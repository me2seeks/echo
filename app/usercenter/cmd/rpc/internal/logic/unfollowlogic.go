package logic

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/kqueue"
	"github.com/me2seeks/echo-hub/common/tool"
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
		return &pb.UnfollowResp{}, errors.Wrapf(xerr.NewErrCode(xerr.UnFollowError), "not followed err:%v", err)
	}

	err = l.svcCtx.UserRelationModel.DeleteSoft(l.ctx, nil, userRelation)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UnFollowError), "unfollow err:%v", err)
	}

	msg := kqueue.CountEvent{
		Type:      kqueue.UnFollow,
		ID:        in.FolloweeId,
		IsComment: false,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.MarshalError), "failed to marshal unfollow event ,err:%v", err)
	}
	followeeIDStr := strconv.FormatInt(in.FolloweeId, 10)

	err = l.svcCtx.KqPusherCounterEventClient.PushWithKey(l.ctx, followeeIDStr, tool.BytesToString(msgBytes))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.KqPusherError), "failed push unfollow event followeeID:%d,err:%v", in.FolloweeId, err)
	}
	return &pb.UnfollowResp{}, err
}
