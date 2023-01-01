package wxpay

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/idoubi/goutils"
	"github.com/idoubi/goutils/crypt"
)

// NotifyMsg: notify message type
type NotifyMsg struct {
	*request.Result
}

// NotifyHandler: notify handler
type NotifyHandler func(*NotifyMsg) *ReplyMsg

// Listen: listen notify
func (svr *Server) Listen(req *http.Request, resp http.ResponseWriter, notifyHandler NotifyHandler) error {
	msg, err := svr.GetNotifyMsg(req)
	if err != nil {
		return fmt.Errorf("get notify message failed: %v", err)
	}

	if msg.GetString("return_code") == "SUCCESS" {
		if reqInfo := msg.GetString("req_info"); reqInfo != "" {
			// decrypt req_info
			decryptMsg, err := svr.DecryptReqInfo(reqInfo)
			if err != nil {
				return fmt.Errorf("decrypt req_info failed: %v", err)
			}
			msg = decryptMsg
		} else {
			if err := svr.VerifyNotifyMsg(msg); err != nil {
				return fmt.Errorf("verify notify message failed: %v", err)
			}
		}
	}

	if notifyHandler == nil {
		return nil
	}

	replyMsg := notifyHandler(msg)

	if replyMsg == nil {
		return nil
	}

	return svr.ReplyMessage(resp, replyMsg)
}

// GetReqBody: get request data in body
func (svr *Server) GetReqBody(req *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("invalid notify data: %v", err)
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}

// VerifyNotifyMsg: verify notify message
func (svr *Server) VerifyNotifyMsg(msg *NotifyMsg) error {
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
		if calSign := SignWithMd5(params, svr.opts.ApiKey); calSign != sign {
			return fmt.Errorf("invalid sign")
		}
	}

	if signType == "HMAC-SHA256" {
		if sign != SignWithHmacSha256(params, svr.opts.ApiKey) {
			return fmt.Errorf("invalid sign")
		}
	}

	return nil
}

// DecryptReqInfo: decrypt req_info
func (svr *Server) DecryptReqInfo(reqInfo string) (*NotifyMsg, error) {
	reqInfoB := goutils.Base64Decode(reqInfo)
	if reqInfoB == "" {
		return nil, fmt.Errorf("invalid req_info: %s", reqInfo)
	}

	md5ApiKey := goutils.MD5(svr.opts.ApiKey)

	info, err := crypt.AesEcbDecrypt([]byte(reqInfoB), []byte(md5ApiKey))
	if err != nil {
		return nil, err
	}

	beginTag := []byte("<root>")
	endTag := []byte("</root>")
	if !bytes.HasPrefix(info, beginTag) || !bytes.HasSuffix(info, endTag) {
		return nil, fmt.Errorf("decrpyt req_info failed: %s", info)
	}

	info = bytes.Replace(info, beginTag, []byte("<xml>"), 1)
	info = bytes.Replace(info, endTag, []byte("</xml>"), 1)

	res := request.NewResult(info)
	res.XmlParsed()

	notifyMsg := &NotifyMsg{
		res,
	}

	return notifyMsg, nil
}

// GetNotifyMsg: get notify message
func (svr *Server) GetNotifyMsg(req *http.Request) (*NotifyMsg, error) {
	reqBody, err := svr.GetReqBody(req)
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
