package wxpay

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/tidwall/gjson"
)

// NotifyMsg: notify message type
type NotifyMsg struct {
	raw []byte
	res gjson.Result
}

// Raw: get raw data
func (n *NotifyMsg) Raw() []byte {
	return n.raw
}

// String: print format
func (n *NotifyMsg) String() string {
	return string(n.raw)
}

// GetString: get field value with string format
func (n *NotifyMsg) GetString(key string) string {
	return n.res.Get(key).String()
}

// GetInt: get field value with int64 format
func (n *NotifyMsg) GetInt(key string) int64 {
	return n.res.Get(key).Int()
}

// NotifyHandler: notify handler
type NotifyHandler func(*NotifyMsg) *ReplyMsg

// Listen: listen notify
func (s *Server) Listen(req *http.Request, resp http.ResponseWriter, notifyHandler NotifyHandler) error {
	msg, err := s.GetNotifyMsg(req)
	if err != nil {
		return fmt.Errorf("get notify message failed: %v", err)
	}

	if msg.GetString("return_code") == "SUCCESS" {
		if err := s.VerifyNotifyMsg(msg); err != nil {
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

	return s.ReplyMessage(resp, replyMsg)
}

// GetReqBody: get request data in body
func (s *Server) GetReqBody(req *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("invalid notify data: %v", err)
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}

// GetNotifyMsg: get notify message
func (s *Server) GetNotifyMsg(req *http.Request) (*NotifyMsg, error) {
	reqBody, err := s.GetReqBody(req)
	if err != nil {
		return nil, err
	}

	notifyMsg := &NotifyMsg{
		raw: reqBody,
		res: request.Result(reqBody).XmlParsed(),
	}

	return notifyMsg, err
}

// VerifyNotifyMsg: verify notify message
func (s *Server) VerifyNotifyMsg(msg *NotifyMsg) error {
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
	for k, v := range msg.res.Map() {
		params[k] = v
	}

	if signType == "MD5" {
		if calSign := SignWithMd5(params, s.opts.ApiKey); calSign != sign {
			return fmt.Errorf("invalid sign")
		}
	}

	if signType == "HMAC-SHA256" {
		if sign != SignWithHmacSha256(params, s.opts.ApiKey) {
			return fmt.Errorf("invalid sign")
		}
	}

	return nil
}
