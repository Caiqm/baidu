package baidu

import (
	"testing"
)

// 百度收银台
func TestClient_TradePolymerPayment(t *testing.T) {
	t.Log("========== TradePolymerPayment ==========")
	client.LoadOptionFunc(WithPayParams("", ""))
	client.LoadAppPrivateKey("")
	var p TradePolymerPayment
	p.TotalAmount = "1"
	p.TpOrderId = "TS13245678997546546"
	p.NotifyUrl = "https://www.com"
	p.DealTitle = "支付测试"
	p.SignFieldsRange = "1"
	r, err := client.TradePolymerPayment(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", r)
}
