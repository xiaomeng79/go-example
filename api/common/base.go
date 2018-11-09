package common

import (
	"encoding/base64"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/json-iterator/go"
	"github.com/xiaomeng79/go-utils/crypto"
	"github.com/xiaomeng79/go-utils/time"
	"math"
	"strings"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type ReqParam struct {
	AppId     int    `json:"-"`                                                    //AppID
	AppKey    string `json:"appKey" valid:"required~appKey必须存在"`                   //密钥ID
	AppSecret string `json:"appSecret" valid:"-"`                                  //密钥
	RequestId string `json:"requestId" valid:"required~required必须存在"`              //32位的唯一请求标识，用于问题排查和防止重复提交
	Timestamp string `json:"timestamp" valid:"required~毫秒时间戳必须存在"`                 //毫秒时间戳
	Custom    string `json:"custom" valid:"-"`                                     //第三方自定义内容
	Nonce     string `json:"nonce" valid:"required~随机数必须存在,length(8|8)~随机数必须满足8位"` //8 位随机数
	Language  string `json:"language" valid:"in(cn|en)"`                           //多语言支持
	Sign      string `json:"sign" valid:"required~签名必须存在"`                         //签名
	SignType  string `json:"signType" valid:"required~签名类型必须存在"`                   //签名类型：MD5 SHA_1 SHA_256 SHA_512
	Encode    bool   `json:"encode" valid:"-"`                                     //响应数据data是否进行base64编码，默认true
	Data      string `json:"data" valid:"-"`                                       //请求的数据
	Remark    string `json:"remark"`                                               //APP备注
	Page      Page   `json:"page"`                                                 //分页
	IsPage    bool   `json:"isPage"`                                               //是否分页
}

//分页
type Page struct {
	PageIndex int64 `json:"pageIndex"` //页面索引
	PageSize  int64 `json:"pageSize"`  //每页大小
	PageTotal int64 `json:"pageTotal"` //总分页数
	Count     int64 `json:"count"`     //当页记录数
	Total     int64 `json:"total"`     //总记录数
}

/**
初始化分页
*/
func (r *ReqParam) InitPage(total int64) {
	r.IsPage = true
	if r.Page.PageIndex <= 0 {
		r.Page.PageIndex = 1
	}
	if r.Page.PageSize <= 0 {
		r.Page.PageSize = 20
	}
	//最大的索引
	maxIndex := int64(math.Ceil(float64(total) / float64(r.Page.PageSize)))
	if maxIndex != 0 && maxIndex < r.Page.PageIndex {
		r.Page.PageIndex = maxIndex
	}
	//总记录数
	r.Page.Total = total
	//总分数页
	if maxIndex <= 0 { //没数据情况
		r.Page.PageTotal = 0
		r.Page.Count = 0
	} else {
		r.Page.PageTotal = maxIndex
	}
	//当页记录数
	if r.Page.PageTotal > r.Page.PageIndex { //总页数大于当前页数
		r.Page.Count = r.Page.PageSize
	}
	if r.Page.PageTotal == r.Page.PageIndex { //总页数等于当前页数
		r.Page.Count = r.Page.Total - (r.Page.PageSize * (r.Page.PageIndex - 1))
	}

}

/**
解析参
*/

func (r *ReqParam) Decode(s string) error {
	return json.Unmarshal([]byte(s), r)
}

/**
验证参数
*/

func (r *ReqParam) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return err
}

/**
**解析data参数
**input: v point
**ouput:  error
 */
func (r *ReqParam) DataDecode(v interface{}) error {
	if r.Data == "" {
		return nil
	}
	//判断参数是否base64编码
	var decoded []byte
	var err error
	if r.Encode { //解码
		decoded, err = base64.StdEncoding.DecodeString(r.Data)
		if err != nil {
			return errors.New("(base64解析失败):" + err.Error())
		}
	} else {
		decoded = []byte(r.Data)
	}
	err = json.Unmarshal(decoded, v)
	if err != nil {
		return errors.New("(json解析失败):" + err.Error())
	}

	return nil
}

/**
**编码data参数
**input: v interface
**ouput:  error
 */
func (r *ReqParam) DataEncode(v interface{}) error {
	if v == "" {
		r.Data = ""
		return nil
	}
	var encoded []byte
	var err error
	//json marshal
	encoded, err = json.Marshal(v)
	if err != nil {
		return errors.New("json编码失败")
	}

	//判断参数是否base64编码
	if r.Encode { //解码
		r.Data = base64.StdEncoding.EncodeToString(encoded)
	} else {
		r.Data = string(encoded)
	}

	return nil
}

/**
生成签名 string
*/
func (r *ReqParam) CreateSign() (string, error) {
	_signType := strings.ToUpper(strings.Trim(r.SignType, " "))
	var originSign string
	//组合签名字符串
	if r.Data == "" {
		originSign = r.AppKey + r.AppSecret + r.Nonce + r.Timestamp
	} else {
		originSign = r.AppKey + r.AppSecret + r.Data + r.Nonce + r.Timestamp
	}
	var _sign string
	var err error
	switch _signType {
	case "MD5":
		_sign = crypto.MD5(originSign)
	case "SHA_1":
		_sign = crypto.SHA1(originSign)
	case "SHA_256":
		_sign = crypto.SHA256(originSign)
	case "SHA_512":
		_sign = crypto.SHA512(originSign)
	default:
		err = errors.New("签名类型不存在")
	}
	if err != nil {
		return "", err
	}
	return _sign, nil
}

/**
根据签名类型比较签名 true:相同 false：不同
*/
func (r *ReqParam) CompareSign() (bool, error) {
	_sign, err := r.CreateSign()
	if err != nil {
		return false, err
	}
	if strings.ToLower(r.Sign) == _sign {
		return true, nil
	} else {
		return false, errors.New("签名不匹配")
	}
}

/**
生成返回数据
*/
func (r *ReqParam) genData(code ErrorNo, errmsg string, v interface{}) (interface{}, error) {

	var err error
	err = r.DataEncode(v)
	if err != nil {
		return nil, err
	}
	r.Timestamp = time.GenMicTime()
	r.Sign, err = r.CreateSign()
	if err != nil {
		return nil, err
	}
	_v := map[string]interface{}{
		"code":      code.String(),
		"message":   ReturnMsg[code] + errmsg,
		"appKey":    r.AppKey,
		"requestId": r.RequestId,
		"timestamp": r.Timestamp,
		"custom":    r.Custom,
		"sign":      r.Sign,
		"signType":  r.SignType,
		"encode":    r.Encode,
		"data":      r.Data,
	}
	if r.IsPage {
		_v["page"] = r.Page
	}
	return _v, err
}

//返回数据
/*
0.生成data数据和page数据
1.encode data数据
2.生成时间错和签名
3.生成错误吗和错误信息
4.返回数据
*/
func (r *ReqParam) R(v interface{}) (interface{}, error) {
	return r.genData(Success, "", v)
}
