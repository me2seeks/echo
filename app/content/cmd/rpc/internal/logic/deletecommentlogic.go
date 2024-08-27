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
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "delete comment failed: %v", err)
	}

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
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.MarshalError), "failed to marshal delete comment event err:%v", err)
	}
	IDStr := strconv.FormatInt(in.Id, 10)

	err = l.svcCtx.KqPusherCounterEventClient.PushWithKey(l.ctx, IDStr, tool.BytesToString(msgBytes))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.KqPusherError), "failed push delete comment event ParentID:%d,IsComment:%t,err:%v", in.ParentID, in.IsComment, err)
	}

	return &pb.DeleteCommentResp{}, nil
}
