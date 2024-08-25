package logic

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/kqueue"
	"github.com/me2seeks/echo-hub/common/tool"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLikeLogic {
	return &DeleteLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLikeLogic) DeleteLike(in *pb.DeleteLikeReq) (*pb.DeleteLikeResp, error) {
	if in.IsComment {
		err := l.svcCtx.CommentLikesModel.DeleteByUserIdContentId(l.ctx, nil, in.UserID, in.Id)
		if err != nil {
			return &pb.DeleteLikeResp{}, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "delete comment like err:%v", err)
		}
	} else {
		err := l.svcCtx.FeedLikesModel.DeleteByUserIdContentId(l.ctx, nil, in.UserID, in.Id)
		if err != nil {
			return &pb.DeleteLikeResp{}, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "delete feed like err:%v", err)
		}
	}

	msg := kqueue.Event{
		Type:      kqueue.UnLike,
		ID:        in.Id,
		IsComment: false,
	}
	if in.IsComment {
		msg.IsComment = true
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.MarshalError), "failed to marshal UnLike event ,err:%v", err)
	}
	contentIDStr := strconv.FormatInt(in.Id, 10)

	err = l.svcCtx.KqPusherClient.PushWithKey(l.ctx, contentIDStr, tool.BytesToString(msgBytes))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.KqPusherError), "failed to push UnLike event userID:%d,contentIDStr:%d,is_comment:%t,err:%v", in.UserID, in.Id, in.IsComment, err)
	}

	return &pb.DeleteLikeResp{}, nil
}
