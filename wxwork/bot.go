package wxwork

import (
	"encoding/xml"
	"errors"
	"strings"

	"github.com/idoubi/cutesdk/wxwork/wxbizmsgcrypt"
	"github.com/idoubi/goutils"
	"github.com/idoubi/goz"
)

// Bot 企业微信机器人
type Bot struct {
	Token          string      // 接入验证的token
	AesKey         string      // 接入验证的encodingAesKey
	WebhookURL     string      // 主动推送消息的地址
	mentionedUsers []string    // 提到的用户
	visibleUsers   []string    // 消息可见的用户
	actions        []MsgAction // 消息操作
	callbackID     string      // 消息操作回调ID
	reqMsg         *ReqMsg     // 请求消息
	respMsg        *RespMsg    // 响应消息
}

// ReqMsg 请求消息
type ReqMsg struct {
	From           ReqMsgFrom    `xml:"From"`
	WebhookURL     string        `xml:"WebhookUrl"`
	ChatID         string        `xml:"ChatId"`
	ChatType       string        `xml:"ChatType"`
	GetChatInfoURL string        `xml:"GetChatInfoUrl"`
	MsgID          string        `xml:"MsgId"`
	MsgType        string        `xml:"MsgType"`
	Text           ReqMsgText    `xml:"Text"`
	Attachment     MsgAttachment `xml:"Attachment"`
}

// ReqMsgFrom 请求消息来源
type ReqMsgFrom struct {
	UserID string `xml:"UserId"`
	Name   string `xml:"Name`
	Alias  string `xml:"Alias"`
}

// ReqMsgText 请求消息内容
type ReqMsgText struct {
	Content string `xml:"Content"`
}

// RespMsg 响应消息
type RespMsg struct {
	MsgType       string           `xml:"MsgType" json:"msgtype"`
	VisibleToUser string           `xml:"VisibleToUser" json:"-"`
	Text          *RespMsgText     `xml:"Text" json:"text,omitempty"`
	Markdown      *RespMsgMarkdown `xml:"Markdown" json:"markdown,omitempty"`
}

// RespMsgText 文本消息类型
type RespMsgText struct {
	Content     string         `xml:"Content" json:"content,omitempty"`
	MentionList *MentionedList `xml:"MentionedList,omitempty" json:"mentioned_list,omitempty"`
}

// RespMsgMarkdown markdown消息类型
type RespMsgMarkdown struct {
	Content    string         `xml:"Content" json:"content,omitempty"`
	Attachment *MsgAttachment `xml:"Attachment,omitempty" json:"attachment,omitempty"`
}

// MsgAttachment markdown附加数据
type MsgAttachment struct {
	CallbackID string      `xml:"CallbackId"`
	Actions    []MsgAction `xml:"Actions"`
}

// MsgAction 操作
type MsgAction struct {
	Name        string `xml:"Name"`
	Value       string `xml:"Value"`
	Text        string `xml:"Text"`
	Type        string `xml:"Type"`
	BorderColor string `xml:"BorderColor"`
	TextColor   string `xml:"TextColor"`
	ReplaceText string `xml:"ReplaceText"`
}

// MentionedList 提到的人
type MentionedList struct {
	Item []string `xml:"Item"`
}

// Valid 接入验证
func (b *Bot) Valid(msgSign, timestamp, nonce, echoStr string) (string, error) {
	wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(b.Token, b.AesKey, "", wxbizmsgcrypt.XmlType)

	bt, err := wxcpt.VerifyURL(msgSign, timestamp, nonce, echoStr)
	if err != nil {
		return "", err
	}

	return string(bt), nil
}

// DecryptMsg 消息解密
func (b *Bot) DecryptMsg(msgSign, timestamp, nonce string, rawMsg []byte) (*ReqMsg, error) {
	wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(b.Token, b.AesKey, "", wxbizmsgcrypt.XmlType)

	bt, err := wxcpt.DecryptMsg(msgSign, timestamp, nonce, rawMsg)
	if err != nil {
		return nil, err
	}

	reqMsg := &ReqMsg{}
	if err := xml.Unmarshal(bt, &reqMsg); err != nil {
		return nil, err
	}

	b.reqMsg = reqMsg

	return reqMsg, nil
}

// Mention 艾特用户
func (b *Bot) Mention(users []string) *Bot {
	b.mentionedUsers = users
	return b
}

// Visible 指定用户可见
func (b *Bot) Visible(users []string) *Bot {
	b.visibleUsers = users
	return b
}

// Actions 设置操作
func (b *Bot) Actions(callbackID string, actions []MsgAction) *Bot {
	b.callbackID = callbackID
	b.actions = actions
	return b
}

// Text 构建文本消息
func (b *Bot) Text(text string) *Bot {
	respMsg := &RespMsg{
		MsgType: "text",
		Text: &RespMsgText{
			Content:     text,
			MentionList: nil,
		},
	}
	if len(b.mentionedUsers) > 0 {
		respMsg.Text.MentionList = &MentionedList{
			Item: b.mentionedUsers,
		}
	}

	b.respMsg = respMsg

	return b
}

// Markdown 构建markdown消息
func (b *Bot) Markdown(markdown string) *Bot {
	respMsg := &RespMsg{
		MsgType: "markdown",
		Markdown: &RespMsgMarkdown{
			Content:    markdown,
			Attachment: nil,
		},
	}

	if b.callbackID != "" && len(b.actions) > 0 {
		respMsg.Markdown.Attachment = &MsgAttachment{
			CallbackID: b.callbackID,
			Actions:    b.actions,
		}
	}

	b.respMsg = respMsg

	return b
}

// Reply 被动回复消息
func (b *Bot) Reply() ([]byte, error) {
	wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(b.Token, b.AesKey, "", wxbizmsgcrypt.XmlType)

	if len(b.visibleUsers) > 0 {
		b.respMsg.VisibleToUser = strings.Join(b.visibleUsers, "|")
	}

	bt, err := xml.Marshal(b.respMsg)
	if err != nil {
		return nil, err
	}

	timestamp := goutils.TimestampStr()
	nonce := goutils.NonceStr(16)
	msg, err := wxcpt.EncryptMsg(string(bt), timestamp, nonce)
	if len(msg) == 0 {
		return nil, errors.New("no reply msg")
	}

	return msg, nil
}

// Send 主动发送消息
func (b *Bot) Send() ([]byte, error) {
	cli := goz.NewClient()
	res, err := cli.Post(b.WebhookURL, goz.Options{
		// Debug: true,
		// Proxy: "http://devnet-proxy.oa.com:8080",
		JSON: b.respMsg,
	})

	if err != nil {
		return nil, err
	}

	body, err := res.GetBody()
	if err != nil {
		return nil, err
	}

	return []byte(body), nil
}
