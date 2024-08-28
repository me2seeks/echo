package feed

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"
	"github.com/me2seeks/echo-hub/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFollowingFeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get following feed list by page
func NewListFollowingFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowingFeedLogic {
	return &ListFollowingFeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFollowingFeedLogic) ListFollowingFeed(req *types.GetFollowingFeedsByPageReq) (*types.GetFollowingFeedsByPageResp, error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)
	LastRequestTime, err := l.svcCtx.UsercenterRPC.LastRequestTime(l.ctx, &usercenter.LastRequestTimeReq{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	followingResp, err := l.svcCtx.UsercenterRPC.Following(l.ctx, &usercenter.FollowingReq{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	getFeedsByUserIDByPageResp, err := l.svcCtx.ContentRPC.GetFeedsByUserIDByPage(l.ctx, &content.GetFeedsByUserIDByPageReq{
		UserIDs:  followingResp.Following,
		Before:   LastRequestTime.LastRequestTime,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var feeds []types.Feed
	for _, feed := range getFeedsByUserIDByPageResp.Feeds {
		feeds = append(feeds, types.Feed{
			ID:          feed.Id,
			UserID:      feed.UserID,
			Content:     feed.Content,
			Media0:      feed.Media0,
			Media1:      feed.Media1,
			Media2:      feed.Media2,
			Media3:      feed.Media3,
			Create_time: feed.CreateTime.AsTime().Unix(),
		})
	}

	return &types.GetFollowingFeedsByPageResp{
		Feeds: feeds,
		Total: getFeedsByUserIDByPageResp.Total,
	}, nil
}
