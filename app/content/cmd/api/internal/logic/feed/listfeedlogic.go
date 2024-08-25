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
	resp, err := l.svcCtx.ContentRPC.GetFeedListByPage(l.ctx, &content.GetFeedListByPageReq{
		UserID:   req.UserID,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var feeds []types.Feed

	for _, feed := range resp.Feeds {
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

	return &types.GetFeedsByPageResp{
		Feeds: feeds,
		Total: resp.Total,
	}, nil
}
