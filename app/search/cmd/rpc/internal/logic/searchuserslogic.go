package logic

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/me2seeks/echo-hub/app/search/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/search/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/es"
	"github.com/me2seeks/echo-hub/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUsersLogic {
	return &SearchUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUsersLogic) SearchUsers(in *pb.SearchReq) (*pb.SearchUsersResp, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"nickname": in.Keyword,
						},
					},
					{
						"match": map[string]interface{}{
							"handle": in.Keyword,
						},
					},
				},
			},
		},
	}
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.MarshalError), "SearchUsers json.Marshal err:%v", err)
	}
	req := esapi.SearchRequest{
		Index: []string{"users"},
		Body:  strings.NewReader(string(queryJSON)),
	}
	res, err := req.Do(l.ctx, l.svcCtx.EsClient)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.EsError), "SearchUsers req.Do err:%v", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.EsError), "SearchUsers res.IsError err:%v", err)
	}

	var searchResponse es.SearchUsersResponse
	if err := json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UnmarshalError), "SearchUsers json.Decode err:%v", err)
	}

	for _, hit := range searchResponse.Hits.Hits {
		logx.Infof("User: %v", hit.Source)
	}
	var users []*pb.User
	for _, hit := range searchResponse.Hits.Hits {
		users = append(users, &pb.User{
			Id:       hit.Source.ID,
			Nickname: hit.Source.Nickname,
			Handle:   hit.Source.Handle,
		})
	}

	return &pb.SearchUsersResp{
		Users: users,
	}, nil
}
