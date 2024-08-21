package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnfollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowLogic {
	return &UnfollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnfollowLogic) Unfollow(in *pb.UnfollowReq) (*pb.UnfollowResp, error) {
	userRelation, err := l.svcCtx.UserRelationModel.FindOneByFollowerIdFolloweeId(l.ctx, in.UserId, in.FolloweeId)
	if err != nil {
		return &pb.UnfollowResp{}, errors.Wrapf(xerr.NewErrCode(xerr.UnFollowError), "Unfollow err:%v", err)
	}

	err = l.svcCtx.UserRelationModel.DeleteSoft(l.ctx, nil, userRelation)

	return &pb.UnfollowResp{}, err
}
