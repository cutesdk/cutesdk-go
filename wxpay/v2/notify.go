package wxpay

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// NotifyMsg: notify message type
type NotifyMsg struct {
	*request.Result
}

// NotifyHandler: notify handler
type NotifyHandler func(*NotifyMsg) *ReplyMsg

// Listen: listen notify
func (ins *Instance) Listen(req *http.Request, resp http.ResponseWriter, notifyHandler NotifyHandler) error {
	msg, err := ins.GetNotifyMsg(req)
	if err != nil {
		return fmt.Errorf("get notify message failed: %v", err)
	}

	if msg.GetString("return_code") == "SUCCESS" {
		if err := ins.VerifyNotifyMsg(msg); err != nil {
			return fmt.Errorf("verify notify message failed: %v", err)
		}
	}

	if notifyHandler == nil {
		return nil
	}

	replyMsg := notifyHandler(msg)

	if replyMsg == nil {
		return nil
	}

	return ins.ReplyMessage(resp, replyMsg)
}

// GetReqBody: get request data in body
func (ins *Instance) GetReqBody(req *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("invalid notify data: %v", err)
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}

// VerifyNotifyMsg: verify notify message
func (ins *Instance) VerifyNotifyMsg(msg *NotifyMsg) error {
	sign := msg.GetString("sign")
	if sign == "" {
		return fmt.Errorf("invalid sign")
	}

	signType := msg.GetString("sign_type")
	if signType == "" {
		signType = "MD5"
	}

	if signType != "MD5" && signType != "HMAC-SHA256" {
		return fmt.Errorf("invalid sign_type")
	}

	params := map[string]interface{}{}
	for k, v := range msg.Map() {
		params[k] = v
	}

	if signType == "MD5" {
		if calSign := SignWithMd5(params, ins.opts.ApiKey); calSign != sign {
			return fmt.Errorf("invalid sign")
		}
	}

	if signType == "HMAC-SHA256" {
		if sign != SignWithHmacSha256(params, ins.opts.ApiKey) {
			return fmt.Errorf("invalid sign")
		}
	}

	return nil
}

// GetNotifyMsg: get notify message
func (ins *Instance) GetNotifyMsg(req *http.Request) (*NotifyMsg, error) {
	reqBody, err := ins.GetReqBody(req)
	if err != nil {
		return nil, err
	}

	res := request.NewResult(reqBody)
	res.XmlParsed()

	notifyMsg := &NotifyMsg{
		res,
	}

	return notifyMsg, err
}
