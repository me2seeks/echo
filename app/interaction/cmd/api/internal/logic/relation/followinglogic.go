package relation

import (
	"context"

	"github.com/me2seeks/echo-hub/app/interaction/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/interaction"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowingsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get followings
func NewFollowingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowingsLogic {
	return &FollowingsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowingsLogic) Following(req *types.FollowingReq) (*types.FollowingResp, error) {
	resp, err := l.svcCtx.InteractionRPC.Following(l.ctx, &interaction.FollowingReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &types.FollowingResp{
		Followings: resp.Following,
	}, nil
}
