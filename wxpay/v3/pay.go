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
	signType := "RSA"

	if v, ok := params["prepay_id"]; ok {
		prepayId = v.(string)
	}

	if prepayId == "" {
		return nil, fmt.Errorf("invalid prepay_id")
	}

	appid := ins.opts.Appid
	nonce := goutils.NonceStr(32)
	timestamp := goutils.TimestampStr()
	_package := fmt.Sprintf("prepay_id=%s", prepayId)

	str := fmt.Sprintf("%s\n%s\n%s\n%s\n", appid, timestamp, nonce, _package)
	sign, err := ins.payClient.Sign(ins.ctx, str)
	if err != nil {
		return nil, err
	}

	return &PayParams{
		Appid:     appid,
		NonceStr:  nonce,
		Timestamp: timestamp,
		Package:   _package,
		SignType:  signType,
		PaySign:   sign.Signature,
	}, nil
}
