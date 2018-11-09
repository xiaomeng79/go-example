package mywrapper

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"github.com/xiaomeng79/go-log"
)

// log wrapper logs every time a request is made
type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	log.Info("客户端记录 服务:"+req.Service()+" 方法:"+req.Method()+"请求信息", ctx)
	return l.Client.Call(ctx, req, rsp)
}

// Implements client.Wrapper as logWrapper
func LogClientWrap(c client.Client) client.Client {
	return &logWrapper{c}
}

// logWrapper is a handler wrapper
func LogServerWrap(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		err := fn(ctx, req, rsp)
		if err != nil {
			log.Info("错误记录,服务:"+req.Service()+" 方法:"+req.Method()+err.Error(), ctx)
		}
		log.Infof("请求信息:%+v\n", req, ctx)
		log.Infof("响应信息:%+v\n", rsp, ctx)
		return err
	}
}
