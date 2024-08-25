package comment

import (
	"context"

	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// delete comment
func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DeleteCommentReq) (resp *types.DeleteCommentResp, err error) {
	_, err = l.svcCtx.ContentRPC.DeleteComment(l.ctx, &content.DeleteCommentReq{
		Id:        req.ID,
		ParentID:  req.ParentID,
		IsComment: req.ParentID != 0,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeleteCommentResp{}, nil
}
