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

type GetCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get comment comment list by page
func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentLogic) GetComment(req *types.GetCommentReq) (*types.GetCommentResp, error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)
	getCommentByIDResp, err := l.svcCtx.ContentRPC.GetCommentByID(l.ctx, &content.GetCommentByIDReq{
		Id: req.CommentID,
	})
	if err != nil {
		return nil, err
	}
	resp := &types.GetCommentResp{
		Comment: types.Comment{
			ID:         strconv.FormatInt(getCommentByIDResp.Comment.Id, 10),
			UserID:     strconv.FormatInt(getCommentByIDResp.Comment.UserID, 10),
			FeedID:     strconv.FormatInt(getCommentByIDResp.Comment.FeedID, 10),
			Content:    getCommentByIDResp.Comment.Content,
			Media0:     getCommentByIDResp.Comment.Media0,
			Media1:     getCommentByIDResp.Comment.Media1,
			Media2:     getCommentByIDResp.Comment.Media2,
			Media3:     getCommentByIDResp.Comment.Media3,
			CreateTime: getCommentByIDResp.Comment.CreateTime.AsTime().Unix(),
		},
	}
	if userID != 0 {
		getLikeStatusResp, _ := l.svcCtx.InteractionRPC.GetLikeStatus(l.ctx, &interaction.GetLikeStatusReq{
			UserID:    userID,
			ContentID: getCommentByIDResp.Comment.Id,
			IsComment: true,
		})
		resp.Comment.IsLiked = getLikeStatusResp.IsLiked
	}

	return resp, nil
}
