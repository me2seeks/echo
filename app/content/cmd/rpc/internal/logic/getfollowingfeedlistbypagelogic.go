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

type GetFollowingFeedListByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowingFeedListByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingFeedListByPageLogic {
	return &GetFollowingFeedListByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowingFeedListByPageLogic) GetFollowingFeedListByPage(in *pb.GetFollowingFeedListByPageReq) (*pb.GetFollowingFeedListByPageResp, error) {
	userLastRequest, err := l.svcCtx.UserLastRequestModel.FindOne(l.ctx, in.UserID)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetfeedListByPage FindOne UserID%d err:%v", in.UserID, err)
	}
	feeds, total, err := l.svcCtx.FeedsModel.FindPageListByPageWithTotal(l.ctx, l.svcCtx.FeedsModel.SelectBuilder().Columns("id, user_id, content, media0, media1, media2, media3, create_at").Where("user_id IN (?,?,...)", in.UserID).Where("create_at > ?", userLastRequest.LastRequestTime), in.Page, in.PageSize, "")
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

	return &pb.GetFollowingFeedListByPageResp{
		Feeds: feedList,
		Total: total,
	}, nil
}
