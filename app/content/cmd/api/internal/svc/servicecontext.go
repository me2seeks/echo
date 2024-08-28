package svc

import (
	"github.com/me2seeks/echo-hub/app/content/cmd/api/internal/config"
	"github.com/me2seeks/echo-hub/app/content/cmd/rpc/content"
	"github.com/me2seeks/echo-hub/app/usercenter/cmd/rpc/usercenter"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	ContentRPC    content.Content
	UsercenterRPC usercenter.Usercenter

	MinioClient *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	minioClient, err := minio.New(c.MiniConf.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.MiniConf.AccessKey, c.MiniConf.SecretKey, ""),
		Secure: true,
	})
	if err != nil {
		logx.Errorf("failed to create minio client: %s", err)
	}
	return &ServiceContext{
		Config:        c,
		ContentRPC:    content.NewContent(zrpc.MustNewClient(c.ContentRPCConf)),
		UsercenterRPC: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRPCConf)),
		MinioClient:   minioClient,
	}
}
