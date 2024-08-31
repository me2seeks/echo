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

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserID)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetUserInfo FindOne failed , UserId:%d , err:%v", in.UserID, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "GetUserInfo  not exit UserID:%d", in.UserID)
	}
	resp := &pb.GetUserInfoResp{}
	resp.User = &pb.User{
		Id:       user.Id,
		Email:    user.Email,
		Nickname: user.Nickname,
		Handel:   user.Handle,
		Sex:      int32(user.Sex),
		Avatar:   user.Avatar,
		Bio:      user.Bio,
	}

	return resp, nil
}
