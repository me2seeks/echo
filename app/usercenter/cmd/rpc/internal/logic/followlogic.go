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
	relation, err := l.svcCtx.UserRelationModel.FindOneByFollowerIdFolloweeIdWithOutDelState(l.ctx, in.UserId, in.FolloweeId)
	if err != nil {
		if err == model.ErrNotFound {
			_, err = l.svcCtx.UserRelationModel.Insert(l.ctx, nil, &model.UserRelation{
				Id:         uniqueid.GenUserRelationID(),
				FollowerId: in.UserId,
				FolloweeId: in.FolloweeId,
			})
			if err != nil {
				return &pb.FollowResp{}, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "insert user relation err:%v", err)
			}
		} else {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "find user relation error userID:%d, followeeID:%d, err:%v", in.UserId, in.FolloweeId, err)
		}
	} else if relation.DelState == globalkey.DelStateYes {
		err = l.svcCtx.UserRelationModel.RestoreSoft(l.ctx, nil, relation.Id)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "restore user relation error userID:%d, followeeID:%d, err:%v", in.UserId, in.FolloweeId, err)
		}
	}

	msg := kqueue.CountEvent{
		Type:      kqueue.Follow,
		SourceID:  in.UserId,
		TargetID:  in.FolloweeId,
		IsComment: false,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.MarshalError), "failed to marshal follow event ,err:%v", err)
	}
	followeeIDStr := strconv.FormatInt(in.FolloweeId, 10)

	err = l.svcCtx.KqPusherCounterEventClient.PushWithKey(l.ctx, followeeIDStr, tool.BytesToString(msgBytes))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.KqPusherError), "failed to push follow event followeeID:%d,err:%v", in.FolloweeId, err)
	}

	return &pb.FollowResp{}, err
}
