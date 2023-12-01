package baidu

// TradePolymerPayment 百度收银台 https://smartprogram.baidu.com/docs/develop/api/open/payment_swan-requestPolymerPayment/#orderInfo
type TradePolymerPayment struct {
	TotalAmount     string `json:"totalAmount"`             // 订单金额（单位：人民币分）。注：小程序测试包测试金额不可超过 1000 分
	TpOrderId       string `json:"tpOrderId"`               // 小程序开发者系统创建的唯一订单 ID ，当支付状态发生变化时，会通过此订单 ID 通知开发者
	NotifyUrl       string `json:"notifyUrl"`               // 通知开发者支付状态的回调地址，必须是合法的 URL ，与开发者平台填写的支付回调地址作用一致，未填写的以平台回调地址为准
	DealTitle       string `json:"dealTitle"`               // 订单的名称
	SignFieldsRange string `json:"signFieldsRange"`         // 用于区分验签字段范围，signFieldsRange 的值：0：原验签字段 appKey+dealId+tpOrderId；1：包含 totalAmount 的验签，验签字段包括appKey+dealId+tpOrderId+totalAmount。固定值为 1 。
	BizInfo         string `json:"bizInfo,omitempty"`       // 订单详细信息，需要是一个可解析为 JSON Object 的字符串
	PayResultUrl    string `json:"payResultUrl,omitempty"`  // 当前页面 path。Web 态小程序支付成功后跳回的页面路径，例如：'/pages/payResult/payResult'
	InlinePaySign   string `json:"inlinePaySign,omitempty"` // 内嵌支付组件返回的支付信息加密 key，与 内嵌支付组件配套使用，此值不传或者传空时默认调起支付面板
	PromotionTag    string `json:"promotionTag,omitempty"`  // 平台营销信息，此处传可使用平台券的 spuid，支持通过英文逗号分割传入多个 spuid 值，仅与百度合作平台类目券的开发者需要填写
}

// TradePolymerPaymentRsp 百度收银台响应参数
type TradePolymerPaymentRsp struct {
	DealId  string `json:"dealId"`  // 跳转百度收银台支付必带参数之一，是百度收银台的财务结算凭证，与账号绑定的结算协议一一对应，每笔交易将结算到 dealId 对应的协议主体
	AppKey  string `json:"appKey"`  // 支付能力开通后分配的支付 appKey，用以表示应用身份的唯一 ID ，在应用审核通过后进行分配，一经分配后不会发生更改，来唯一确定一个应用
	RsaSign string `json:"rsaSign"` // 对appKey+dealId+totalAmount+tpOrderId进行 RSA 加密后的签名，防止订单被伪造
	TradePolymerPayment
}

type BizInfo struct {
	TpData TpData `json:"tpData"` // bizInfo 组装键值对集合
}

type TpData struct {
	AppKey      string                 `json:"appKey"`       // 表示应用身份的唯一 ID
	DealId      string                 `json:"dealId"`       // 百度收银台的财务结算凭证
	TpOrderId   string                 `json:"tpOrderId"`    // 业务方唯一订单号
	TotalAmount string                 `json:"totalAmount"`  // 订单总金额（单位：分）
	ReturnData  map[string]interface{} `json:"returnData"`   // 业务方用于透传的业务变量 支付成功后会以 query 形式注入到 payResultUrl 页面中（query 可以在页面的 onLoad 生命周期内获取）
	DisplayData string                 `json:"display_data"` // 收银台定制页面展示属性，非定制业务请置空 用于支付页面展示订单详细信息
}

// TradeOrderQuery 查询订单 https://smartprogram.baidu.com/docs/third/pay/get_by_tp_order_id/
type TradeOrderQuery struct {
	AuxParam
	AccessToken string `json:"access_token"` // 授权小程序的接口调用凭据
	TpOrderId   string `json:"tpOrderId"`    // 开发者订单 ID
	PmAppKey    string `json:"pmAppKey"`     // 调起百度收银台的支付服务 appKey
}

func (aux TradeOrderQuery) NeedSign() bool {
	return false
}

func (aux TradeOrderQuery) NeedAppId() bool {
	return false
}

// TradeOrderQueryRsp 查询订单响应参数
type TradeOrderQueryRsp struct {
	ErrorNo
	Data struct {
		BizInfo       string `json:"bizInfo"`       // 业务扩展字段
		Count         int    `json:"count"`         // 数量
		CreateTime    int    `json:"createTime"`    // 创建时间
		DealId        int    `json:"dealId"`        // 跳转百度收银台支付必带参数之一
		OrderId       int    `json:"orderId"`       // 百度订单ID
		OriPrice      int    `json:"oriPrice"`      // 原价
		ParentOrderId int    `json:"parentOrderId"` // 购物车订单父订单ID
		ParentType    int    `json:"parentType"`    // 订单类型
		PayMoney      int    `json:"payMoney"`      // 支付金额
		SettlePrice   int    `json:"settlePrice"`   // 结算金额
		Status        int    `json:"status"`        // 订单状态 -1订单已取消/订单已异常退款  1订单支付中 2订单已支付
		SubStatus     int    `json:"subStatus"`     // 订单子状态
		TotalMoney    int    `json:"totalMoney"`    // 总金额
		TpId          int    `json:"tpId"`          // tpid
		TpOrderId     string `json:"tpOrderId"`     // 开发者订单ID
		TradeNo       string `json:"tradeNo"`       // 支付单号
		Type          int    `json:"type"`          // ordertype
		OpenId        int    `json:"openId"`        // 小程序用户id
		AppKey        int    `json:"appKey"`        // 小程序appkey
		AppId         int    `json:"appId"`         // 小程序appid
		UserId        int    `json:"userId"`        // 用户 id 与支付状态通知中的保持一致
	} `json:"data"`
}

