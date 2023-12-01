package baidu

import "net/url"

// TradePolymerPayment 百度收银台 https://smartprogram.baidu.com/docs/develop/api/open/payment_swan-requestPolymerPayment/#orderInfo
func (c *Client) TradePolymerPayment(param TradePolymerPayment) (result TradePolymerPaymentRsp, err error) {
	result.DealId = c.dealId
	result.AppKey = c.appKey
	// 签名字段
	data := url.Values{}
	data.Set("appKey", result.AppKey)
	data.Set("dealId", result.DealId)
	data.Set("tpOrderId", param.TpOrderId)
	// signFieldsRange 的值：0：原验签字段 appKey+dealId+tpOrderId；1：包含 totalAmount 的验签，验签字段包括appKey+dealId+tpOrderId+totalAmount。固定值为 1
	if param.SignFieldsRange == "" || param.SignFieldsRange == "1" {
		param.SignFieldsRange = "1"
		data.Set("totalAmount", param.TotalAmount)
	} else if param.SignFieldsRange != "0" {
		param.SignFieldsRange = "0"
	}
	result.TradePolymerPayment = param
	// 进行 RSA 加密后的签名，防止订单被伪造
	if result.RsaSign, err = c.sign(data); err != nil {
		return
	}
	return
}

// TradeOrderQuery 查询订单 https://smartprogram.baidu.com/docs/third/pay/get_by_tp_order_id/
// GET https://openapi.baidu.com/rest/2.0/smartapp/pay/paymentservice/tp/findByTpOrderId
func (c *Client) TradeOrderQuery(param TradeOrderQuery) (result *TradeOrderQueryRsp, err error) {
	err = c.doRequest("GET", param, &result)
	return
}

// TradeCloseOrder 关闭订单 https://smartprogram.baidu.com/docs/third/pay/close_order/
// GET https://openapi.baidu.com/rest/2.0/smartapp/pay/paymentservice/tp/cancelOrder
func (c *Client) TradeCloseOrder(param TradeCloseOrder) (result *TradeCloseOrderRsp, err error) {
	err = c.doRequest("GET", param, &result)
	return
}

// TradeRefund 申请退款 https://smartprogram.baidu.com/docs/third/pay/apply_order_refund/
// POST https://openapi.baidu.com/rest/2.0/smartapp/pay/paymentservice/tp/applyOrderRefund
func (c *Client) TradeRefund(param TradeRefund) (result *TradeRefundRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}

// TradeRefundQuery 查询退款 https://smartprogram.baidu.com/docs/third/pay/get_order_refund/
// GET https://openapi.baidu.com/rest/2.0/smartapp/pay/paymentservice/tp/findOrderRefund
func (c *Client) TradeRefundQuery(param TradeRefundQuery) (result *TradeRefundQueryRsp, err error) {
	err = c.doRequest("GET", param, &result)
	return
}
