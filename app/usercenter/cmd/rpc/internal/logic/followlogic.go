package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/usercenter/model"
	"github.com/me2seeks/echo-hub/common/uniqueid"
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
			return &pb.FollowResp{}, errors.Wrap(err, "insert user relation error")
		}
		return &pb.FollowResp{}, err
	}
	if err != nil {
		return nil, errors.Wrap(err, "restore user relation error")
	}
	return &pb.FollowResp{}, err
}
