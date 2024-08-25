package logic

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

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
	feeds, total, err := l.svcCtx.FeedsModel.FindPageListByPageWithTotal(l.ctx, l.svcCtx.FeedsModel.SelectBuilder().Columns("id, user_id, content, media0, media1, media2, media3, create_at").Where("user_id = ?", in.UserID), in.Page, in.PageSize, "")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetfeedListByPage FindPageListByPageWithTotal UserID%d err:%v", in.UserID, err)
	}

	var feedList []*pb.Feed
	for _, feed := range feeds {
		feedList = append(feedList, &pb.Feed{
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

	return &pb.GetFeedListByPageResp{
		Total: total,
		Feeds: feedList,
	}, nil
}
