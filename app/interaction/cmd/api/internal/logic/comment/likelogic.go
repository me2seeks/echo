package comment

import (
	"context"

	"github.com/me2seeks/echo-hub/app/interaction/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/app/interaction/cmd/rpc/interaction"
	"github.com/me2seeks/echo-hub/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// like
func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.CreateCommentLikeReq) (*types.CreateCommentLikeResp, error) {
	userID := ctxdata.GetUIDFromCtx(l.ctx)
	_, err := l.svcCtx.InteractionRPC.CreateLike(l.ctx, &interaction.CreateLikeReq{
		UserID:    userID,
		Id:        req.ID,
		IsComment: true,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateCommentLikeResp{}, nil
}
