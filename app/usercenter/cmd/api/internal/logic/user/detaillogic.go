package user

import (
	"context"
	"strconv"

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
	userID := ctxdata.GetUIDFromCtx(l.ctx)
	if req.UserID != 0 {
		userID = req.UserID
	}

	getUserInfoResp, err := l.svcCtx.UsercenterRPC.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	var getFollowStatusResp *usercenter.GetFollowStatusResp

	if req.UserID != 0 {
		getFollowStatusResp, err = l.svcCtx.UsercenterRPC.GetFollowStatus(l.ctx, &usercenter.GetFollowStatusReq{
			UserID:   userID,
			TargetID: req.UserID,
		})
		if err != nil {
			return nil, err
		}
	}

	return &types.UserInfoResp{
		UserInfo: types.User{
			Id:       strconv.FormatInt(getUserInfoResp.User.Id, 10),
			Email:    getUserInfoResp.User.Email,
			Nickname: getUserInfoResp.User.Nickname,
			Handle:   getUserInfoResp.User.Handle,
			Sex:      getUserInfoResp.User.Sex,
			Avatar:   getUserInfoResp.User.Avatar,
			Bio:      getUserInfoResp.User.Bio,
			IsFollow: getFollowStatusResp.IsFollow,
		},
	}, nil
}
