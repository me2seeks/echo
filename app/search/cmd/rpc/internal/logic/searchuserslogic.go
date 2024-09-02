package logic

import (
	"context"
	"encoding/json"
	"io"

	"github.com/me2seeks/echo-hub/app/search/cmd/rpc/internal/svc"
	"github.com/me2seeks/echo-hub/app/search/cmd/rpc/pb"
	"github.com/me2seeks/echo-hub/common/es"
	"github.com/me2seeks/echo-hub/common/xerr"
	"google.golang.org/protobuf/types/known/timestamppb"

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
	// TODO 1、search users page page_size 2、 match nickname handle
	// query := elastic.NewMultiMatchQuery(in.Keyword, "handle", "nickname").
	// 	Type("best_fields").
	// 	Fuzziness("AUTO")

	// // 将查询转换为 JSON
	// queryJSON, err := query.Source()
	// if err != nil {
	// 	return nil, errors.Wrapf(xerr.NewErrCode(xerr.EsError), "SearchUsers query.Source err:%v", err)
	// }
	// queryBody, err := json.Marshal(queryJSON)
	// if err != nil {
	// 	return nil, errors.Wrapf(xerr.NewErrCode(xerr.EsError), "SearchUsers json.Marshal err:%v", err)
	// }

	// // 创建 SearchRequest
	// searchRequest := &elastic.SearchRequest{
	// 	Index:          []string{"users"},
	// 	Body:           bytes.NewReader(queryBody),
	// 	TrackTotalHits: true,
	// 	Pretty:         true,
	// 	Size:           int(in.PageSize),
	// }

	// // 执行搜索请求
	// searchResp, err := l.svcCtx.EsClient.Search(
	// 	l.svcCtx.EsClient.Search.WithContext(l.ctx),
	// 	l.svcCtx.EsClient.Search.WithIndex(searchRequest.Index...),
	// 	l.svcCtx.EsClient.Search.WithBody(searchRequest.Body),
	// 	l.svcCtx.EsClient.Search.WithTrackTotalHits(searchRequest.TrackTotalHits),
	// 	l.svcCtx.EsClient.Search.WithPretty(searchRequest.Pretty),
	// 	l.svcCtx.EsClient.Search.WithSize(searchRequest.Size),
	// )
	searchResp, err := l.svcCtx.EsClient.Search(
		l.svcCtx.EsClient.Search.WithContext(l.ctx),
		l.svcCtx.EsClient.Search.WithIndex("users"),
		l.svcCtx.EsClient.Search.WithQuery(in.Keyword),
		l.svcCtx.EsClient.Search.WithTrackTotalHits(true),
		l.svcCtx.EsClient.Search.WithPretty(),
		l.svcCtx.EsClient.Search.WithSize(int(in.PageSize)),
	)
	if err != nil || searchResp.StatusCode != 200 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.EsError), "SearchUsers err:%v", err)
	}

	body, err := io.ReadAll(searchResp.Body)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ReadBodyError), "SearchUsers io.ReadAll err:%v", err)
	}
	var response es.SearchUsersResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UnmarshalError), "SearchUsers json.Unmarshal err:%v", err)
	}

	resp := &pb.SearchUsersResp{}
	for _, hit := range response.Hits.Hits {
		user := &pb.User{
			Id:       hit.Source.ID,
			Nickname: hit.Source.Nickname,
			Handle:   hit.Source.Handle,
			CreateAt: timestamppb.New(hit.Source.CreatedAt),
		}
		if hit.Source.Avatar != "" {
			user.Avatar = l.svcCtx.Config.BaseURL + hit.Source.Avatar
		}
		resp.Users = append(resp.Users, user)
	}
	return resp, nil
}
