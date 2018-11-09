package common

import (
	"github.com/labstack/echo"
	"github.com/xiaomeng79/go-example/internal/errors"
	"net/http"
	"strconv"
	"strings"
)

//定义错误信息

type ErrorNo int64

func (e ErrorNo) String() string {
	return strconv.FormatInt(int64(e), 10)
}

const (
	//成功
	Success ErrorNo = 0
	//请求
	ReqPathError          = 10001
	ReqVersionNoExist     = 10002
	ReqInterfaceNoExist   = 10003
	ReqCommandNoExist     = 10004
	ReqInterfaceNoSupport = 10005
	ReqNoAllow            = 10006

	//公共
	CommonPageError         = 20001 //分页错误
	CommonParamError        = 20002 //参数错误
	CommonParamConvertError = 20003 //参数转换错误
	CommonSignError         = 20004 //签名错误
	CommonAppError          = 20008 //app错误

	//业务参数
	BusParamError        = 30001 //参数错误
	BusParamConvertError = 30002 //参数转换错误

	//权限

	//服务端处理异常
	ServiceError = 50001
)

var ReturnMsg map[ErrorNo]string = map[ErrorNo]string{
	//成功
	Success: "success",
	//请求
	ReqPathError:          "请求路径错误",
	ReqVersionNoExist:     "请求版本不存在",
	ReqInterfaceNoExist:   "请求接口不存在",
	ReqCommandNoExist:     "请求命令不存在",
	ReqInterfaceNoSupport: "接口在当前请求命令中未被支持",
	ReqNoAllow:            "请求不允许",

	//公共
	CommonPageError:         "分页错误",        //分页错误
	CommonParamError:        "公共参数错误",      //参数错误
	CommonParamConvertError: "公共参数转换错误",    //参数转换错误
	CommonSignError:         "公共参数签名错误",    //签名错误
	CommonAppError:          "商户账户不存在或不可用", //

	//业务参数
	BusParamError:        "业务参数错误",   //参数错误
	BusParamConvertError: "业务参数转换错误", //参数转换错误

	//服务端处理异常
	ServiceError: "服务端处理异常",
}

/**
异常错误公共
*/
func errCommon(code ErrorNo, errmsg ...string) interface{} {
	return map[string]interface{}{
		"code":    code.String(),
		"message": ReturnMsg[code] + strings.Join(errmsg, " "),
	}
}

/**
RPC错误处理
*/
func RpcErr(c echo.Context, err error) error {
	errr := errors.Parse(err.Error())
	if errr.Code == http.StatusBadRequest {
		return c.JSON(http.StatusBadRequest, errCommon(BusParamError, ":", errr.Detail))
	} else if errr.Code == http.StatusUnauthorized {
		return c.JSON(http.StatusUnauthorized, errCommon(BusParamError, ":", errr.Detail))
	} else {
		return c.JSON(http.StatusOK, errCommon(ServiceError))

	}
}

/**
正常返回
*/
func HandleSuccess(c echo.Context, i interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    Success,
		"message": ReturnMsg[Success],
		"data":    i,
	})
}

/**
错误返回
*/
func HandleError(c echo.Context, errcode ErrorNo, errmsg ...string) error {
	co := int64(errcode)
	var code int
	switch {
	case co == 0:
		code = http.StatusOK
	case co < 20000:
		code = http.StatusUnauthorized
	case co < 50000:
		code = http.StatusBadRequest
	case co < 60000:
		code = http.StatusInternalServerError
	default:
		code = http.StatusInternalServerError
	}
	return c.JSON(code, map[string]interface{}{
		"code":    errcode.String(),
		"message": ReturnMsg[errcode] + strings.Join(errmsg, " "),
	})
}
