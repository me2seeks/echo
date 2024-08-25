package feed

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"
	"github.com/me2seeks/echo-hub/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// create feed
func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateFeedReq) (*types.CreateFeedResp, error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)
	resp, err := l.svcCtx.ContentRPC.CreateFeed(l.ctx, &content.CreateFeedReq{
		UserID:  userID,
		Content: req.Content,
		Media0:  req.Media0,
		Media1:  req.Media1,
		Media2:  req.Media2,
		Media3:  req.Media3,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateFeedResp{
		ID: resp.Id,
	}, nil
}
