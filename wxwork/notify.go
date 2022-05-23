package wxwork

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cutesdk/cutesdk-go/common/crypt"
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/tidwall/gjson"
)

// NotifyMsg: notify message type
type NotifyMsg struct {
	receiveId string
	raw       []byte
	res       gjson.Result
}

// ReceiveId: get receiveId
func (n *NotifyMsg) ReceiveId() string {
	return n.receiveId
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
	params := req.URL.Query()

	echostr := params.Get("echostr")
	timestamp := params.Get("timestamp")
	nonce := params.Get("nonce")
	msgSignature := params.Get("msg_signature")

	if echostr != "" {
		// verify notify url
		calSign := crypt.GenMsgSignature(s.opts.VerifyToken, timestamp, nonce, echostr)
		if calSign != msgSignature {
			return fmt.Errorf("invalid signature")
		}

		echoStrB, receiveId, err := crypt.DecryptWithAesKey(s.opts.aesKey, echostr)
		if err != nil {
			return fmt.Errorf("decrypt echostr failed: %v", err)
		}

		// todo check receiveId
		_ = receiveId

		resp.Write(echoStrB)

		return nil
	}

	msg, err := s.GetNotifyMsg(req)
	if err != nil {
		return fmt.Errorf("get notify message failed: %v", err)
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

	msg := request.Result(reqBody).XmlParsed()
	msgEncrypt := msg.Get("Encrypt").String()
	if msgEncrypt == "" {
		return nil, fmt.Errorf("invalid msg_encrypt")
	}

	// verify message
	params := req.URL.Query()
	timestamp := params.Get("timestamp")
	nonce := params.Get("nonce")
	msgSignature := params.Get("msg_signature")

	calSign := crypt.GenMsgSignature(s.opts.VerifyToken, timestamp, nonce, msgEncrypt)
	if calSign != msgSignature {
		return nil, fmt.Errorf("invalid signature")
	}

	// decrypt message
	contentB, receiveId, err := crypt.DecryptWithAesKey(s.opts.aesKey, msgEncrypt)
	if err != nil {
		return nil, fmt.Errorf("decrypt message failed: %v", err)
	}

	notifyMsg := &NotifyMsg{
		receiveId: receiveId,
		raw:       contentB,
		res:       request.Result(contentB).XmlParsed(),
	}

	return notifyMsg, err
}
