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
	nonce := goutils.NonceStr(32)
	timestamp := goutils.TimestampStr()
	_package := fmt.Sprintf("prepay_id=%s", prepayId)

	str := fmt.Sprintf("%s\n%s\n%s\n%s\n", appid, timestamp, nonce, _package)
	sign, err := cli.payClient.Sign(cli.ctx, str)
	if err != nil {
		return nil, err
	}

	payParams := &PayParams{
		Appid:     appid,
		NonceStr:  nonce,
		Timestamp: timestamp,
		Package:   _package,
		SignType:  signType,
		PaySign:   sign.Signature,
	}

	return payParams, nil
}
