package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	service.ServiceConf
	KqConsumerConf kq.KqConf
	EsConf         struct {
		Address            []string
		Username           string
		Password           string
		CertFile           string
		KeyFile            string
		InsecureSkipVerify bool
	}
}