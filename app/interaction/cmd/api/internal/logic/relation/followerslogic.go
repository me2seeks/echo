package relation

import (
	"context"

	"github.com/me2seeks/echo-hub/app/interaction/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/interaction"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get followers
func NewFollowersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowersLogic {
	return &FollowersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowersLogic) Followers(req *types.FollowersReq) (*types.FollowersResp, error) {
	resp, err := l.svcCtx.InteractionRPC.Followers(l.ctx, &interaction.FollowersReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &types.FollowersResp{
		Followers: resp.Followers,
	}, nil
}
