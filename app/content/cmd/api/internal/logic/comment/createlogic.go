package comment

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"
	"github.com/me2seeks/echo-hub/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// create comment
func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateCommentReq) (*types.CreateCommentResp, error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)
	resp, err := l.svcCtx.ContentRPC.CreateComment(l.ctx, &content.CreateCommentReq{
		CommentID: req.CommentID,
		UserID:    userID,
		Content:   req.Content,
		Media0:    req.Media0,
		Media1:    req.Media1,
		Media2:    req.Media2,
		Media3:    req.Media3,
		IsComment: true,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateCommentResp{
		ID: resp.Id,
	}, nil
}
