package wxpay

import (
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

// GetPayParams: get pay params
func (ins *Instance) GetPayParams(params map[string]interface{}) (*PayParams, error) {
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

	appid := ins.opts.Appid
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
		sign = SignWithHmacSha256(payParams, ins.opts.ApiKey)
	} else {
		sign = SignWithMd5(payParams, ins.opts.ApiKey)
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
