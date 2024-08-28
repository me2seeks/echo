package logic

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/usercenter/model"
	"github.com/me2seeks/echo-hub/common/globalkey"
	"github.com/me2seeks/echo-hub/common/kqueue"
	"github.com/me2seeks/echo-hub/common/tool"
	"github.com/me2seeks/echo-hub/common/uniqueid"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowLogic) Follow(in *pb.FollowReq) (*pb.FollowResp, error) {
	relation, err := l.svcCtx.UserRelationModel.
		FindOneByFollowerIdFolloweeIdWithOutDelState(l.ctx, in.UserID, in.FolloweeID)
	if err != nil {
		if err == model.ErrNotFound {
			_, err = l.svcCtx.UserRelationModel.Insert(l.ctx, nil, &model.UserRelation{
				Id:         uniqueid.GenUserRelationID(),
				FollowerId: in.UserID,
				FolloweeId: in.FolloweeID,
			})
			if err != nil {
				return &pb.FollowResp{}, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Follow insert user relation err:%v", err)
			}
		} else {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Follow FindOneByFollowerIdFolloweeIdWithOutDelState failed, userID:%d, followeeID:%d, err:%v", in.UserID, in.FolloweeID, err)
		}
	} else if relation.DelState == globalkey.DelStateYes {
		err = l.svcCtx.UserRelationModel.RestoreSoft(l.ctx, nil, relation.Id)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Follow restore user relation error userID:%d, followeeID:%d, err:%v", in.UserID, in.FolloweeID, err)
		}
	}

	go func() {
		msg := kqueue.CountEvent{
			Type:      kqueue.Follow,
			SourceID:  in.UserID,
			TargetID:  in.FolloweeID,
			IsComment: false,
		}
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			logx.Errorf("Follow Marshal  SourceID:%d,TargetID:%d,err:%v", in.UserID, in.FolloweeID, err)
			return
		}
		followeeIDStr := strconv.FormatInt(in.FolloweeID, 10)

		err = l.svcCtx.KqPusherCounterEventClient.PushWithKey(l.ctx, followeeIDStr, tool.BytesToString(msgBytes))
		if err != nil {
			logx.Errorf("Follow PushWithKey  SourceID:%d,TargetID:%d,err:%v", in.UserID, in.FolloweeID, err)
			return
		}
	}()

	return &pb.FollowResp{}, err
}
