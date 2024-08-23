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

type CreateFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFeedLogic {
	return &CreateFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateFeedLogic) CreateFeed(in *pb.CreateFeedReq) (*pb.CreateFeedResp, error) {
	_, err := l.svcCtx.FeedsModel.Insert(l.ctx, nil, &model.Feeds{
		Id:      uniqueid.GenFeedID(),
		Content: in.Content,
		UserId:  in.UserID,
		Media0:  sql.NullString{String: in.Media0, Valid: in.Media0 != ""},
		Media1:  sql.NullString{String: in.Media1, Valid: in.Media1 != ""},
		Media2:  sql.NullString{String: in.Media2, Valid: in.Media2 != ""},
		Media3:  sql.NullString{String: in.Media3, Valid: in.Media3 != ""},
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "insert feed failed: %v", err)
	}
	return &pb.CreateFeedResp{}, nil
}
