package wxpay

import (
	"fmt"
	"sort"
	"strings"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/idoubi/goutils"
)

// UnifiedOrder: unifiedorder pay
func (c *Client) UnifiedOrder(params map[string]interface{}) (request.Result, error) {
	uri := "/pay/unifiedorder"

	data, err := c.BuildParams(params)
	if err != nil {
		return nil, fmt.Errorf("build params failed: %v", err)
	}

	res, err := c.Post(uri, data)

	return res, err
}

// PayParams: pay params
type PayParams struct {
	Appid     string `json:"appId"`
	Timestamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

// GetPayParams: get pay params
func (c *Client) GetPayParams(params map[string]interface{}) (*PayParams, error) {
	if params == nil {
		return nil, fmt.Errorf("invalid params")
	}

	prepayId := ""
	signType := "MD5"

	if v, ok := params["prepay_id"]; ok {
		prepayId = v.(string)
	}
	if v, ok := params["sign_type"]; ok {
		signType = v.(string)
	}

	if prepayId == "" {
		return nil, fmt.Errorf("invalid prepay_id")
	}
	if signType != "MD5" && signType != "HMAC-SHA256" {
		return nil, fmt.Errorf("invalid sign_type")
	}

	appid := c.GetAppid()
	nonce := goutils.NonceStr(32)
	timestamp := goutils.TimestampStr()
	_package := fmt.Sprintf("prepay_id=%s", prepayId)

	payParams := map[string]interface{}{
		"appId":     appid,
		"nonceStr":  nonce,
		"timeStamp": timestamp,
		"package":   _package,
		"signType":  signType,
	}

	var sign string

	if signType == "HMAC-SHA256" {
		sign = c.SignWithHmacSha256(payParams)
	} else {
		sign = c.SignWithMd5(payParams)
	}

	payParams["paySign"] = sign

	return &PayParams{
		Appid:     appid,
		NonceStr:  nonce,
		Timestamp: timestamp,
		Package:   _package,
		SignType:  signType,
		PaySign:   sign,
	}, nil
}

// BuildParams: build params with sign
func (c *Client) BuildParams(params map[string]interface{}) (map[string]interface{}, error) {
	if params == nil {
		return nil, fmt.Errorf("invalid params")
	}

	signType := "MD5"
	if v, ok := params["sign_type"]; ok {
		signType = v.(string)
	}
	if signType != "MD5" && signType != "HMAC-SHA256" {
		return nil, fmt.Errorf("invalid sign_type")
	}

	params["mch_id"] = c.GetMchId()
	params["appid"] = c.GetAppid()
	params["nonce_str"] = goutils.NonceStr(32)
	params["sign_type"] = signType

	if signType == "HMAC-SHA256" {
		params["sign"] = c.SignWithHmacSha256(params)
	} else {
		params["sign"] = c.SignWithMd5(params)
	}

	return params, nil
}

func (c *Client) SignWithHmacSha256(params map[string]interface{}) string {
	signStr := c.buildSignStr(params)
	sign := strings.ToUpper(goutils.HmacSha256(signStr, c.GetApiKey()))

	return sign
}

func (c *Client) SignWithMd5(params map[string]interface{}) string {
	signStr := c.buildSignStr(params)
	sign := strings.ToUpper(goutils.MD5(signStr))

	return sign
}

func (c *Client) buildSignStr(params map[string]interface{}) string {
	keys := []string{}
	for k := range params {
		keys = append(keys, k)
	}

	// sort param keys
	sort.Strings(keys)

	arr := []string{}
	str := ""

	for _, k := range keys {
		if params[k] == nil {
			continue
		}
		if v, ok := params[k].(string); ok && v == "" {
			continue
		}
		arr = append(arr, fmt.Sprintf("%s=%v", k, params[k]))
	}

	str = strings.Join(arr, "&")
	str += fmt.Sprintf("&%s=%s", "key", c.GetApiKey())

	return str
}
