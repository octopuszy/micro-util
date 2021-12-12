package util

import (
	"context"
	"log"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
)

type clientWrap struct {
	client.Client
}

func (c *clientWrap) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption)  error{
	err := hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(err error) error {
		log.Println("hystrix error...")
		log.Println(err)
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

func NewHystrixClientWrap() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrap{c}
	}
}