package relation

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"

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

func (l *FollowingsLogic) Followings(req *types.FollowingsReq) (*types.FollowingsResp, error) {
	resp, err := l.svcCtx.UsercenterRPC.Followings(l.ctx, &usercenter.FollowingsReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &types.FollowingsResp{
		Followings: resp.Followings,
	}, nil
}
