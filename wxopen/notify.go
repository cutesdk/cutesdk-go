package wxopen

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/idoubi/goutils"
)

// NotifyInfo: notify info data
type NotifyInfo struct {
	Appid                 string `xml:"AppId"`
	CreateTime            string `xml:"CreateTime"`
	InfoType              string `xml:"InfoType"`
	ComponentVerifyTicket string `xml:"ComponentVerifyTicket,omitempty"`
}

// NotifyMsg: notify msg
type NotifyMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content,omitempty"`
	MsgId        string `xml:"MsgId"`
}

// NotifyData: notify data
type NotifyData struct {
	Appid      string `xml:"AppId,omitempty"`
	ToUserName string `xml:"ToUserName,omitempty"`
	Encrypt    string `xml:"Encrypt"`
}

// NewText: new text reply msg
func (msg *NotifyMsg) NewText(content string) *ReplyMsg {
	return &ReplyMsg{
		ToUserName:   CDATAText(msg.FromUserName),
		FromUserName: CDATAText(msg.ToUserName),
		CreateTime:   goutils.TimestampStr(),
		MsgType:      "text",
		Content:      CDATAText(content),
	}
}

// MsgHandler: notify message handler
type MsgHandler func(msg request.Result)

// Listen: listen notify
func (s *Server) Listen(req *http.Request, resp http.ResponseWriter, msgHandler MsgHandler) error {
	res, err := s.GetNotifyData(req)
	if err != nil {
		return err
	}

	if msgHandler != nil {
		msgHandler(res)
	}

	return nil
}

// GetNotifyInfo: get notify info
func (s *Server) GetNotifyInfo(req *http.Request) (*NotifyInfo, error) {
	res, err := s.GetNotifyData(req)
	if err != nil {
		return nil, fmt.Errorf("get notify data failed: %v", err)
	}

	notifyInfo := &NotifyInfo{}
	if err := xml.Unmarshal(res, &notifyInfo); err != nil {
		return nil, err
	}

	return notifyInfo, nil
}

// GetNotifyMsg: get notify message
func (s *Server) GetNotifyMsg(req *http.Request) (*NotifyMsg, error) {
	res, err := s.GetNotifyData(req)
	if err != nil {
		return nil, fmt.Errorf("get notify data failed: %v", err)
	}

	notifyMsg := &NotifyMsg{}
	if err := xml.Unmarshal(res, &notifyMsg); err != nil {
		return nil, err
	}

	return notifyMsg, nil
}

// GetNotifyData: get notify data
func (s *Server) GetNotifyData(req *http.Request) (request.Result, error) {
	queryParams := req.URL.Query()
	timestamp := queryParams.Get("timestamp")
	nonce := queryParams.Get("nonce")
	msgSignature := queryParams.Get("msg_signature")

	if timestamp == "" || nonce == "" || msgSignature == "" {
		return nil, fmt.Errorf("invalid notify params")
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("invalid notify data: %v", err)
	}
	defer req.Body.Close()

	notifyData := &NotifyData{}
	err = xml.Unmarshal(body, &notifyData)
	if err != nil || notifyData.Encrypt == "" {
		return nil, fmt.Errorf("notify data unmarshal failed")
	}

	if notifyData.Appid != "" && notifyData.Appid != s.GetComponentAppid() {
		return nil, fmt.Errorf("notify data with invalid appid")
	}

	if err := s.VerifyMsg(timestamp, nonce, msgSignature, notifyData.Encrypt); err != nil {
		return nil, fmt.Errorf("notify data verify failed: %v", err)
	}

	res, err := s.DecryptMsg(notifyData.Encrypt)
	if err != nil {
		return nil, fmt.Errorf("notify data decrypt failed: %v", err)
	}

	return request.Result(res), err
}
