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

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentLogic) DeleteComment(in *pb.DeleteCommentReq) (*pb.DeleteCommentResp, error) {
	err := l.svcCtx.CommentsModel.DeleteSoft(l.ctx, nil, &model.Comments{
		Id: in.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "DeleteComment delete comment failed: %v", err)
	}

	go func() {
		msg := kqueue.CountEvent{
			Type:      kqueue.UnComment,
			TargetID:  in.Id,
			IsComment: false,
		}

		if in.IsComment {
			msg.TargetID = in.ParentID
			msg.IsComment = true
		}

		msgBytes, err := json.Marshal(msg)
		if err != nil {
			if !in.IsComment {
				logx.Errorf("CreateFeed Marshal CountEvent failed  Type:%d,TargetID:%d,IsComment:%t,err:%v", kqueue.Comment, in.Id, true, err)
			} else {
				logx.Errorf("CreateFeed Marshal CountEvent failed  Type:%d,TargetID:%d,IsComment:%t,err:%v", kqueue.UnComment, in.ParentID, false, err)
			}
			return
		}
		IDStr := strconv.FormatInt(in.Id, 10)

		err = l.svcCtx.KqPusherCounterEventClient.PushWithKey(l.ctx, IDStr, tool.BytesToString(msgBytes))
		if err != nil {
			if !in.IsComment {
				logx.Errorf("CreateFeed Push CountEvent failed  Type:%d,TargetID:%d,IsComment:%t,err:%v", kqueue.Comment, in.Id, true, err)
			} else {
				logx.Errorf("CreateFeed Push CountEvent failed  Type:%d,TargetID:%d,IsComment:%t,err:%v", kqueue.UnComment, in.ParentID, false, err)
			}
			return
		}
	}()

	return &pb.DeleteCommentResp{}, nil
}
