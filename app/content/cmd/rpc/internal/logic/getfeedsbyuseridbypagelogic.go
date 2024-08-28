package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/tool"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFeedsByUserIDByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFeedsByUserIDByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFeedsByUserIDByPageLogic {
	return &GetFeedsByUserIDByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFeedsByUserIDByPageLogic) GetFeedsByUserIDByPage(in *pb.GetFeedsByUserIDByPageReq) (*pb.GetFeedsByUserIDByPageResp, error) {
	findPageListByPageWithTotalResp, total, err := l.svcCtx.FeedsModel.FindPageListByPageWithTotal(l.ctx, l.svcCtx.FeedsModel.SelectBuilder().
		// Columns("id, user_id, content, media0, media1, media2, media3, create_at").
		Where("user_id in "+tool.BuildQuery(in.UserIDs)).
		Where("create_at > "+in.Before.AsTime().UTC().String()), in.Page, in.PageSize, "id DESC")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetFeedsByUserIDByPage FindPageListByPageWithTotal err:%v", err)
	}

	var feeds []*pb.Feed
	for _, feed := range findPageListByPageWithTotalResp {
		feeds = append(feeds, &pb.Feed{
			Id:         feed.Id,
			UserID:     feed.UserId,
			Content:    feed.Content,
			Media0:     feed.Media0.String,
			Media1:     feed.Media1.String,
			Media2:     feed.Media2.String,
			Media3:     feed.Media3.String,
			CreateTime: timestamppb.New(feed.CreateAt),
		})
	}

	return &pb.GetFeedsByUserIDByPageResp{
		Feeds: feeds,
		Total: total,
	}, nil
}
