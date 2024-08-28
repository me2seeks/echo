package user

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"
	"github.com/me2seeks/echo-hub/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// update user info
func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateUserInfoReq) (*types.UpdateUserInfoResp, error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)

	_, err := l.svcCtx.UsercenterRPC.UpdateUserInfo(l.ctx, &usercenter.UpdateUserInfoReq{
		UserID:   userID,
		Nickname: req.Nickname,
		Sex:      req.Sex,
		Avatar:   req.Avatar,
		Bio:      req.Bio,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateUserInfoResp{}, nil
}
