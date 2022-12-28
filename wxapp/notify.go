package wxapp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cutesdk/cutesdk-go/common/crypt"
	"github.com/cutesdk/cutesdk-go/common/request"
)

// NotifyMsg: notify message
type NotifyMsg struct {
	*request.Result
	receiveId string
}

// ReceiveId: get receiveId
func (n *NotifyMsg) ReceiveId() string {
	return n.receiveId
}

// NotifyHandler: notify handler
type NotifyHandler func(*NotifyMsg) *ReplyMsg

// Listen: listen notify message
func (svr *Server) Listen(req *http.Request, resp http.ResponseWriter, notifyHandler NotifyHandler) error {
	params := req.URL.Query()

	echostr := params.Get("echostr")

	if echostr != "" {
		// verify notify url
		if err := svr.VerifyNotifyMsg(req, echostr); err != nil {
			return fmt.Errorf("verify notify url failed: %v", err)
		}

		resp.Write([]byte(echostr))

		return nil
	}

	msg, err := svr.GetNotifyMsg(req)
	if err != nil {
		return fmt.Errorf("get notify message failed: %v", err)
	}

	if notifyHandler == nil {
		return nil
	}

	replyMsg := notifyHandler(msg)

	if replyMsg == nil {
		return svr.ReplySuccess(resp)
	}

	return svr.ReplyPlaintext(resp, replyMsg)
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
func (svr *Server) VerifyNotifyMsg(req *http.Request, msgEncrypt string) error {
	params := req.URL.Query()

	timestamp := params.Get("timestamp")
	nonce := params.Get("nonce")
	msgSignature := params.Get("signature")

	calSign := crypt.GenMsgSignature(svr.opts.VerifyToken, timestamp, nonce, "")

	if calSign != msgSignature {
		return fmt.Errorf("invalid signature")
	}

	return nil
}

// GetNotifyMsg: get notify message
func (svr *Server) GetNotifyMsg(req *http.Request) (*NotifyMsg, error) {
	reqBody, err := svr.GetReqBody(req)
	if err != nil {
		return nil, err
	}

	msg := request.NewResult(reqBody)
	msg.XmlParsed()

	msgEncrypt := msg.GetString("Encrypt")

	// verify notify message
	if err := svr.VerifyNotifyMsg(req, ""); err != nil {
		return nil, fmt.Errorf("verify notify msg failed: %v", err)
	}

	// decrypt message
	if msgEncrypt != "" {
		contentB, receiveId, err := crypt.DecryptWithAesKey(svr.opts.aesKey, msgEncrypt)
		if err != nil {
			return nil, fmt.Errorf("decrypt message failed: %v", err)
		}

		res := request.NewResult(contentB)
		res.XmlParsed()

		notifyMsg := &NotifyMsg{res, receiveId}

		return notifyMsg, nil
	}

	notifyMsg := &NotifyMsg{msg, ""}

	return notifyMsg, nil
}
