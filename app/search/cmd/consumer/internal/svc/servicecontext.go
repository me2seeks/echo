package svc

import (
	"crypto/tls"
	"log"
	"net/http"

	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/me2seeks/echo-hub/app/search/cmd/consumer/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	EsClient *es.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	var cert tls.Certificate
	var err error
	if !c.EsConf.InsecureSkipVerify {
		cert, err = tls.LoadX509KeyPair(c.EsConf.CertFile, c.EsConf.KeyFile)
		if err != nil {
			log.Fatalf("Error loading key pair: %s", err)
		}
	}
	client, err := es.NewClient(es.Config{
		Addresses: c.EsConf.Address,
		Username:  c.EsConf.Username,
		Password:  c.EsConf.Password,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: c.EsConf.InsecureSkipVerify,
			},
		},
	})
	if err != nil {
		log.Fatalf("Error creating the es client: %s", err)
	}
	resp, err := client.Cat.Health()
	if err != nil || resp.StatusCode != 200 {
		log.Fatalf("Error getting response: %s", err)
	}
	return &ServiceContext{
		Config:   c,
		EsClient: client,
	}
}
