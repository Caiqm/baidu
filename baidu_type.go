package baidu

import "fmt"

const (
	kContentType = "application/x-www-form-urlencoded;charset=utf-8"
	kTimeFormat  = "2006-01-02 15:04:05"
)

const (
	kFieldAppId       = "app_id"
	kFieldSecret      = "secret"
	kFieldAccessToken = "access_token"
	kFieldSign        = "sign"
	kFieldErr         = "error"
	kFieldErrNo       = "errno"
	kFieldErrCode     = "error_code"
)

type Param interface {
	// NeedSign 是否需要签名，有的接口不需要签名，比如：小程序登录与获取手机号接口
	NeedSign() bool

	// NeedAppId 是否需要APPID，有的接口不需要APPID，比如：获取应用授权调用凭证
	NeedAppId() bool

	// NeedSecret 是否需要密钥
	NeedSecret() bool
}

type AuxParam struct {
}

func (aux AuxParam) NeedSign() bool {
	return true
}

func (aux AuxParam) NeedAppId() bool {
	return true
}

func (aux AuxParam) NeedSecret() bool {
	return false
}

// TOKEN错误
type Err struct {
	ErrorS           string `json:"error"`             // 异常提示信息
	ErrorDescription string `json:"error_description"` // 异常情况详细的提示信息
}

func (e Err) Error() string {
	return fmt.Sprintf("%s - %s", e.ErrorS, e.ErrorDescription)
}

// 基本错误
type ErrorNo struct {
	Errno  int    `json:"errno"`
	Errmsg string `json:"errmsg"`
	Msg    string `json:"msg"`
}

func (e ErrorNo) Error() string {
	errMsg := fmt.Sprintf("%d", e.Errno)
	if e.Errmsg != "" {
		errMsg = fmt.Sprintf("%s - %s", errMsg, e.Errmsg)
	}
	if e.Msg != "" {
		errMsg = fmt.Sprintf("%s - %s", errMsg, e.Msg)
	}
	return errMsg
}

// openapi错误
type ErrorCode struct {
	ErrorNo
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func (e ErrorCode) Error() string {
	var errMsg string
	if e.ErrorCode > 0 {
		errMsg = fmt.Sprintf("%d - %s", e.ErrorCode, e.ErrorMsg)
	}
	if e.Errno > 0 {
		errMsg = fmt.Sprintf("%s，%d - %s", errMsg, e.ErrorCode, e.ErrorMsg)
	}
	return errMsg
}
