package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/xiaomeng79/go-example/data"
	"github.com/xiaomeng79/go-example/internal/errors"
	pb "github.com/xiaomeng79/go-example/srv/user/proto"
	"github.com/xiaomeng79/go-log"
	"net/http"
)

type UserHandler struct {
	name2info map[string]*pb.UserInfo
	id2info   map[int32]*pb.UserInfo
}

//从json文件导入测试数据
func loadUserInfo(path string) (map[string]*pb.UserInfo, map[int32]*pb.UserInfo) {
	file := data.MustAsset(path)
	infos := []*pb.UserInfo{}
	if err := json.Unmarshal(file, &infos); err != nil {
		log.Fatalf("Failed to load json:%v", err)
	}
	id2info := make(map[int32]*pb.UserInfo)
	name2info := make(map[string]*pb.UserInfo)
	for _, info := range infos {
		id2info[info.Id] = info
		name2info[info.Username] = info
	}
	return name2info, id2info
}

func (u *UserHandler) Login(ctx context.Context, req *pb.LoginRequest, rsp *pb.LoginResponse) error {
	fmt.Println(req.Username)
	fmt.Println(req.Password)
	log.Debugf("%+v\n", u.name2info, ctx)
	if info, ok := u.name2info[req.Username]; ok {
		log.Debugf("%+v\n", info, ctx)
		if info.Password != req.Password {
			return errors.New("密码不正确", http.StatusBadRequest)
		}
		rsp.Username = info.Username
		rsp.Id = info.Id
		rsp.Email = info.Email
		return nil
	}
	return errors.New("用户不存在", http.StatusBadRequest)

}

func (u *UserHandler) UserInfo(ctx context.Context, req *pb.UserInfoRequest, rsp *pb.UserInfoResponse) error {
	if _, ok := u.id2info[req.Id]; !ok {
		return errors.New("用户不存在", http.StatusBadRequest)
	}
	info := u.id2info[req.Id]
	rsp.Id = info.Id
	rsp.Email = info.Email
	rsp.Username = info.Username
	return nil
}
