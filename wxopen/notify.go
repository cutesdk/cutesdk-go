package wxopen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cutesdk/cutesdk-go/common/crypt"
	"github.com/cutesdk/cutesdk-go/common/request"
)

// NotifyMsg: notify message type
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

// Listen: listen notify
func (ins *Instance) Listen(req *http.Request, resp http.ResponseWriter, notifyHandler NotifyHandler) error {
	msg, err := ins.GetNotifyMsg(req)
	if err != nil {
		return fmt.Errorf("get notify message failed: %v", err)
	}

	if notifyHandler == nil {
		return nil
	}

	replyMsg := notifyHandler(msg)

	if replyMsg == nil {
		return ins.ReplySuccess(resp)
	}

	return ins.ReplySuccess(resp)
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
func (ins *Instance) VerifyNotifyMsg(req *http.Request, msgEncrypt string) error {
	params := req.URL.Query()

	timestamp := params.Get("timestamp")
	nonce := params.Get("nonce")
	msgSignature := params.Get("msg_signature")

	calSign := crypt.GenMsgSignature(ins.opts.VerifyToken, timestamp, nonce, msgEncrypt)
	if calSign != msgSignature {
		return fmt.Errorf("invalid signature")
	}

	return nil
}

// GetNotifyMsg: get notify message
func (ins *Instance) GetNotifyMsg(req *http.Request) (*NotifyMsg, error) {
	reqBody, err := ins.GetReqBody(req)
	if err != nil {
		return nil, err
	}

	msg := request.NewResult(reqBody)
	msg.XmlParsed()

	msgEncrypt := msg.GetString("Encrypt")
	if msgEncrypt == "" {
		return nil, fmt.Errorf("invalid msg_encrypt")
	}

	// verify notify message
	if err := ins.VerifyNotifyMsg(req, msgEncrypt); err != nil {
		return nil, fmt.Errorf("verify notify msg failed: %v", err)
	}

	fmt.Printf("%s,%s", ins.opts.aesKey, msgEncrypt)

	// decrypt message
	contentB, receiveId, err := crypt.DecryptWithAesKey(ins.opts.aesKey, msgEncrypt)
	if err != nil {
		return nil, fmt.Errorf("decrypt message failed: %v", err)
	}

	res := request.NewResult(contentB)
	res.XmlParsed()

	notifyMsg := &NotifyMsg{
		res,
		receiveId,
	}

	return notifyMsg, err
}
