package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/counter/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCounterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCounterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCounterLogic {
	return &GetUserCounterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCounterLogic) GetUserCounter(in *pb.GetUserCounterRequest) (*pb.GetUserCounterResponse, error) {
	// TODO use redis hash
	userCount, err := l.svcCtx.UserStateModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetUserCounter error: %v", err)
	}

	return &pb.GetUserCounterResponse{
		FollowingCount: userCount.FollowingCount,
		FollowerCount:  userCount.FollowerCount,
		FeedCount:      userCount.FeedCount,
	}, nil
}
