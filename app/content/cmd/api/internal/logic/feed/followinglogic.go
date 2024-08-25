package feed

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"
	"github.com/me2seeks/echo-hub/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get following feed list by page
func NewFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowingLogic {
	return &FollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowingLogic) Following(req *types.GetFollowingFeedListByPageReq) (*types.GetFollowingFeedListByPageResp, error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)
	followings, err := l.svcCtx.UsercenterRPC.Following(l.ctx, &usercenter.FollowingReq{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	resp, err := l.svcCtx.ContentRPC.GetFollowingFeedListByPage(l.ctx, &content.GetFollowingFeedListByPageReq{
		UserID:   userID,
		TargetID: followings.Following,
		Page:     0,
		PageSize: 0,
	})
	if err != nil {
		return nil, err
	}

	var feeds []types.Feed

	for _, feed := range resp.Feeds {
		feeds = append(feeds, types.Feed{
			ID:          feed.Id,
			UserID:      feed.UserID,
			Content:     feed.Content,
			Media0:      feed.Media0,
			Media1:      feed.Media1,
			Media2:      feed.Media2,
			Media3:      feed.Media3,
			Create_time: feed.CreateTime.AsTime().Unix(),
		})
	}

	return &types.GetFollowingFeedListByPageResp{
		Feeds: feeds,
		Total: resp.Total,
	}, nil
}
