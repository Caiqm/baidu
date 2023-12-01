package baidu

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"hash"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	appId          string
	secret         string
	appKey         string
	dealId         string
	secretKey      string
	publicKey      string
	privateKey     string
	host           string
	signType       string
	location       *time.Location
	client         *http.Client
	onReceivedData func(method string, data []byte)
}

type OptionFunc func(c *Client)

// 设置请求连接
func WithApiHost(host string) OptionFunc {
	return func(c *Client) {
		if host != "" {
			c.host = host
		}
	}
}

// 设置加密方式
func WithSignType(signType string) OptionFunc {
	return func(c *Client) {
		if signType != "" {
			c.signType = signType
		}
	}
}

// 设置支付参数
func WithPayParams(appKey, dealId string) OptionFunc {
	return func(c *Client) {
		if appKey != "" {
			c.appKey = appKey
		}
		if dealId != "" {
			c.dealId = dealId
		}
	}
}

// 初始化
func New(appId, secret string, opts ...OptionFunc) (nClient *Client, err error) {
	if appId == "" || secret == "" {
		return nil, errors.New("bad params")
	}
	nClient = &Client{}
	nClient.appId = appId
	nClient.secret = secret
	nClient.signType = "RSA"
	nClient.client = http.DefaultClient
	nClient.location = time.Local
	nClient.LoadOptionFunc(opts...)
	return
}

// 加载接口链接
func (c *Client) LoadOptionFunc(opts ...OptionFunc) {
	for _, opt := range opts {
		if opt != nil {
			opt(c)
		}
	}
}

// 加载公钥
func (c *Client) LoadAppPublicKey(pubKey string) {
	if pubKey != "" {
		c.publicKey = pubKey
	}
}

// 加载私钥
func (c *Client) LoadAppPrivateKey(priKey string) {
	if priKey != "" {
		c.privateKey = priKey
	}
}

// 结构体转map
func (c *Client) structToMap(stu interface{}) map[string]interface{} {
	// 结构体转map
	m, _ := json.Marshal(&stu)
	var parameters map[string]interface{}
	_ = json.Unmarshal(m, &parameters)
	return parameters
}

// 请求参数
func (c *Client) URLValues(param Param) (value url.Values, err error) {
	var values = url.Values{}
	// 是否需要APPID
	if param.NeedAppId() {
		values.Add(kFieldAppId, c.appId)
	}
	// 是否需要密钥
	if param.NeedSecret() {
		values.Add(kFieldSecret, c.secret)
	}
	// 结构体转MAP
	var params = c.structToMap(param)
	for k, v := range params {
		if v == "" {
			continue
		}
		values.Add(k, v.(string))
	}
	// 判断是否需要签名
	if param.NeedSign() {
		var signature string
		if signature, err = c.sign(values); err != nil {
			return
		}
		// 添加签名
		values.Add(kFieldSign, signature)
	}
	return values, nil
}

// 生成签名
func (c *Client) sign(data url.Values) (signature string, err error) {
	signContentBytes, _ := url.QueryUnescape(data.Encode())
	var h hash.Hash
	var hType crypto.Hash
	switch c.signType {
	case "RSA":
		h = sha1.New()
		hType = crypto.SHA1
	case "RSA2":
		h = sha256.New()
		hType = crypto.SHA256
	}
	h.Write([]byte(signContentBytes))
	d := h.Sum(nil)
	pk, err := c.parsePrivateKey(c.privateKey)
	if err != nil {
		return
	}
	bs, err := rsa.SignPKCS1v15(rand.Reader, pk, hType, d)
	if err != nil {
		return
	}
	signature = base64.StdEncoding.EncodeToString(bs)
	return
}

// 处理密钥
func (c *Client) parsePrivateKey(prvKey string) (pk *rsa.PrivateKey, err error) {
	keyByts, _ := base64.StdEncoding.DecodeString(prvKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(keyByts)
	if err != nil {
		privateKey, err = x509.ParsePKCS1PrivateKey(keyByts)
		if err != nil {
			err = fmt.Errorf("解析密钥失败，err: %v", err)
			return
		}
	}
	pk = privateKey.(*rsa.PrivateKey)
	return
}

// 请求接口
func (c *Client) doRequest(method string, param Param, result interface{}) (err error) {
	// 创建一个请求
	req, _ := http.NewRequest(method, c.host, nil)
	// 判断参数是否为空
	if param != nil {
		var values url.Values
		values, err = c.URLValues(param)
		if err != nil {
			return err
		}
		if method == http.MethodPost {
			req.PostForm = values
		} else if method == http.MethodGet {
			req.URL, _ = url.Parse(c.host + "?" + values.Encode())
		}
	}
	// 添加header头
	req.Header.Add("Content-Type", kContentType)
	// 发起请求数据
	rsp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	err = c.decode(bodyBytes, method, result)
	return
}

// 解密返回数据
func (c *Client) decode(data []byte, method string, result interface{}) (err error) {
	// 返回结果
	if c.onReceivedData != nil {
		c.onReceivedData(method, data)
	}
	var raw = make(map[string]json.RawMessage)
	if err = json.Unmarshal(data, &raw); err != nil {
		return
	}
	var errTBytes = raw[kFieldErr]
	var errCBytes = raw[kFieldErrCode]
	var errNBytes = raw[kFieldErrNo]
	if len(errTBytes) > 0 {
		var tErr Err
		if err = json.Unmarshal(data, &tErr); err != nil {
			return err
		}
		return tErr
	}
	if len(errCBytes) > 0 {
		var cErr ErrorCode
		if err = json.Unmarshal(data, &cErr); err != nil {
			return err
		}
		return cErr
	} else if len(errCBytes) == 0 && len(errNBytes) > 0 {
		var nErr ErrorNo
		if err = json.Unmarshal(data, &nErr); err != nil {
			return err
		}
		return nErr
	}
	// 返回数据
	if err = json.Unmarshal(data, result); err != nil {
		return err
	}
	return nil
}

// 验证签名
func (c *Client) VerifySign(signContent, sign string) (checkRes bool, err error) {
	// 步骤1，加载RSA的公钥
	publicKey := c.formatPublicKey(c.publicKey)
	block, _ := pem.Decode([]byte(publicKey))
	// keyByts, _ := base64.StdEncoding.DecodeString(publicKey)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	checkRes = false
	if err != nil {
		return
	}
	rsaPub, _ := pub.(*rsa.PublicKey)
	// 步骤2，计算代签名字串的SHA1哈希
	hashed := sha1.Sum([]byte(signContent))
	hashs := crypto.SHA1
	digest := hashed[:]
	// 步骤3，base64 decode,必须步骤，支付宝对返回的签名做过base64 encode必须要反过来decode才能通过验证
	data, _ := base64.StdEncoding.DecodeString(sign)
	// 步骤4，调用rsa包的VerifyPKCS1v15验证签名有效性
	err = rsa.VerifyPKCS1v15(rsaPub, hashs, digest, data)
	if err != nil {
		fmt.Println("Verify sig error, reason: ", err)
		return
	}
	checkRes = true
	return
}

// 格式化普通支付宝公钥
func (c *Client) formatPublicKey(publicKey string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString("-----BEGIN PUBLIC KEY-----\n")
	rawLen := 64
	keyLen := len(publicKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(publicKey[start:])
		} else {
			buffer.WriteString(publicKey[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END PUBLIC KEY-----\n")
	pKey = buffer.String()
	return
}

// 返回内容
func (c *Client) OnReceivedData(fn func(method string, data []byte)) {
	c.onReceivedData = fn
}
