package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/content/model"
	"github.com/me2seeks/echo-hub/common/kqueue"
	"github.com/me2seeks/echo-hub/common/tool"
	"github.com/me2seeks/echo-hub/common/uniqueid"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommentLogic) CreateComment(in *pb.CreateCommentReq) (*pb.CreateCommentResp, error) {
	var comment *model.Comments
	id := uniqueid.GenCommentID()
	comment = &model.Comments{
		Id:      id,
		Content: in.Content,
		FeedId:  in.FeedID,
		UserId:  in.UserID,
		Media0:  sql.NullString{String: in.Media0, Valid: in.Media0 != ""},
		Media1:  sql.NullString{String: in.Media1, Valid: in.Media1 != ""},
		Media2:  sql.NullString{String: in.Media2, Valid: in.Media2 != ""},
		Media3:  sql.NullString{String: in.Media3, Valid: in.Media3 != ""},
	}

	if in.IsComment {
		comment.ParentId = sql.NullInt64{Int64: in.CommentID, Valid: in.CommentID != 0}
	}
	_, err := l.svcCtx.CommentsModel.Insert(l.ctx, nil, comment)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "insert comment failed: %v", err)
	}

	msg := kqueue.Event{
		Type:      kqueue.Comment,
		ID:        in.FeedID,
		IsComment: false,
	}

	if in.IsComment {
		msg.ID = in.CommentID
		msg.IsComment = true
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.MarshalError), "failed to marshal comment event err:%v", err)
	}
	feedIDStr := strconv.FormatInt(in.FeedID, 10)

	err = l.svcCtx.KqPusherClient.PushWithKey(l.ctx, feedIDStr, tool.BytesToString(msgBytes))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.KqPusherError), "failed to push follow event feed:%d,comment:%d,err:%v", in.FeedID, in.CommentID, err)
	}

	return &pb.CreateCommentResp{
		Id: id,
	}, nil
}
