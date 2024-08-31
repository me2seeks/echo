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

type GetCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsLogic {
	return &GetCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentsLogic) GetComments(in *pb.GetCommentsReq) (*pb.GetCommentsResp, error) {
	var comments []*model.Comments
	var err error
	if !in.IsComment {
		comments, err = l.svcCtx.CommentsModel.FindAll(l.ctx, l.svcCtx.CommentsModel.SelectBuilder().
			// Columns("id, user_id, content, media0, media1, media2, media3, create_at").
			Where("feed_id = ?", in.Id), "id DESC")
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetCommentListByPage FindAll feedID%d err:%v", in.Id, err)
		}
	} else {
		comments, err = l.svcCtx.CommentsModel.FindAll(l.ctx, l.svcCtx.CommentsModel.SelectBuilder().
			// Columns("id, user_id, content, media0, media1, media2, media3, create_at").
			Where("parent_id = ?", in.Id), "")
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetCommentListByPage FindAll commentID%d err:%v", in.Id, err)
		}
	}
	go func() {
		for _, comment := range comments {
			msg := kqueue.CountEvent{
				Type:      kqueue.View,
				TargetID:  comment.Id,
				IsComment: true,
			}
			msgBytes, err := json.Marshal(msg)
			if err != nil {
				logx.Errorf("IncreaseCommentView Marshal CountEvent failed Type:%d,TargetID:%d,IsComment:%v,err:%v", kqueue.Like, comment.Id, true, err)
			}
			contentIDStr := strconv.FormatInt(comment.Id, 10)

			err = l.svcCtx.KqPusherCounterEventClient.PushWithKey(l.ctx, contentIDStr, tool.BytesToString(msgBytes))
			if err != nil {
				logx.Errorf("IncreaseCommentView PushWithKey failed Type:%d,TargetID:%d,IsComment:%v,err:%v", kqueue.Like, comment.Id, true, err)
			}
		}
	}()

	resp := &pb.GetCommentsResp{}

	for _, comment := range comments {
		resp.Comments = append(resp.Comments, &pb.Comment{
			Id:         comment.Id,
			UserID:     comment.UserId,
			Content:    comment.Content,
			Media0:     tool.GenMediaURL(comment.Media0, l.svcCtx.Config.BaseURL),
			Media1:     tool.GenMediaURL(comment.Media1, l.svcCtx.Config.BaseURL),
			Media2:     tool.GenMediaURL(comment.Media2, l.svcCtx.Config.BaseURL),
			Media3:     tool.GenMediaURL(comment.Media3, l.svcCtx.Config.BaseURL),
			CreateTime: timestamppb.New(comment.CreateAt),
		})
	}
	return resp, nil
}
