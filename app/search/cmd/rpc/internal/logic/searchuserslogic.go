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
	searchResp, err := l.svcCtx.EsClient.Search(
		l.svcCtx.EsClient.Search.WithContext(l.ctx),
		l.svcCtx.EsClient.Search.WithIndex("users"),
		l.svcCtx.EsClient.Search.WithQuery(in.Keyword),
		l.svcCtx.EsClient.Search.WithTrackTotalHits(true),
		l.svcCtx.EsClient.Search.WithPretty(),
		// l.svcCtx.EsClient.Search.WithSize(10),
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

	var users []*pb.User
	for _, hit := range response.Hits.Hits {
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
