package feed

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"
	"github.com/me2seeks/echo-hub/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFollowingFeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get following feed list by page
func NewListFollowingFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowingFeedLogic {
	return &ListFollowingFeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFollowingFeedLogic) ListFollowingFeed(req *types.GetFollowingFeedsByPageReq) (*types.GetFollowingFeedsByPageResp, error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)
	followings, err := l.svcCtx.UsercenterRPC.Following(l.ctx, &usercenter.FollowingReq{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	contentResp, err := l.svcCtx.ContentRPC.GetFollowingFeedListByPage(l.ctx, &content.GetFollowingFeedListByPageReq{
		UserID:   userID,
		TargetID: followings.Following,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	logx.Debugf("contentResp: %v", contentResp)

	var resp *types.GetFollowingFeedsByPageResp

	_ = copier.Copy(resp, contentResp)

	return resp, nil
}
