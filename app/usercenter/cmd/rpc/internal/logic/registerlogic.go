package logic

import (
	"context"
	"encoding/json"
	"time"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"
	"github.com/me2seeks/echo-hub/app/usercenter/model"
	"github.com/me2seeks/echo-hub/common/kqueue"
	"github.com/me2seeks/echo-hub/common/tool"
	"github.com/me2seeks/echo-hub/common/uniqueid"
	"github.com/me2seeks/echo-hub/common/xerr"

	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrUserAlreadyRegisterError = xerr.NewErrMsg("user has been registered")

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		if err != model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Register email:%s,err:%v", in.Email, err)
		}
	}
	if user != nil {
		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "Register user exists email:%s,err:%v", in.Email, err)
	}

	user = &model.User{
		Email:    in.Email,
		Nickname: in.Nickname,
		Handle:   in.Handle,
	}

	if len(user.Nickname) == 0 {
		user.Nickname = tool.Krand(8, tool.KcRandKindAll)
	}
	user.Password, err = tool.EncryptWithBcrypt(in.Password)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.EncryptError), "Register EncryptWithBcrypt err:%v,user:%+v", err, user)
	}
	user.Id = uniqueid.GenUserID()

	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		_, err := l.svcCtx.UserModel.Insert(ctx, session, user)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Register db user Insert err:%v,user:%+v", err, user)
		}
		userAuth := new(model.UserAuth)
		userAuth.UserId = user.Id
		userAuth.AuthKey = in.AuthKey
		userAuth.AuthType = in.AuthType
		if _, err := l.svcCtx.UserAuthModel.Insert(ctx, session, userAuth); err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Register db user_auth Insert err:%v,userAuth:%v", err, userAuth)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	go func() {
		msg := kqueue.EsEvent{
			Type:      kqueue.Register,
			UserID:    user.Id,
			Handle:    user.Handle,
			Content:   user.Nickname,
			CreatedAt: time.Now(),
		}
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			logx.Errorf("Register Marshal  UserId:%d,Handle:%s,NickName:%s,CreatedAt:%v,err:%v", user.Id, user.Handle, user.Nickname, time.Now(), err)
			return
		}
		err = l.svcCtx.KqPusherEsEventClient.Push(l.ctx, tool.BytesToString(msgBytes))
		if err != nil {
			logx.Errorf("Register PushWithKey  UserId:%d,Handle:%s,NickName:%s,CreatedAt:%v,err:%v", user.Id, user.Handle, user.Nickname, time.Now(), err)
			return
		}
	}()

	// 2、Generate the token, so that the service doesn't call rpc internally
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserID: user.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", user.Id)
	}

	return &usercenter.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
