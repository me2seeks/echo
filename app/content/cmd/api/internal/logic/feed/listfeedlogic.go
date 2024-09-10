package feed

import (
	"context"
	"strconv"

	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/interaction"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"
	"github.com/me2seeks/echo-hub/common/ctxdata"

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
	userID := ctxdata.GetUIDFromCtx(l.ctx)
	getFeedsByPageResp, err := l.svcCtx.ContentRPC.GetFeedsByPage(l.ctx, &content.GetFeedsByPageReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	resp := &types.GetFeedsByPageResp{}
	resp.Total = getFeedsByPageResp.Total

	for _, feed := range getFeedsByPageResp.Feeds {
		getLikeStatusResp, _ := l.svcCtx.InteractionRPC.GetLikeStatus(l.ctx, &interaction.GetLikeStatusReq{
			UserID:    userID,
			ContentID: feed.Id,
			IsComment: false,
		})
		getFollowStatusResp, _ := l.svcCtx.UsercenterRPC.GetFollowStatus(l.ctx, &usercenter.GetFollowStatusReq{
			UserID:   userID,
			TargetID: feed.UserID,
		})
		resp.Feeds = append(resp.Feeds, types.Feed{
			ID:         strconv.FormatInt(feed.Id, 10),
			UserID:     strconv.FormatInt(feed.UserID, 10),
			Content:    feed.Content,
			Media0:     feed.Media0,
			Media1:     feed.Media1,
			Media2:     feed.Media2,
			Media3:     feed.Media3,
			CreateTime: feed.CreateTime.AsTime().Unix(),
			IsLiked:    getLikeStatusResp.IsLiked,
			IsFollowed: getFollowStatusResp.IsFollowing,
		})
	}

	return resp, nil
}
