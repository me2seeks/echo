package feed

import (
	"context"

	"github.com/me2seeks/echo-hub/app/interaction/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/interaction"
	"github.com/me2seeks/echo-hub/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnlikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// unlike
func NewUnlikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnlikeLogic {
	return &UnlikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnlikeLogic) Unlike(req *types.DeleteFeedLikeReq) (*types.DeleteFeedLikeResp, error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)
	_, err := l.svcCtx.InteractionRPC.DeleteLike(l.ctx, &interaction.DeleteLikeReq{
		UserID:    userID,
		Id:        req.ID,
		IsComment: false,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeleteFeedLikeResp{}, nil
}
