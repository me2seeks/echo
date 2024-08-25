package feed

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get comment list by page
func NewListCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentLogic {
	return &ListCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCommentLogic) ListComment(req *types.GetFeedCommentsByPageReq) (*types.GetFeedCommentsByPageResp, error) {
	resp, err := l.svcCtx.ContentRPC.GetCommentListByPage(l.ctx, &content.GetCommentListByPageReq{
		Id:        req.FeedID,
		Page:      req.Page,
		PageSize:  req.PageSize,
		IsComment: false,
	})
	if err != nil {
		return nil, err
	}

	var comments []types.Comment

	for _, comment := range resp.Comments {
		comments = append(comments, types.Comment{
			ID:          comment.Id,
			UserID:      comment.UserID,
			Content:     comment.Content,
			Media0:      comment.Media0,
			Media1:      comment.Media1,
			Media2:      comment.Media2,
			Media3:      comment.Media3,
			Create_time: comment.CreateTime.AsTime().Unix(),
		})
	}

	return &types.GetFeedCommentsByPageResp{
		Comments: comments,
	}, nil
}
