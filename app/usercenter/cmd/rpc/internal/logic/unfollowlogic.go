package logic

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/usercenter/model"
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
	userRelation, err := l.svcCtx.UserRelationModel.FindOneByFollowerIdFolloweeId(l.ctx, in.UserID, in.FolloweeID)
	if err != nil {
		if err == model.ErrNotFound {
			return &pb.UnfollowResp{}, nil
		}
		return &pb.UnfollowResp{}, errors.Wrapf(xerr.NewErrCode(xerr.UnFollowError), "Unfollow err:%v", err)
	}

	err = l.svcCtx.UserRelationModel.DeleteSoft(l.ctx, nil, userRelation)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UnFollowError), "Unfollow err:%v", err)
	}

	go func() {
		msg := kqueue.CountEvent{
			Type:      kqueue.UnFollow,
			SourceID:  in.UserID,
			TargetID:  in.FolloweeID,
			IsComment: false,
		}
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			logx.Errorf("Unfollow Marshal  TargetID:%d,err:%v", in.FolloweeID, err)
			return
		}
		followeeIDStr := strconv.FormatInt(in.FolloweeID, 10)

		err = l.svcCtx.KqPusherCounterEventClient.PushWithKey(l.ctx, followeeIDStr, tool.BytesToString(msgBytes))
		if err != nil {
			logx.Errorf("Unfollow PushWithKey  TargetID:%d,err:%v", in.FolloweeID, err)
			return
		}
	}()

	return &pb.UnfollowResp{}, err
}
