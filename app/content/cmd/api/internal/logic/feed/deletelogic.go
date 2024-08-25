package feed

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"
	"github.com/me2seeks/echo-hub/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// delete feed
func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DeleteFeedReq) (*types.DeleteFeedResp, error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)
	_, err := l.svcCtx.ContentRPC.DeleteFeed(l.ctx, &content.DeleteFeedReq{
		Id:     req.ID,
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeleteFeedResp{}, nil
}
