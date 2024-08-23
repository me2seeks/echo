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

type GetFeedListByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFeedListByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFeedListByPageLogic {
	return &GetFeedListByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFeedListByPageLogic) GetFeedListByPage(in *pb.GetFeedListByPageReq) (*pb.GetFeedListByPageResp, error) {
	feeds, total, err := l.svcCtx.FeedsModel.FindPageListByPageWithTotal(l.ctx, l.svcCtx.FeedsModel.SelectBuilder().Columns("*").Where("user_id = ?", in.UserID).Where("create_at > ?", in.StartTime.AsTime()), in.Page, in.PageSize, "")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetfeedListByPage FindPageListByPageWithTotal UserID%d err:%v", in.UserID, err)
	}

	var feedList []*pb.Feed
	err = copier.Copy(&feedList, &feeds)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CopyError), "GetFeedListByPage copier.Copy err:%v", err)
	}

	return &pb.GetFeedListByPageResp{
		Total: total,
		Feed:  feedList,
	}, nil
}
