package feed

import (
	"context"
	"strconv"

	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/interaction"
	"github.com/me2seeks/echo-hub/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get feed by feedID
func NewGetFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFeedLogic {
	return &GetFeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFeedLogic) GetFeed(req *types.GetFeedReq) (resp *types.GetFeedResp, err error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)
	getFeedsByIDResp, err := l.svcCtx.ContentRPC.GetFeedsByID(l.ctx, &content.GetFeedsByIDReq{
		IDs: []int64{req.ID},
	})
	if err != nil {
		return nil, err
	}

	if len(getFeedsByIDResp.Feeds) == 0 {
		return nil, nil
	}

	getLikeStatusResp, _ := l.svcCtx.InteractionRPC.GetLikeStatus(l.ctx, &interaction.GetLikeStatusReq{
		UserID:    userID,
		ContentID: getFeedsByIDResp.Feeds[0].Id,
		IsComment: false,
	})

	return &types.GetFeedResp{
		Feed: types.Feed{
			ID:         strconv.FormatInt(getFeedsByIDResp.Feeds[0].Id, 10),
			UserID:     strconv.FormatInt(getFeedsByIDResp.Feeds[0].UserID, 10),
			Content:    getFeedsByIDResp.Feeds[0].Content,
			Media0:     getFeedsByIDResp.Feeds[0].Media0,
			Media1:     getFeedsByIDResp.Feeds[0].Media1,
			Media2:     getFeedsByIDResp.Feeds[0].Media2,
			Media3:     getFeedsByIDResp.Feeds[0].Media3,
			CreateTime: getFeedsByIDResp.Feeds[0].CreateTime.AsTime().Unix(),
			IsLiked:    getLikeStatusResp.IsLiked,
		},
	}, nil
}
