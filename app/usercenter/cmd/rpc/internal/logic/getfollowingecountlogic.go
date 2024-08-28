package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingeCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowingeCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingeCountLogic {
	return &GetFollowingeCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowingeCountLogic) GetFollowingeCount(in *pb.GetFollowingeCountReq) (*pb.GetFollowingeCountResp, error) {
	count, err := l.svcCtx.UserRelationModel.FindCount(l.ctx, l.svcCtx.UserAuthModel.SelectBuilder().
		Where("follower_id=?", in.UserID), "id")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetFollowingeCount FindCount failed, follower_id:%d , err:%v", in.UserID, err)
	}

	return &pb.GetFollowingeCountResp{
		Count: count,
	}, nil
}
