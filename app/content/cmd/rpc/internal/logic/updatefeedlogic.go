package logic

import (
	"context"
	"database/sql"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/content/model"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFeedLogic {
	return &UpdateFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFeedLogic) UpdateFeed(in *pb.UpdateFeedReq) (*pb.UpdateFeedResp, error) {
	err := l.svcCtx.FeedsModel.UpdateWithVersion(l.ctx, nil, &model.Feeds{
		Id:      in.Id,
		Content: in.Content,
		Media0:  sql.NullString{String: in.Media0, Valid: in.Media0 != ""},
		Media1:  sql.NullString{String: in.Media1, Valid: in.Media1 != ""},
		Media2:  sql.NullString{String: in.Media2, Valid: in.Media2 != ""},
		Media3:  sql.NullString{String: in.Media3, Valid: in.Media3 != ""},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "UpdateFeed  failed feed id: %d", in.Id)
	}

	return &pb.UpdateFeedResp{}, nil
}
