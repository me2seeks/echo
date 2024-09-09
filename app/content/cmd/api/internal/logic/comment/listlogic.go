package comment

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

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get comment list by page
func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.GetCommentsByPageReq) (*types.GetCommentsByPageResp, error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)
	getCommentsByPageResp, err := l.svcCtx.ContentRPC.GetCommentsByPage(l.ctx, &content.GetCommentsByPageReq{
		Id:        req.CommentID,
		Page:      req.Page,
		PageSize:  req.PageSize,
		IsComment: true,
	})
	if err != nil {
		return nil, err
	}
	resp := &types.GetCommentsByPageResp{}
	resp.Total = getCommentsByPageResp.Total

	for _, comment := range getCommentsByPageResp.Comments {
		getLikeStatusResp, _ := l.svcCtx.InteractionRPC.GetLikeStatus(l.ctx, &interaction.GetLikeStatusReq{
			UserID:    userID,
			ContentID: comment.Id,
			IsComment: true,
		})
		resp.Comments = append(resp.Comments, types.Comment{
			ID:         strconv.FormatInt(comment.Id, 10),
			UserID:     strconv.FormatInt(comment.UserID, 10),
			Content:    comment.Content,
			Media0:     comment.Media0,
			Media1:     comment.Media1,
			Media2:     comment.Media2,
			Media3:     comment.Media3,
			CreateTime: comment.CreateTime.AsTime().Unix(),
			IsLiked:    getLikeStatusResp.IsLiked,
		})
	}
	return resp, nil
}
