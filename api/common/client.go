package common

import (
	"github.com/micro/go-micro/client"
	wo "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/xiaomeng79/go-example/srv/user/proto"
	"time"
)

var (
	UserClient user.UserService
)

func NewUserClient(srvname string) {
	c := client.NewClient(
		client.Retries(0),
		client.WrapCall(wo.NewCallWrapper(opentracing.GlobalTracer())), //tracing
		client.DialTimeout(time.Minute*2),
	)
	UserClient = user.NewUserService(srvname, c)
}
