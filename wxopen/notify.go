package wxopen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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
func (svr *Server) Listen(req *http.Request, resp http.ResponseWriter, notifyHandler NotifyHandler) error {
	msg, err := svr.GetNotifyMsg(req)
	if err != nil {
		return fmt.Errorf("get notify message failed: %v", err)
	}

	if svr.opts.Appid == "" && msg.receiveId != "" {
		svr.opts.Appid = msg.receiveId
	}

	infoType := msg.GetString("InfoType")

	// auto cache component_verify_ticket
	if infoType == "component_verify_ticket" {
		ticket := msg.GetString("ComponentVerifyTicket")
		svr.cli.SetVerifyTicket(ticket, 12*time.Hour)
	}

	if notifyHandler == nil {
		return nil
	}

	replyMsg := notifyHandler(msg)

	if replyMsg == nil {
		return svr.ReplySuccess(resp)
	}

	return svr.ReplyEncryptedMsg(resp, replyMsg)
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
	msgSignature := params.Get("msg_signature")

	calSign := crypt.GenMsgSignature(svr.opts.VerifyToken, timestamp, nonce, msgEncrypt)
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
	if msgEncrypt == "" {
		return nil, fmt.Errorf("invalid msg_encrypt")
	}

	// verify notify message
	if err := svr.VerifyNotifyMsg(req, msgEncrypt); err != nil {
		return nil, fmt.Errorf("verify notify msg failed: %v", err)
	}

	// decrypt message
	contentB, receiveId, err := crypt.DecryptWithAesKey(svr.opts.aesKey, msgEncrypt)
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
