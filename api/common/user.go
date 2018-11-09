package common

import (
	"context"
	"github.com/labstack/echo"
	"github.com/xiaomeng79/go-example/cinit"
	"github.com/xiaomeng79/go-example/internal/jwt"
	"github.com/xiaomeng79/go-example/internal/trace"
	"github.com/xiaomeng79/go-example/srv/user/proto"
	"github.com/xiaomeng79/go-log"
)

//获取用户信息
func userinfo(c echo.Context) error {
	ctx := c.Get(cinit.TRACE_CONTEXT).(context.Context)
	ctx, span, _ := trace.TraceIntoContext(ctx, "userinfo")
	defer span.Finish()
	//解析请求参数
	_req := new(user.UserInfoRequest)
	//获取用户ID
	_u := c.Get(cinit.JWT_MSG).(jwt.JWTMsg)
	_req.Id = _u.UserId
	//请求服务

	_rsp, err := UserClient.UserInfo(ctx, _req)
	if err != nil {
		//解析返回的错误信息
		log.Error(err.Error(), ctx)
		return RpcErr(c, err)
	}
	return HandleSuccess(c, _rsp)

}

//登录
func login(c echo.Context) error {
	ctx := c.Get(cinit.TRACE_CONTEXT).(context.Context)
	ctx, span, _ := trace.TraceIntoContext(ctx, "login")
	defer span.Finish()
	//解析请求参数
	_req := new(user.LoginRequest)
	err := c.Bind(&_req)
	if err != nil {
		//解析返回的错误信息
		log.Error(err.Error(), ctx)
		return HandleError(c, BusParamConvertError, err.Error())
	}
	_rsp, err := UserClient.Login(ctx, _req)
	if err != nil {
		//解析返回的错误信息
		log.Error(err.Error(), ctx)
		return RpcErr(c, err)
	}
	//通过验证,设置JWT
	s, err := jwt.Encode(jwt.JWTMsg{_rsp.Id, _rsp.Username})
	if err != nil {
		log.Error(err.Error(), ctx)
		return HandleError(c, ServiceError, err.Error())
	}
	c.Response().Header().Set(cinit.JWT_NAME, "Bearer "+s)
	return HandleError(c, Success)
}
