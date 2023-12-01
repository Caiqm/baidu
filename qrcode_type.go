package baidu

// ShortQrcode 二维码短链 https://smartprogram.baidu.com/docs/develop/serverapi/get_qrcode/
// 获取小程序二维码短链，长度固定 35 个字符，适用于需要的码数量较少的业务场景。通过该接口生成的二维码，永久有效，有数量限制。
type ShortQrcode struct {
	AuxParam
	Path    string `json:"path"`     // 扫码进入的小程序页面路径，最大长度 4000 字节，可以为空。
	Width   int    `json:"width"`    // 二维码的宽度（单位：px）。最小 280px，最大 1280px
	Mf      int    `json:"mf"`       // 是否包含二维码内嵌 logo 标识，1001 为不包含，默认包含
	ImgFlag int    `json:"img_flag"` // 返回值选项，默认或传 1 时只返回二维码 base64 编码字符串，传 0 只返回 url
}

func (q ShortQrcode) NeedSign() bool {
	return false
}

func (q ShortQrcode) NeedAppId() bool {
	return false
}

// LongQrcode 二维码长链 https://smartprogram.baidu.com/docs/develop/serverapi/get_unlimited_qrcode/
// 获取小程序二维码长链接，适用于需要的码数量较多的业务场景。通过该接口生成的二维码，永久有效，无数量限制。
type LongQrcode struct {
	AuxParam
	Path    string `json:"path"`     // 扫码进入的小程序页面路径，最大长度 4000 字节，可以为空。
	Width   int    `json:"width"`    // 二维码的宽度（单位：px）。最小 280px，最大 1280px
	Mf      int    `json:"mf"`       // 是否包含二维码内嵌 logo 标识，1001 为不包含，默认包含
	ImgFlag int    `json:"img_flag"` // 返回值选项，默认或传 1 时只返回二维码 base64 编码字符串，传 0 只返回 url
}

func (q LongQrcode) NeedSign() bool {
	return false
}

func (q LongQrcode) NeedAppId() bool {
	return false
}

// QrcodeRsp 二维码短链响应参数
type QrcodeRsp struct {
	ErrorNo
	RequestId string `json:"request_id"` // 请求 ID，标识一次请求
	Timestamp int    `json:"timestamp"`  // 时间戳
	Data      struct {
		Base64Str string `json:"base64_str,omitempty"` // 二维码 base64 编码字符串
		Url       string `json:"url,omitempty"`        // 小程序二维码短链
	} `json:"data"`
}
