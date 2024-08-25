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

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFeedLogic {
	return &DeleteFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteFeedLogic) DeleteFeed(in *pb.DeleteFeedReq) (*pb.DeleteFeedResp, error) {
	err := l.svcCtx.FeedsModel.DeleteSoft(l.ctx, nil, &model.Feeds{
		Id:     in.Id,
		UserId: in.UserID,
	})
	if err != nil {
		return nil, err
	}

	msg := kqueue.Event{
		Type:      kqueue.DeleteFeed,
		ID:        in.UserID,
		IsComment: false,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.MarshalError), "failed to marshal DeleteFeed event, err: %v", err)
	}
	userIDStr := strconv.FormatInt(in.UserID, 10)
	err = l.svcCtx.KqPusherClient.PushWithKey(l.ctx, userIDStr, tool.BytesToString(msgBytes))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.KqPusherError), "failed to push DeleteFeed event userID:%d,err:%v", in.UserID, err)
	}

	return &pb.DeleteFeedResp{}, nil
}
