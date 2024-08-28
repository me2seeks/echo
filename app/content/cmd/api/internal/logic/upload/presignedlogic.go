package upload

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PresignedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPresignedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PresignedLogic {
	return &PresignedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PresignedLogic) Presigned(req *types.PresignedReq) (*types.PresignedResp, error) {
	req.FileName = uuid.New().String() + "_" + req.FileName
	presignedURL, err := l.svcCtx.MinioClient.PresignedPutObject(
		l.ctx, l.svcCtx.Config.MiniConf.BucketName,
		req.FileName,
		time.Duration(l.svcCtx.Config.MiniConf.Expires)*time.Minute)
	if err != nil {
		return nil, err
	}

	return &types.PresignedResp{
		Url: presignedURL.String(),
	}, nil
}
