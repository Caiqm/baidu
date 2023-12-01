package baidu

// GetAccessToken 获取小程序全局唯一后台接口调用凭据（access_token）https://smartprogram.baidu.com/docs/develop/serverapi/power_exp/
// GET https://openapi.baidu.com/oauth/2.0/token?grant_type=client_credentials&client_id=CLIENT_ID&client_secret=CLIENT_SECRET&scope=smartapp_snsapi_base
func (c *Client) GetAccessToken(param GetAccessToken) (result *GetAccessTokenRsp, err error) {
	if param.Scope == "" {
		param.Scope = "smartapp_snsapi_base"
	}
	if param.GrantType == "" {
		param.GrantType = "client_credentials"
	}
	err = c.doRequest("GET", param, &result)
	return
}

// SessionKey 用户登陆凭证 https://smartprogram.baidu.com/docs/develop/api/open/getSessionKey/
// GET https://openapi.baidu.com/rest/2.0/smartapp/getsessionkey?access_token=ACCESS_TOKEN&code=CODE
func (c *Client) SessionKey(param SessionKey) (result *SessionKeyRsp, err error) {
	err = c.doRequest("GET", param, &result)
	return
}

// GetUnionId 获取unionID https://smartprogram.baidu.com/docs/develop/api/open/log_getunionid/
// POST https://openapi.baidu.com/rest/2.0/smartapp/getunionid?access_token=ACCESS_TOKEN
func (c *Client) GetUnionId(param GetUnionId) (result *GetUnionIdRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}
