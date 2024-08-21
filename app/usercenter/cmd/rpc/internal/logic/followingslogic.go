package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowingsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowingsLogic {
	return &FollowingsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// follow list
func (l *FollowingsLogic) Followings(in *pb.FollowingsReq) (*pb.FollowingsResp, error) {
	ids, err := l.svcCtx.UserRelationModel.FindFollowees(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find followings")
	}

	return &pb.FollowingsResp{
		Followings: ids,
	}, err
}
