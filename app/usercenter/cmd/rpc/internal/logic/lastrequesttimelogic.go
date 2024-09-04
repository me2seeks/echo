package logic

import (
	"context"
	"time"

	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/app/usercenter/model"
	"github.com/me2seeks/echo-hub/common/xerr"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type LastRequestTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLastRequestTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LastRequestTimeLogic {
	return &LastRequestTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LastRequestTimeLogic) LastRequestTime(in *pb.LastRequestTimeReq) (*pb.LastRequestTimeResp, error) {
	userLastRequest, err := l.svcCtx.UserLastRequestModel.FindOne(l.ctx, in.UserID)
	if err != nil {
		if err != model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "LastRequestTime FindOne failed: %v", err)
		}
		_, err = l.svcCtx.UserLastRequestModel.Insert(l.ctx, nil, &model.UserLastRequest{
			UserId:          in.UserID,
			LastRequestTime: time.Now(),
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "LastRequestTime Insert failed: %v", err)
		}
		userLastRequest.LastRequestTime = time.Time{}
	}
	go func() {
		_, err = l.svcCtx.UserLastRequestModel.Update(l.ctx, nil, &model.UserLastRequest{
			UserId:          userLastRequest.UserId,
			LastRequestTime: time.Now(),
		})
		if err != nil {
			l.Error("LastRequestTime Update failed: %v", err)
		}
	}()

	return &pb.LastRequestTimeResp{
		LastRequestTime: timestamppb.New(userLastRequest.LastRequestTime),
	}, nil
}
