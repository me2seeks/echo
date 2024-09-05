package feed

import (
	"context"
	"strconv"

	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"
	"github.com/me2seeks/echo-hub/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// create feed comment
func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateFeedCommentReq) (*types.CreateFeedCommentResp, error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)
	resp, err := l.svcCtx.ContentRPC.CreateComment(l.ctx, &content.CreateCommentReq{
		FeedID:    req.FeedID,
		UserID:    userID,
		Content:   req.Content,
		Media0:    req.Media0,
		Media1:    req.Media1,
		Media2:    req.Media2,
		Media3:    req.Media3,
		IsComment: false,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateFeedCommentResp{
		ID: strconv.FormatInt(resp.Id, 10),
	}, nil
}
