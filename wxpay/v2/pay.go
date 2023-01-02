package wxpay

import (
	"encoding/json"
	"fmt"

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

// String: output as string
func (p *PayParams) String() string {
	b, _ := json.Marshal(p)

	return string(b)
}

// GetPayParams: get pay params
func (cli *Client) GetPayParams(appid, prepayId, signType string) (*PayParams, error) {
	if prepayId == "" {
		return nil, fmt.Errorf("invalid prepay_id")
	}
	if signType != "MD5" && signType != "HMAC-SHA256" {
		return nil, fmt.Errorf("invalid sign_type")
	}

	nonce := goutils.NonceStr(32)
	timestamp := goutils.TimestampStr()
	_package := fmt.Sprintf("prepay_id=%s", prepayId)

	params := map[string]interface{}{
		"appId":     appid,
		"nonceStr":  nonce,
		"timeStamp": timestamp,
		"package":   _package,
		"signType":  signType,
	}

	var sign string

	if signType == "HMAC-SHA256" {
		sign = SignWithHmacSha256(params, cli.opts.ApiKey)
	} else {
		sign = SignWithMd5(params, cli.opts.ApiKey)
	}

	payParams := &PayParams{
		Appid:     appid,
		NonceStr:  nonce,
		Timestamp: timestamp,
		Package:   _package,
		SignType:  signType,
		PaySign:   sign,
	}

	return payParams, nil
}
