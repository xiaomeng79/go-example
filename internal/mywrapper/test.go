package mywrapper

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/server"
)

// logWrapper is a handler wrapper
func TestServerWrap(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Println("这里是中间件")
		return nil
	}
}
