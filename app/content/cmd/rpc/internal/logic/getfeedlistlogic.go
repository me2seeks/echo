//nolint:dupl
package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFeedListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFeedListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFeedListLogic {
	return &GetFeedListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFeedListLogic) GetFeedList(in *pb.GetFeedListReq) (*pb.GetFeedListResp, error) {
	feeds, err := l.svcCtx.FeedsModel.FindAll(l.ctx, l.svcCtx.FeedsModel.SelectBuilder().Columns("*").Where("user_id=?", in.UserID), "")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetfeedList FindAll userID%d err:%v", in.UserID, err)
	}

	var feedList []*pb.Feed
	err = copier.Copy(&feedList, &feeds)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CopyError), "GetfeedList copier.Copy err:%v", err)
	}

	return &pb.GetFeedListResp{}, nil
}
