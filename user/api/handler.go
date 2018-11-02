package api

import (
	"context"
	api "github.com/micro/micro/api/proto"
	pb "github.com/xiaomeng79/go-example/user/srv/proto"
)

type UserHandler struct {
	Client pb.UserService
}

func (u *UserHandler) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	return nil
}