// TradeCloseOrder 关闭订单 https://smartprogram.baidu.com/docs/third/pay/close_order/
type TradeCloseOrder struct {
	AuxParam
	AccessToken string `json:"access_token"` // 授权小程序的接口调用凭据
	TpOrderId   string `json:"tpOrderId"`    // 开发者订单 ID
	PmAppKey    string `json:"pmAppKey"`     // 调起百度收银台的支付服务 appKey
}

func (aux TradeCloseOrder) NeedSign() bool {
	return false
}

func (aux TradeCloseOrder) NeedAppId() bool {
	return false
}

// TradeCloseOrderRsp 关闭订单响应参数
type TradeCloseOrderRsp struct {
	ErrorNo
	Data bool `json:"data"`
}

// TradeRefund 申请退款 https://smartprogram.baidu.com/docs/third/pay/apply_order_refund/
type TradeRefund struct {
	AuxParam
	AccessToken      string `json:"access_token" form:"access_token"`                   // 授权小程序的接口调用凭据
	ApplyRefundMoney int64  `json:"applyRefundMoney,omitempty" form:"applyRefundMoney"` // 退款金额（单位：分），该字段最大不能超过支付回调中的总金额（totalMoney） 1.如不填金额时，默认整单发起退款 2.含有百度平台营销的订单，目前只支持整单发起退款，不支持部分多次退款
	BizRefundBatchID string `json:"bizRefundBatchId" form:"bizRefundBatchId"`           // 开发者退款批次
	IsSkipAudit      int64  `json:"isSkipAudit" form:"isSkipAudit"`                     // 是否跳过审核，不需要百度请求开发者退款审核请传 1，默认为0； 0：不跳过开发者业务方审核；1：跳过开发者业务方审核。 若不跳过审核，请对接请求业务方退款审核接口
	OrderID          int64  `json:"orderId" form:"orderId"`                             // 百度收银台订单 ID
	RefundReason     string `json:"refundReason" form:"refundReason"`                   // 退款原因
	RefundType       int64  `json:"refundType" form:"refundType"`                       // 退款类型 1：用户发起退款；2：开发者业务方客服退款；3：开发者服务异常退款。
	TpOrderID        string `json:"tpOrderId" form:"tpOrderId"`                         // 开发者订单 ID
	UserID           int64  `json:"userId" form:"userId"`                               // 百度收银台用户 ID
	RefundNotifyURL  string `json:"refundNotifyUrl,omitempty" form:"refundNotifyUrl"`   // 退款通知 url ，不传时默认为在开发者后台配置的 url
	PmAppKey         string `json:"pmAppKey" form:"pmAppKey"`                           // 调起百度收银台的支付服务 appKey
}

func (aux TradeRefund) NeedSign() bool {
	return false
}

func (aux TradeRefund) NeedAppId() bool {
	return false
}

// TradeRefundRsp 申请退款响应参数
type TradeRefundRsp struct {
	ErrorNo
	Data struct {
		RefundBatchId  string `json:"refundBatchId"`  // 平台退款批次号
		RefundPayMoney int    `json:"refundPayMoney"` // 平台可退退款金额【分为单位】
	} `json:"data"`
}

// TradeRefundQuery 查询退款 https://smartprogram.baidu.com/docs/third/pay/get_order_refund/
type TradeRefundQuery struct {
	AuxParam
	AccessToken string `json:"access_token"` // 授权小程序的接口调用凭据
	TpOrderId   string `json:"tpOrderId"`    // 开发者订单 ID
	PmAppKey    string `json:"pmAppKey"`     // 调起百度收银台的支付服务 appKey
	UserId      string `json:"userId"`       // 百度收银台用户 ID
}

func (aux TradeRefundQuery) NeedSign() bool {
	return false
}

func (aux TradeRefundQuery) NeedAppId() bool {
	return false
}

// TradeRefundQueryRsp 查询退款响应参数
type TradeRefundQueryRsp struct {
	ErrorNo
	Data []struct {
		BizRefundBatchId string `json:"bizRefundBatchId"` // 开发者退款批次id
		OrderId          int    `json:"orderId"`          // 退款订单号
		RefundBatchId    int    `json:"refundBatchId"`    // 退款批次id
		RefundStatus     int    `json:"refundStatus"`     // 退款状态 1 退款中 2 退款成功 3 退款失败
		UserId           int    `json:"userId"`           // 退款用户id
	} `json:"data"`
}
