package user

import (
	"context"

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

// search users
func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (*types.SearchUsersResp, error) {
	searchUsersResp, err := l.svcCtx.SearchRPC.SearchUsers(l.ctx, &search.SearchReq{
		Keyword:  req.Keyword,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	resp := &types.SearchUsersResp{}
	for _, user := range searchUsersResp.Users {
		resp.Users = append(resp.Users, types.User{
			Id:       user.Id,
			Nickname: user.Nickname,
			Handle:   user.Handle,
			Avatar:   user.Avatar,
			CreateAt: user.CreateAt.AsTime().Unix(),
		})
	}

	return resp, nil
}
