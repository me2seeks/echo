package logic

import (
	"context"
	"encoding/json"
	"io"

	"github.com/me2seeks/echo-hub/app/search/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/search/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/es"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFeedsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchFeedsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFeedsLogic {
	return &SearchFeedsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchFeedsLogic) SearchFeeds(in *pb.SearchReq) (*pb.SearchFeedsResp, error) {
	searchResp, err := l.svcCtx.EsClient.Search(
		l.svcCtx.EsClient.Search.WithContext(l.ctx),
		l.svcCtx.EsClient.Search.WithIndex("feeds"),
		l.svcCtx.EsClient.Search.WithQuery(in.Keyword),
		l.svcCtx.EsClient.Search.WithTrackTotalHits(true),
		l.svcCtx.EsClient.Search.WithPretty(),
		// l.svcCtx.EsClient.Search.WithSize(10),
	)
	if err != nil || searchResp.StatusCode != 200 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.EsError), "SearchContent err:%v", err)
	}

	body, err := io.ReadAll(searchResp.Body)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ReadBodyError), "SearchContent io.ReadAll err:%v", err)
	}
	var response es.SearchFeedsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UnmarshalError), "SearchContent json.Unmarshal err:%v", err)
	}

	var contentID []int64
	for _, hit := range response.Hits.Hits {
		contentID = append(contentID, hit.Source.ID)
	}

	return &pb.SearchFeedsResp{
		ContentID: contentID,
	}, nil
}
