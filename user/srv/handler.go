package srv

import (
	"context"
	"fmt"
	pb "github.com/xiaomeng79/go-example/user/srv/proto"
)

type UserHandler struct {
}

func (u UserHandler) Login(ctx context.Context, req *pb.LoginRequest, rsp *pb.LoginResponse) error {
	fmt.Println(req.Username)
	fmt.Println(req.Password)
	rsp.Base.Code = 0
	rsp.Base.Msg = "success"
	return nil
}
