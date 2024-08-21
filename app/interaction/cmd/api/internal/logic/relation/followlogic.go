package relation

import (
	"context"

	"github.com/me2seeks/echo-hub/app/interaction/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/interaction"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"
	"github.com/me2seeks/echo-hub/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// follow
func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowReq) (*types.FollowResp, error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)
	resp, err := l.svcCtx.UsercenterRPC.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		UserId: req.UserId,
	})
	if err != nil || resp.User == nil {
		return nil, err
	}

	_, err = l.svcCtx.InteractionRPC.Follow(l.ctx, &interaction.FollowReq{
		UserId:     userID,
		FolloweeId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &types.FollowResp{}, nil
}
