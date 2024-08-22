package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowingLogic {
	return &FollowingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowingLogic) Following(in *pb.FollowingReq) (*pb.FollowingResp, error) {
	ids, err := l.svcCtx.UserRelationModel.FindFollowees(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find followings")
	}

	return &pb.FollowingResp{
		Following: ids,
	}, err
}
