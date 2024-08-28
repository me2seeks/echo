package feed

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get feed list by page
func NewListFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFeedLogic {
	return &ListFeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFeedLogic) ListFeed(req *types.GetFeedsByPageReq) (*types.GetFeedsByPageResp, error) {
	getFeedsByUserIDByPageResp, err := l.svcCtx.ContentRPC.GetFeedsByUserIDByPage(l.ctx, &content.GetFeedsByUserIDByPageReq{
		UserIDs:  []int64{req.UserID},
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	resp := &types.GetFeedsByPageResp{}
	resp.Total = getFeedsByUserIDByPageResp.Total

	for _, feed := range getFeedsByUserIDByPageResp.Feeds {
		resp.Feeds = append(resp.Feeds, types.Feed{
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

	return resp, nil
}
