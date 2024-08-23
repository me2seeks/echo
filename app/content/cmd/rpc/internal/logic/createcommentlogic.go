package logic

import (
	"context"
	"database/sql"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/content/model"
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
	_, err := l.svcCtx.CommentsModel.Insert(l.ctx, nil, &model.Comments{
		Id:      uniqueid.GenCommentID(),
		Content: in.Content,
		FeedId:  in.FeedID,
		UserId:  in.UserID,
		Media0:  sql.NullString{String: in.Media0, Valid: in.Media0 != ""},
		Media1:  sql.NullString{String: in.Media1, Valid: in.Media1 != ""},
		Media2:  sql.NullString{String: in.Media2, Valid: in.Media2 != ""},
		Media3:  sql.NullString{String: in.Media3, Valid: in.Media3 != ""},
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "insert comment failed: %v", err)
	}
	return &pb.CreateCommentResp{}, nil
}
