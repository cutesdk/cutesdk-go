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

// NotifyData 原始通知数据
type NotifyData struct {
	Appid      string `xml:"AppId"`
	MsgEncrypt string `xml:"Encrypt"`
}

// ReplyMsg 回复消息
type ReplyMsg struct {
	Type        string
	TextContent string
}

// MsgHandler 消息处理器
type MsgHandler func(msg *Result) *ReplyMsg

// HandleNotify 处理通知
func (w *WxOpen) HandleNotify(req *http.Request, resp http.ResponseWriter, msgHandler MsgHandler) error {
	res, err := w.GetNotifyData(req)
	if err != nil {
		return err
	}

	if msgHandler != nil {
		replyMsg := msgHandler(&res)
		if replyMsg != nil {
			resp.Write(([]byte(fmt.Sprintf("reply:%v", replyMsg))))
		}
	}

	return nil
}

// ReplySuccess 回复字符串success
func (w *WxOpen) ReplySuccess(resp http.ResponseWriter) error {
	_, err := resp.Write([]byte("success"))

	return err
}

// Reply 回复消息
func (w *WxOpen) Reply() error {
	return nil
}

// GetNotifyData 获取通知数据
func (w *WxOpen) GetNotifyData(req *http.Request) (Result, error) {
	queryParams := req.URL.Query()
	timestamp := queryParams.Get("timestamp")
	nonce := queryParams.Get("nonce")
	// signature := queryParams.Get("signature")
	msgSignature := queryParams.Get("msg_signature")

	if timestamp == "" || nonce == "" || msgSignature == "" {
		return nil, fmt.Errorf("notify data with invalid params")
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("notify data with invalid body: %v", err)
	}
	defer req.Body.Close()

	notifyData := &NotifyData{}
	err = xml.Unmarshal(body, &notifyData)
	if err != nil || notifyData.MsgEncrypt == "" || notifyData.Appid != w.opts.Appid {
		return nil, fmt.Errorf("notify data unmarshal failed")
	}

	if err := w.VerifyNotifyData(timestamp, nonce, msgSignature, notifyData.MsgEncrypt); err != nil {
		return nil, fmt.Errorf("notify data verify failed: %v", err)
	}

	res, err := w.DecryptNotifyData(notifyData.MsgEncrypt)
	if err != nil {
		return nil, fmt.Errorf("notify data decrypt failed: %v", err)
	}

	return res, nil
}

// VerifyNotifyData 验证通知数据
// signature=sha1(sort(Token、timestamp、nonce, msg_encrypt))
func (w *WxOpen) VerifyNotifyData(timestamp, nonce, msgSignature, msgEncrypt string) error {
	arr := []string{w.opts.VerifyToken, timestamp, nonce, msgEncrypt}
	// 字典序排列
	sort.Strings(arr)
	// sha1 加密
	signature := crypt.Sha1Encode([]byte(strings.Join(arr, "")))
	fmt.Println(signature)
	if signature != msgSignature {
		return errors.New("invalid signature")
	}

	return nil
}

// DecryptNotifyData 解密通知数据
func (w *WxOpen) DecryptNotifyData(encryptedData string) (Result, error) {
	rawData, err := crypt.Base64Decode(encryptedData)
	if err != nil {
		return nil, err
	}

	decryptedData, err := crypt.AesCbcDecrypt(rawData, w.opts.aesKey, nil)
	if err != nil {
		return nil, err
	}

	if len(decryptedData) < 20 {
		return nil, errors.New("decrypt error: invalid data length")
	}

	// 读取有效内容长度
	contentLen := getBytesLen(decryptedData[16:20])
	if contentLen > len(decryptedData)-20 {
		return nil, errors.New("decrypt error: invalid content length")
	}

	// 有效内容
	contentB := decryptedData[20 : 20+contentLen]

	// 尾部的开放平台appid
	appidB := decryptedData[20+contentLen:]

	if string(appidB) != w.opts.Appid {
		return nil, errors.New("decrypt error: invalid appid")
	}

	return Result(contentB), nil
}

// 获取网络字节序
func getBytesLen(bytes []byte) int {
	var num = 0
	for i := 0; i < 4; i++ {
		num <<= 8
		num |= (int)(bytes[i] & 0xff)
	}

	return num
}
