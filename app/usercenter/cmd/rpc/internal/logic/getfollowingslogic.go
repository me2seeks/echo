package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingsLogic {
	return &GetFollowingsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowingsLogic) GetFollowings(in *pb.GetFollowingsReq) (*pb.GetFollowingsResp, error) {
	ids, err := l.svcCtx.UserRelationModel.FindFollowees(l.ctx, in.UserID)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetFollowings  FindFollowees failed err:%v", err)
	}

	return &pb.GetFollowingsResp{
		IDs: ids,
	}, nil
}
