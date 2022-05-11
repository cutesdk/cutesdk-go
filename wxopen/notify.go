package wxopen

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/idoubi/goutils/crypt"
)

// NotifyInfo: notify info data
type NotifyInfo struct {
	Appid                 string `xml:"AppId"`
	InfoType              string `xml:"InfoType,omitempty"`
	ComponentVerifyTicket string `xml:"ComponentVerifyTicket,omitempty"`
}

// NotifyMessage: notify msg
type NotifyMessage struct {
	ToUserName   string `xml:"ToUserName,omitempty"`
	FromUserName string `xml:"FromUserName,omitempty"`
	MsgType      string `xml:"MsgType,omitempty"`
	Content      string `xml:"Content,omitempty"`
	MsgId        string `xml:"MsgId,omitempty"`
}

// NotifyData: notify data
type NotifyData struct {
	NotifyInfo
	NotifyMessage
	MsgEncrypt string `xml:"Encrypt,omitempty"`
	CreateTime string `xml:"CreateTime,omitempty"`
}

// ReplyMsg 回复消息
type ReplyMsg struct {
	Type        string
	TextContent string
}

// MsgHandler: notify message handler
type MsgHandler func(msg *NotifyData)

// ReplySuccess 回复字符串success
func (s *Server) ReplySuccess(resp http.ResponseWriter) error {
	_, err := resp.Write([]byte("success"))

	return err
}

// Reply 回复消息
func (s *Server) Reply() error {
	return nil
}

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

// GetNotifyData: get notify data
func (s *Server) GetNotifyData(req *http.Request) (*NotifyData, error) {
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
	if err != nil || notifyData.MsgEncrypt == "" {
		return nil, fmt.Errorf("notify data unmarshal failed")
	}

	if notifyData.Appid != "" && notifyData.Appid != s.GetComponentAppid() {
		return nil, fmt.Errorf("notify data with invalid appid")
	}

	if err := s.VerifyNotifyData(timestamp, nonce, msgSignature, notifyData.MsgEncrypt); err != nil {
		return nil, fmt.Errorf("notify data verify failed: %v", err)
	}

	res, err := s.DecryptNotifyData(notifyData.MsgEncrypt)
	if err != nil {
		return nil, fmt.Errorf("notify data decrypt failed: %v", err)
	}

	return res, nil
}

// VerifyNotifyData: verify notify data
// signature=sha1(sort(Token、timestamp、nonce, msg_encrypt))
func (s *Server) VerifyNotifyData(timestamp, nonce, msgSignature, msgEncrypt string) error {
	arr := []string{s.opts.VerifyToken, timestamp, nonce, msgEncrypt}
	sort.Strings(arr)
	signature := crypt.Sha1Encode([]byte(strings.Join(arr, "")))

	if signature != msgSignature {
		return errors.New("invalid signature")
	}

	return nil
}

// DecryptNotifyData: decrypt encrypted data in notify data
func (s *Server) DecryptNotifyData(encryptedData string) (*NotifyData, error) {
	rawData, err := crypt.Base64Decode(encryptedData)
	if err != nil {
		return nil, err
	}

	decryptedData, err := crypt.AesCbcDecrypt(rawData, s.opts.aesKey, nil)
	if err != nil {
		return nil, err
	}

	if len(decryptedData) < 20 {
		return nil, errors.New("decrypt error: invalid data length")
	}

	contentLen := getBytesLen(decryptedData[16:20])
	if contentLen > len(decryptedData)-20 {
		return nil, errors.New("decrypt error: invalid content length")
	}

	contentB := decryptedData[20 : 20+contentLen]

	// parse appid
	appidB := decryptedData[20+contentLen:]

	if string(appidB) != s.GetComponentAppid() {
		return nil, errors.New("decrypt error: invalid appid")
	}

	fmt.Println(string(contentB))

	notifyData := &NotifyData{}
	if err := xml.Unmarshal(contentB, &notifyData); err != nil {
		return nil, err
	}

	return notifyData, nil
}

// get bytes length
func getBytesLen(bytes []byte) int {
	var num = 0
	for i := 0; i < 4; i++ {
		num <<= 8
		num |= (int)(bytes[i] & 0xff)
	}

	return num
}
