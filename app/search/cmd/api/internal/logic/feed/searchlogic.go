package feed

import (
	"context"

	"github.com/jinzhu/copier"
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
	searchContentResp, err := l.svcCtx.SearchRPC.SearchContent(l.ctx, &search.SearchReq{
		Keyword:  req.Keyword,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	contentResp, err := l.svcCtx.ContentRPC.GetFeedsByIDByPage(l.ctx, &content.GetFeedsByIDByPageReq{
		FeedID: searchContentResp.ContentID,
	})
	if err != nil {
		return nil, err
	}
	var resp types.SearchFeedsResp
	_ = copier.Copy(&resp, contentResp.Feeds)

	return &resp, nil
}