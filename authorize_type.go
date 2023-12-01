package baidu

// GetAccessToken 获取小程序全局唯一后台接口调用凭据（access_token）https://smartprogram.baidu.com/docs/develop/serverapi/power_exp/
type GetAccessToken struct {
	AuxParam
	GrantType string `json:"grant_type"` // 固定为：client_credentials
	Scope     string `json:"scope"`      // 固定为：smartapp_snsapi_base
}

func (a GetAccessToken) NeedSign() bool {
	return false
}

func (a GetAccessToken) NeedSecret() bool {
	return true
}

// GetAccessTokenRsp 获取小程序全局唯一后台接口调用凭据响应参数
type GetAccessTokenRsp struct {
	Err
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒
}

// SessionKey 用户登陆凭证 https://smartprogram.baidu.com/docs/develop/api/open/getSessionKey/
type SessionKey struct {
	AuxParam
	AccessToken string `json:"access_token"` // 接口调用凭证
	Code        string `json:"code"`         // 通过 swan.getLoginCode 获取 Authorization Code 特殊说明：code 中有 @ 符号时，会请求对应的开源宿主，用户身份校验及 SessionKey 生成过程由开源宿主实现
}

func (a SessionKey) NeedSign() bool {
	return false
}

func (a SessionKey) NeedAppId() bool {
	return false
}

func (a SessionKey) NeedSecret() bool {
	return false
}

// SessionKeyRsp 用户登陆凭证响应参数
type SessionKeyRsp struct {
	ErrorCode
	Data struct {
		SessionKey string `json:"session_key"` // 用户的 SessionKey
		OpenId     string `json:"open_id"`     // 用户身份标识 不同用户登录同一个小程序获取到的 openid 不同，同一个用户登录不同小程序获取到的 openid 也不同
	} `json:"data"`
	RequestId string `json:"request_id"` // 请求 ID ，标识一次请求
	Timestamp int    `json:"timestamp"`  // 时间戳 1640140013
}

// GetUnionId 获取unionID https://smartprogram.baidu.com/docs/develop/api/open/log_getunionid/
type GetUnionId struct {
	AuxParam
	Openid string `json:"openid"`
}

func (a GetUnionId) NeedSign() bool {
	return false
}

func (a GetUnionId) NeedSecret() bool {
	return true
}

type GetUnionIdRsp struct {
	ErrorNo
	Data struct {
		Unionid string `json:"unionid"` // 小程序用户 + 开发者主体维度唯一的 id
	} `json:"data"`
	RequestId string `json:"request_id"` // 请求 ID ，标识一次请求
	Timestamp int    `json:"timestamp"`  // 时间戳 1640140013
}
