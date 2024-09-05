package relation

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"
	"github.com/me2seeks/echo-hub/common/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowingsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get following
func NewFollowingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowingsLogic {
	return &FollowingsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowingsLogic) Followings(req *types.FollowingsReq) (*types.FollowingsResp, error) {
	getFollowingsResp, err := l.svcCtx.UsercenterRPC.GetFollowings(l.ctx, &usercenter.GetFollowingsReq{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &types.FollowingsResp{
		Followings: tool.ConvertInt64SliceToStringSlice(getFollowingsResp.IDs),
	}, nil
}
