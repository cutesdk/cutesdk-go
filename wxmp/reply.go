package wxmp

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

type CDATAText string

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

// ReplyText: new text reply msg
func (msg *NotifyMsg) ReplyText(content string) *ReplyMsg {
	if content == "" {
		return nil
	}

	reply := map[string]interface{}{
		"MsgType": "text",
		"Content": content,
	}

	return msg.Reply(reply)
}

// ReplyImage: new image reply msg
func (msg *NotifyMsg) ReplyImage(mediaId string) *ReplyMsg {
	if mediaId == "" {
		return nil
	}

	reply := map[string]interface{}{
		"MsgType": "image",
		"Image": map[string]interface{}{
			"MediaId": mediaId,
		},
	}

	return msg.Reply(reply)
}

// ReplyVoice: new voice reply msg
func (msg *NotifyMsg) ReplyVoice(mediaId string) *ReplyMsg {
	if mediaId == "" {
		return nil
	}

	reply := map[string]interface{}{
		"MsgType": "voice",
		"Voice": map[string]interface{}{
			"MediaId": mediaId,
		},
	}

	return msg.Reply(reply)
}

// ReplyVideo: new video reply msg
func (msg *NotifyMsg) ReplyVideo(mediaId, title, description string) *ReplyMsg {
	if mediaId == "" {
		return nil
	}

	reply := map[string]interface{}{
		"MsgType": "video",
		"Video": map[string]interface{}{
			"MediaId":     mediaId,
			"Title":       title,
			"Description": description,
		},
	}

	return msg.Reply(reply)
}

// ReplyNews: new news reply msg
func (msg *NotifyMsg) ReplyNews(title, description, url, picUrl string) *ReplyMsg {
	reply := map[string]interface{}{
		"MsgType":      "news",
		"ArticleCount": 1,
		"Articles": []map[string]interface{}{
			{
				"Title":       title,
				"Description": description,
				"PicUrl":      picUrl,
				"Url":         url,
			},
		},
	}

	return msg.Reply(reply)
}

// EncryptReplyMsg: encrypt ReplyMsg
func (ins *Instance) EncryptReplyMsg(msg *ReplyMsg) ([]byte, error) {
	xb := msg.Raw()

	msgEncrypt, err := crypt.EncryptMsg(ins.opts.aesKey, xb, ins.opts.Appid)
	if err != nil {
		return nil, fmt.Errorf("encrypt reply_msg failed: %v", err)
	}

	timestamp := goutils.TimestampStr()
	nonce := goutils.NonceStr(16)
	msgSignature := crypt.GenMsgSignature(ins.opts.VerifyToken, timestamp, nonce, msgEncrypt)

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
func (ins *Instance) ReplyEncryptedMsg(resp http.ResponseWriter, msg *ReplyMsg) error {
	encryptedMsg, err := ins.EncryptReplyMsg(msg)
	if err != nil {
		return err
	}

	resp.Header().Set("Content-Type", "text/xml")
	_, err = resp.Write(encryptedMsg)

	return err
}

// ReplyPlaintext: reply unencrypted msg
func (ins *Instance) ReplyPlaintext(resp http.ResponseWriter, msg *ReplyMsg) error {
	resp.Header().Set("Content-Type", "text/xml")
	_, err := resp.Write(msg.Raw())

	return err
}

// ReplySuccess: reply success status
func (ins *Instance) ReplySuccess(resp http.ResponseWriter) error {
	_, err := resp.Write([]byte("success"))

	return err
}
