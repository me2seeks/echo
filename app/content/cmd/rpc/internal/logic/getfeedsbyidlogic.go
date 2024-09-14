package logic

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/content/model"
	"github.com/me2seeks/echo-hub/common/kqueue"
	"github.com/me2seeks/echo-hub/common/tool"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFeedsByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFeedsByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFeedsByIDLogic {
	return &GetFeedsByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFeedsByIDLogic) GetFeedsByID(in *pb.GetFeedsByIDReq) (*pb.GetFeedsByIDResp, error) {
	if len(in.IDs) == 0 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.InvalidParameter), "GetFeedsByID invalid params")
	}
	findResp := make([]*model.Feeds, 0)
	var err error
	if len(in.IDs) == 1 {
		findResp[0], err = l.svcCtx.FeedsModel.FindOne(l.ctx, in.IDs[0])
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetFeedsByID FindOne err:%v", err)
		}
	} else {
		findResp, err = l.svcCtx.FeedsModel.FindAll(l.ctx, l.svcCtx.FeedsModel.
			SelectBuilder().Where("id in "+tool.BuildQuery(in.IDs)), "id DESC")
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetFeedsByID FindAll err:%v", err)
		}
	}
	go func() {
		for _, feed := range findResp {
			msg := kqueue.CountEvent{
				Type:      kqueue.View,
				TargetID:  feed.Id,
				IsComment: false,
			}
			msgBytes, err := json.Marshal(msg)
			if err != nil {
				logx.Errorf("IncreaseFeedView Marshal CountEvent failed Type:%d,TargetID:%d,IsComment:%v,err:%v", kqueue.Like, feed.Id, false, err)
			}
			contentIDStr := strconv.FormatInt(feed.Id, 10)

			err = l.svcCtx.KqPusherCounterEventClient.PushWithKey(l.ctx, contentIDStr, tool.BytesToString(msgBytes))
			if err != nil {
				logx.Errorf("IncreaseFeedView PushWithKey failed Type:%d,TargetID:%d,IsComment:%v,err:%v", kqueue.Like, feed.Id, false, err)
			}
		}
	}()

	resp := &pb.GetFeedsByIDResp{}

	for _, feed := range findResp {
		resp.Feeds = append(resp.Feeds, &pb.Feed{
			Id:         feed.Id,
			UserID:     feed.UserId,
			Content:    feed.Content,
			Media0:     tool.GenMediaURL(feed.Media0, l.svcCtx.Config.BaseURL),
			Media1:     tool.GenMediaURL(feed.Media1, l.svcCtx.Config.BaseURL),
			Media2:     tool.GenMediaURL(feed.Media2, l.svcCtx.Config.BaseURL),
			Media3:     tool.GenMediaURL(feed.Media3, l.svcCtx.Config.BaseURL),
			CreateTime: timestamppb.New(feed.CreateAt),
		})
	}

	return resp, nil
}
