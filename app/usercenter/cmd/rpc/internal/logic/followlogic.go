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
	err := l.svcCtx.UserRelationModel.RestoreSoft(l.ctx, nil, in.UserId, in.FolloweeId)
	if err == model.ErrNoRowsUpdate {
		_, err = l.svcCtx.UserRelationModel.Insert(l.ctx, nil, &model.UserRelation{
			Id:         uniqueid.GenUserRelationID(),
			FollowerId: in.UserId,
			FolloweeId: in.FolloweeId,
		})
		if err != nil {
			return &pb.FollowResp{}, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "insert user relation err:%v", err)
		}
		return &pb.FollowResp{}, err
	}
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "restore user relation error userID:%d,followeeID:%d,err:%v", in.UserId, in.FolloweeId, err)
	}

	msg := kqueue.Event{
		Type:      kqueue.Follow,
		ID:        in.FolloweeId,
		IsComment: false,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, errors.Wrap(err, "marshal event error")
	}
	FolloweeIDStr := strconv.FormatInt(in.FolloweeId, 10)

	err = l.svcCtx.KqPusherClient.PushWithKey(l.ctx, FolloweeIDStr, tool.BytesToString(msgBytes))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.KqPusherError), "failed push follow event followeeID:%d,err:%v", in.FolloweeId, err)
	}

	return &pb.FollowResp{}, err
}
