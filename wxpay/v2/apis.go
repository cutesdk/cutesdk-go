package wxpay

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/idoubi/goutils"
)

// PayParams: pay params
type PayParams struct {
	Appid     string `json:"appId"`
	Timestamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

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

	appid := c.opts.Appid
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
		sign = SignWithHmacSha256(payParams, c.opts.ApiKey)
	} else {
		sign = SignWithMd5(payParams, c.opts.ApiKey)
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

// QueryOrder: query order
func (c *Client) QueryOrder(params map[string]interface{}) (request.Result, error) {
	uri := "/pay/orderquery"

	data, err := c.BuildParams(params)
	if err != nil {
		return nil, fmt.Errorf("build params failed: %v", err)
	}

	res, err := c.Post(uri, data)

	return res, err
}

// QueryOrderByOutTradeNo: query order by out_trade_no
func (c *Client) QueryOrderByOutTradeNo(outTradeNo string) (request.Result, error) {
	params := map[string]interface{}{
		"out_trade_no": outTradeNo,
	}

	return c.QueryOrder(params)
}

// QueryOrderByTransactionId: query order by transaction_id
func (c *Client) QueryOrderByTransactionId(transactionId string) (request.Result, error) {
	params := map[string]interface{}{
		"transaction_id": transactionId,
	}

	return c.QueryOrder(params)
}
