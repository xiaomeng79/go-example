package api

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"
	pb "github.com/xiaomeng79/go-example/user/srv/proto"
)

type User struct {
	Client pb.UserService
}

func (u *User) Login(ctx context.Context, req *api.Request, rsp *api.Response) error {
	username, ok := req.Post["username"]

	if !ok || len(username.Values) == 0 {
		return errors.BadRequest("api.user", "username is not found")
	}

	password, ok := req.Post["password"]

	if !ok || len(password.Values) == 0 {
		return errors.BadRequest("api.user", "password is not found")
	}

	response, err := u.Client.Login(ctx, &pb.LoginRequest{Username: username.Values[0], Password: password.Values[0]})
	if err != nil {
		return nil
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code": response.Base.Code,
		"msg":  response.Base.Msg,
	})
	rsp.Body = string(b)
	return nil
}
