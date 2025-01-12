package upload

import (
	"context"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/svc"
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/types"
	"github.com/me2seeks/echo-hub/common/upload"

	"github.com/zeromicro/go-zero/core/logx"
)

type PresignedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPresignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PresignedLogic {
	return &PresignedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PresignedLogic) Presign(req *types.PresignReq) (*types.PresignResp, error) {
	resp := &types.PresignResp{}
	for _, object := range req.Objects {
		switch object.FileType {
		case upload.Avatar:
			object.FileName = "avatars/" + uuid.New().String() + filepath.Ext(object.FileName)
		case upload.FeedImg:
			object.FileName = "feed_imgs/" + uuid.New().String() + filepath.Ext(object.FileName)
		case upload.FeedVideo:
			object.FileName = "feed_videos/" + uuid.New().String() + filepath.Ext(object.FileName)
		case upload.FeedGIF:
			object.FileName = "feed_gifs/" + uuid.New().String() + filepath.Ext(object.FileName)
		default:
			object.FileName = "unknown/" + uuid.New().String()
		}
		presignedURL, err := l.svcCtx.MinioClient.PresignedPutObject(
			l.ctx, l.svcCtx.Config.MinioConf.BucketName,
			object.FileName,
			time.Duration(l.svcCtx.Config.MinioConf.Expires)*time.Minute)
		if err != nil {
			return nil, err
		}
		resp.Urls = append(resp.Urls, presignedURL.String())

	}
	return resp, nil
}
