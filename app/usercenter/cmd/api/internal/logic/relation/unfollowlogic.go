package relation

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"
	"github.com/me2seeks/echo-hub/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnfollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// unfollow
func NewUnfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowLogic {
	return &UnfollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnfollowLogic) Unfollow(req *types.UnfollowReq) (*types.UnfollowResp, error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)

	_, err := l.svcCtx.UsercenterRPC.Unfollow(l.ctx, &usercenter.UnfollowReq{
		UserID:     userID,
		FolloweeID: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &types.UnfollowResp{}, nil
}
