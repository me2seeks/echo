package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserInfoResp{}, nil
}
