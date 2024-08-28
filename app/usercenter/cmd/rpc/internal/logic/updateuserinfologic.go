package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/usercenter/model"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

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
	user := &model.User{
		Id: in.UserId,
	}
	if in.Nickname != "" {
		user.Nickname = in.Nickname
	}
	if in.Sex != 0 {
		user.Sex = int64(in.Sex)
	}
	if in.Avatar != "" {
		user.Avatar = in.Avatar
	}
	if in.Bio != "" {
		user.Bio = in.Bio
	}
	err := l.svcCtx.UserModel.UpdateWithVersion(l.ctx, nil, user)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "UpdateUserInfo err:%v", err)
	}
	return &pb.UpdateUserInfoResp{}, nil
}
