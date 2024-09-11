package relation

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"
	"github.com/me2seeks/echo-hub/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get follow status
func NewFollowStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowStatusLogic {
	return &FollowStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowStatusLogic) FollowStatus(req *types.FollowStatusReq) (*types.FollowStatusResp, error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)

	getFollowStatusResp, err := l.svcCtx.UsercenterRPC.GetFollowStatus(l.ctx, &usercenter.GetFollowStatusReq{
		UserID:   userID,
		TargetID: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &types.FollowStatusResp{
		IsFollow: getFollowStatusResp.IsFollow,
	}, nil
}
