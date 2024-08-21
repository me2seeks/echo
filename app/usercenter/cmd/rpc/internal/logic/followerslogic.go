package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowersLogic {
	return &FollowersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 被关注列表
func (l *FollowersLogic) Followers(in *pb.FollowersReq) (*pb.FollowersResp, error) {
	ids, err := l.svcCtx.UserRelationModel.FindFollowers(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find followers")
	}

	return &pb.FollowersResp{
		Followers: ids,
	}, err
}
