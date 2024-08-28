package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerCountLogic {
	return &GetFollowerCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerCountLogic) GetFollowerCount(in *pb.GetFollowerCountReq) (*pb.GetFollowerCountResp, error) {
	count, err := l.svcCtx.UserRelationModel.FindCount(l.ctx, l.svcCtx.UserAuthModel.SelectBuilder().
		Where("followee_id=?", in.UserID), "id")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetFollowerCount FindCount failed, followee_id:%d , err:%v", in.UserID, err)
	}
	return &pb.GetFollowerCountResp{
		Count: count,
	}, nil
}
