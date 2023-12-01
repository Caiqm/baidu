# 百度相关接口（简易版）
百度小程序支付、登录、生成二维码等相关接口

## 安装

#### 启用 Go module

```go
go get github.com/Caiqm/baidu
```

```go
import "github.com/Caiqm/baidu"
```

#### 未启用 Go module

```go
go get github.com/Caiqm/baidu
```

```go
import "github.com/Caiqm/baidu"
```

## 如何使用

```go
// 实例化百度接口
var client, err = baidu.New(appID, Secret)
```

#### 关于密钥（Secret）

密钥是百度应用的唯一凭证密钥，即 AppSecret，获取方式同 appid

## 加载支持配置

```go
// 加载接口链接，支付参数AppKey与DealId
client.LoadOptionFunc(WithApiHost(HOST), WithPayParams(APPKEY, DEALID))

// 或者一开始就加载支付参数AppKey与DealId
var client, err = wxpay.New(appID, Secret, WithPayParams(APPKEY, DEALID))
```

#### 关于加载支持配置
系统内置了几种可支持的配置

```go
// 设置请求链接，可自定义请求接口，传入host字符串
WithApiHost(HOST)

// 设置加密方式
WithSignType(signType)

// 设置支付参数
WithPayParams(appKey, dealId)

// 也可自定义传入配置，返回以下类型即可
type OptionFunc func(c *Client)
```

## 用户登录凭证

```go
// 用户登陆凭证
func TestClient_SessionKey(t *testing.T) {
	t.Log("========== SessionKey ==========")
	client.LoadOptionFunc(WithApiHost("https://openapi.com/rest/2.0/smartapp/getsessionkey"))
	var p SessionKey
	p.Code = "123456"
	p.AccessToken = "123456123456123456123456123456123456123456123456123456"
	r, err := client.SessionKey(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
```

## 百度收银台

```go
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
```