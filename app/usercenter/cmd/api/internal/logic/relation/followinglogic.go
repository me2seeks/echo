package relation

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get following
func NewFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowingLogic {
	return &FollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowingLogic) Following(req *types.FollowingReq) (*types.FollowingResp, error) {
	resp, err := l.svcCtx.UsercenterRPC.Following(l.ctx, &usercenter.FollowingReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &types.FollowingResp{
		Followings: resp.Following,
	}, nil
}
