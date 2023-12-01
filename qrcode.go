package baidu

// ShortQrcode 二维码短链 https://smartprogram.baidu.com/docs/develop/serverapi/get_qrcode/
// POST https://openapi.baidu.com/rest/2.0/smartapp/qrcode/getv2?access_token=ACCESS_TOKEN
func (c *Client) ShortQrcode(param ShortQrcode) (result *QrcodeRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// LongQrcode 二维码长链 https://smartprogram.baidu.com/docs/develop/serverapi/get_unlimited_qrcode/
// POST https://openapi.baidu.com/rest/2.0/smartapp/qrcode/getunlimitedv2?access_token=ACCESS_TOKEN
func (c *Client) LongQrcode(param LongQrcode) (result *QrcodeRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}
