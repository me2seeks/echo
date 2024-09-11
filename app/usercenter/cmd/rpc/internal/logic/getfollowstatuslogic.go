package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowStatusLogic {
	return &GetFollowStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowStatusLogic) GetFollowStatus(in *pb.GetFollowStatusReq) (*pb.GetFollowStatusResp, error) {
	var resp pb.GetFollowStatusResp
	findOneByFollowerIDFolloweeIDResp, err := l.svcCtx.UserRelationModel.FindOneByFollowerIdFolloweeId(l.ctx, in.UserID, in.TargetID)
	if err == nil && findOneByFollowerIDFolloweeIDResp != nil {
		resp.IsFollow = true
	}

	return &resp, nil
}
