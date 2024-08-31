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
			return &pb.DeleteLikeResp{}, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "DeleteLike delete comment like err:%v", err)
		}
	} else {
		err := l.svcCtx.FeedLikesModel.DeleteByUserIdContentId(l.ctx, nil, in.UserID, in.Id)
		if err != nil {
			return &pb.DeleteLikeResp{}, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "DeleteLike delete feed like err:%v", err)
		}
	}

	go func() {
		msg := kqueue.CountEvent{
			Type:      kqueue.UnLike,
			TargetID:  in.Id,
			IsComment: false,
		}
		if in.IsComment {
			msg.IsComment = true
		}
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			logx.Errorf("DeleteLike Marshal CountEvent failed Type:%d,TargetID:%d,IsComment:%v,err:%v", kqueue.Like, in.Id, in.IsComment, err)
			return
		}
		contentIDStr := strconv.FormatInt(in.Id, 10)

		err = l.svcCtx.KqPusherClient.PushWithKey(l.ctx, contentIDStr, tool.BytesToString(msgBytes))
		if err != nil {
			logx.Errorf("DeleteLike PushWithKey failed contentIDStr:%s, msg:%s, err:%v", contentIDStr, string(msgBytes), err)
			return
		}
	}()

	return &pb.DeleteLikeResp{}, nil
}
