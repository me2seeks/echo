package logic

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/me2seeks/echo-hub/app/search/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/search/cmd/rpc/mqs"
	"github.com/me2seeks/echo-hub/app/search/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchContentLogic {
	return &SearchContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchContentLogic) SearchContent(in *pb.SearchReq) (*pb.SearchContentResp, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"content": in.Keyword,
						},
					},
				},
			},
		},
	}
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.MarshalError), "SearchContent json.Marshal err:%v", err)
	}
	req := esapi.SearchRequest{
		Index: []string{"feeds"},
		Body:  strings.NewReader(string(queryJSON)),
	}
	res, err := req.Do(l.ctx, l.svcCtx.EsClient)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.EsError), "es request err:%v", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.EsError), "es response err:%v", res.String())
	}

	var searchResponse mqs.SearchFeedsResponse
	if err := json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UnmarshalError), "unmarshal searchResponse err:%v", err)
	}
	for _, hit := range searchResponse.Hits.Hits {
		logx.Infof("Feed: %v", hit.Source)
	}

	var contentID []int64
	for _, hit := range searchResponse.Hits.Hits {
		contentID = append(contentID, hit.Source.ID)
	}

	return &pb.SearchContentResp{
		ContentID: contentID,
	}, nil
}
