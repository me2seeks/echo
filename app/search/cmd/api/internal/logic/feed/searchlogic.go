package feed

import (
	"context"
	"strconv"

	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"
	"github.com/me2seeks/echo-hub/app/search/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/search/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/search/cmd/rpc/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// search feeds
func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (*types.SearchFeedsResp, error) {
	searchContentResp, err := l.svcCtx.SearchRPC.SearchFeeds(l.ctx, &search.SearchReq{
		Keyword:  req.Keyword,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	resp := &types.SearchFeedsResp{}

	if len(searchContentResp.ContentID) != 0 {
		contentResp, err := l.svcCtx.ContentRPC.GetFeedsByID(l.ctx, &content.GetFeedsByIDReq{
			IDs: searchContentResp.ContentID,
		})
		if err != nil {
			return nil, err
		}
		// _ = copier.Copy(&resp.Feeds, contentResp.Feeds)
		for _, feed := range contentResp.Feeds {
			resp.Feeds = append(resp.Feeds, types.Feed{
				Id:        strconv.FormatInt(feed.Id, 10),
				Content:   feed.Content,
				UserID:    strconv.FormatInt(feed.UserID, 10),
				Media0:    feed.Media0,
				Media1:    feed.Media1,
				Media2:    feed.Media2,
				Media3:    feed.Media3,
				CreatTime: feed.CreateTime.AsTime().Unix(),
			})
		}
	}

	return resp, nil
}
