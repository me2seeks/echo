package user

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"

	"github.com/me2seeks/echo-hub/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get user info
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (*types.UserInfoResp, error) {
	if req.UserID == 0 {
		req.UserID = ctxdata.GetUIDFromCtx(l.ctx)
	}

	getUserInfoResp, err := l.svcCtx.UsercenterRPC.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	resp := &types.UserInfoResp{}
	_ = copier.Copy(&resp.UserInfo, getUserInfoResp.User)

	return resp, nil
}
