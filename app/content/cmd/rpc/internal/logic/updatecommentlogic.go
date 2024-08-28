package logic

import (
	"context"
	"database/sql"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/content/model"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommentLogic) UpdateComment(in *pb.UpdateCommentReq) (*pb.UpdateCommentResp, error) {
	err := l.svcCtx.CommentsModel.UpdateWithVersion(l.ctx, nil, &model.Comments{
		Id:      in.Id,
		Content: in.Content,
		Media0:  sql.NullString{String: in.Media0, Valid: in.Media0 != ""},
		Media1:  sql.NullString{String: in.Media1, Valid: in.Media1 != ""},
		Media2:  sql.NullString{String: in.Media2, Valid: in.Media2 != ""},
		Media3:  sql.NullString{String: in.Media3, Valid: in.Media3 != ""},
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "UpdateComment failed comment id: %d", in.Id)
	}

	return &pb.UpdateCommentResp{}, nil
}
