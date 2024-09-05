package relation

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"
	"github.com/me2seeks/echo-hub/common/tool"

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
	getFollowersResp, err := l.svcCtx.UsercenterRPC.GetFollowers(l.ctx, &usercenter.GetFollowersReq{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &types.FollowersResp{
		Followers: tool.ConvertInt64SliceToStringSlice(getFollowersResp.IDs),
	}, nil
}
