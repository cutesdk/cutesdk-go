package wxapp

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"

	"github.com/cutesdk/cutesdk-go/common/crypt"
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/idoubi/goutils"
	"github.com/idoubi/goutils/convert"
)

// CDATAText: xml data with CDATA
type CDATAText string

// MarshalXML: marshal xml data
func (c CDATAText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

// EncryptedMsg: encrypted message
type EncryptedMsg struct {
	XMLName      xml.Name  `xml:"xml"`
	Encrypt      CDATAText `xml:"Encrypt"`
	MsgSignature CDATAText `xml:"MsgSignature"`
	TimeStamp    string    `xml:"TimeStamp"`
	Nonce        CDATAText `xml:"Nonce"`
}

// ReplyMsg: raw reply message
type ReplyMsg struct {
	*request.Result
}

// Transfer: transfer message to service system
func (msg *NotifyMsg) Transfer() *ReplyMsg {
	reply := map[string]interface{}{
		"MsgType": "transfer_customer_service",
	}

	return msg.Reply(reply)
}

// Reply: reply msg with map data
func (msg *NotifyMsg) Reply(data map[string]interface{}) *ReplyMsg {
	if data == nil {
		return nil
	}

	data["ToUserName"] = msg.GetString("FromUserName")
	data["FromUserName"] = msg.GetString("ToUserName")
	data["CreateTime"] = goutils.Timestamp()

	xb, err := convert.Map2Xml(data)
	if err != nil {
		return nil
	}

	xs := string(xb)
	if strings.Contains(xs, "Articles_child") {
		xs = strings.ReplaceAll(xs, "Articles_child", "item")
		xb = []byte(xs)
	}

	res := request.NewResult(xb)
	res.XmlParsed()

	return &ReplyMsg{res}
}

// EncryptReplyMsg: encrypt ReplyMsg
func (svr *Server) EncryptReplyMsg(msg *ReplyMsg) ([]byte, error) {
	xb := msg.Raw()

	msgEncrypt, err := crypt.EncryptMsg(svr.opts.aesKey, xb, svr.opts.Appid)
	if err != nil {
		return nil, fmt.Errorf("encrypt reply_msg failed: %v", err)
	}

	timestamp := goutils.TimestampStr()
	nonce := goutils.NonceStr(16)
	msgSignature := crypt.GenMsgSignature(svr.opts.Token, timestamp, nonce, msgEncrypt)

	encryptedMsg := &EncryptedMsg{
		Encrypt:      CDATAText(msgEncrypt),
		MsgSignature: CDATAText(msgSignature),
		TimeStamp:    timestamp,
		Nonce:        CDATAText(nonce),
	}

	rxb, err := xml.MarshalIndent(encryptedMsg, " ", "  ")
	if err != nil {
		return nil, fmt.Errorf("format encrypted msg failed: %v", err)
	}

	return rxb, nil
}

// ReplyEncryptedMsg: reply encrypted msg
func (svr *Server) ReplyEncryptedMsg(resp http.ResponseWriter, msg *ReplyMsg) error {
	encryptedMsg, err := svr.EncryptReplyMsg(msg)
	if err != nil {
		return err
	}

	resp.Header().Set("Content-Type", "text/xml")
	_, err = resp.Write(encryptedMsg)

	return err
}

// ReplyPlaintext: reply unencrypted msg
func (svr *Server) ReplyPlaintext(resp http.ResponseWriter, msg *ReplyMsg) error {
	resp.Header().Set("Content-Type", "text/xml")
	_, err := resp.Write(msg.Raw())

	return err
}

// ReplySuccess: reply success status
func (svr *Server) ReplySuccess(resp http.ResponseWriter) error {
	_, err := resp.Write([]byte("success"))

	return err
}
