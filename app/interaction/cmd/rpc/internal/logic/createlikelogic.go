package logic

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/interaction/model"
	"github.com/me2seeks/echo-hub/common/kqueue"
	"github.com/me2seeks/echo-hub/common/tool"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLikeLogic {
	return &CreateLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLikeLogic) CreateLike(in *pb.CreateLikeReq) (*pb.CreateLikeResp, error) {
	if in.IsComment {
		_, err := l.svcCtx.CommentLikesModel.Insert(l.ctx, nil, &model.CommentLikes{
			UserId:    in.UserID,
			ContentId: in.Id,
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "insert comment like err:%v", err)
		}
	} else {
		_, err := l.svcCtx.FeedLikesModel.Insert(l.ctx, nil, &model.FeedLikes{
			UserId:    in.UserID,
			ContentId: in.Id,
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "insert feed like err:%v", err)
		}
	}

	msg := kqueue.Event{
		Type:      kqueue.Like,
		ID:        in.Id,
		IsComment: false,
	}
	if in.IsComment {
		msg.IsComment = true
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.MarshalError), "failed to marshal Like event ,err:%v", err)
	}
	contentIDStr := strconv.FormatInt(in.Id, 10)

	err = l.svcCtx.KqPusherClient.PushWithKey(l.ctx, contentIDStr, tool.BytesToString(msgBytes))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.KqPusherError), "failed to push like event userID:%d,contentIDStr:%d,is_comment:%t,err:%v", in.UserID, in.Id, in.IsComment, err)
	}

	return &pb.CreateLikeResp{}, nil
}
