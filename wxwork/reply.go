package wxwork

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/cutesdk/cutesdk-go/common/crypt"
	"github.com/idoubi/goutils"
)

type CDATAText string

func (c CDATAText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

// ReplyMsg: reply message
type ReplyMsg struct {
	XMLName      xml.Name  `xml:"xml"`
	ToUserName   CDATAText `xml:"ToUserName"`
	FromUserName CDATAText `xml:"FromUserName"`
	CreateTime   string    `xml:"CreateTime"`
	MsgType      CDATAText `xml:"MsgType"`
	Content      CDATAText `xml:"Content,omitempty"`
}

// ReplyData: encrypted data reply to user
type ReplyData struct {
	XMLName      xml.Name  `xml:"xml"`
	Encrypt      CDATAText `xml:"Encrypt"`
	MsgSignature CDATAText `xml:"MsgSignature"`
	TimeStamp    string    `xml:"TimeStamp"`
	Nonce        CDATAText `xml:"Nonce"`
}

func (s *Server) EncryptReplyMsg(msg *ReplyMsg) ([]byte, error) {
	xb, err := xml.MarshalIndent(msg, " ", "  ")
	if err != nil {
		return nil, fmt.Errorf("format reply_msg failed: %v", err)
	}

	msgEncrypt, err := crypt.EncryptMsg(s.opts.aesKey, xb, s.opts.Corpid)
	if err != nil {
		return nil, fmt.Errorf("encrypt reply_msg failed: %v", err)
	}

	timestamp := goutils.TimestampStr()
	nonce := goutils.NonceStr(16)
	msgSignature := crypt.GenMsgSignature(s.opts.VerifyToken, timestamp, nonce, msgEncrypt)

	replyData := &ReplyData{
		Encrypt:      CDATAText(msgEncrypt),
		MsgSignature: CDATAText(msgSignature),
		TimeStamp:    timestamp,
		Nonce:        CDATAText(nonce),
	}

	rxb, err := xml.MarshalIndent(replyData, " ", "  ")
	if err != nil {
		return nil, fmt.Errorf("format reply data failed: %v", err)
	}

	return rxb, nil
}

// ReplySuccess 回复字符串success
func (s *Server) ReplySuccess(resp http.ResponseWriter) error {
	_, err := resp.Write([]byte("success"))

	return err
}

// ReplyMessage 回复消息
func (s *Server) ReplyMessage(resp http.ResponseWriter, msg *ReplyMsg) error {
	reply, err := s.EncryptReplyMsg(msg)
	if err != nil {
		return err
	}

	resp.Header().Set("Content-Type", "text/xml")
	_, err = resp.Write(reply)

	return err
}
