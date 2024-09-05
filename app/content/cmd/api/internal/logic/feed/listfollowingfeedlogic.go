package feed

import (
	"context"
	"strconv"

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
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	followingResp, err := l.svcCtx.UsercenterRPC.GetFollowings(l.ctx, &usercenter.GetFollowingsReq{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	getFeedsByUserIDByPageResp, err := l.svcCtx.ContentRPC.GetFeedsByUserIDByPage(l.ctx, &content.GetFeedsByUserIDByPageReq{
		UserIDs:  followingResp.IDs,
		Before:   LastRequestTime.LastRequestTime,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	resp := &types.GetFollowingFeedsByPageResp{}
	resp.Total = getFeedsByUserIDByPageResp.Total

	for _, feed := range getFeedsByUserIDByPageResp.Feeds {
		resp.Feeds = append(resp.Feeds, types.Feed{
			ID:         strconv.FormatInt(feed.Id, 10),
			UserID:     strconv.FormatInt(feed.UserID, 10),
			Content:    feed.Content,
			Media0:     feed.Media0,
			Media1:     feed.Media1,
			Media2:     feed.Media2,
			Media3:     feed.Media3,
			CreateTime: feed.CreateTime.AsTime().Unix(),
		})
	}

	return resp, nil
}
