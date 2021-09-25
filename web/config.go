package web

import (
	kratos "github.com/ory/kratos-client-go"
)

var (
	KratosPublicURL    = "http://127.0.0.1:4000"
	KratosPublicClient = NewKratosClient(KratosPublicURL)
)

func NewKratosClient(endpoint string) *kratos.APIClient {
	conf := kratos.NewConfiguration()
	conf.Servers = kratos.ServerConfigurations{{URL: endpoint}}
	return kratos.NewAPIClient(conf)
}
