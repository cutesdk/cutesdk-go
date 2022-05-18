package wxopen

import (
	"fmt"
	"net/http"

	"github.com/idoubi/goutils"
)

// ReplyText: new text reply msg
func (msg *NotifyMsg) ReplyText(content string) *ReplyMsg {
	if content == "" {
		return nil
	}

	return &ReplyMsg{
		ToUserName:   CDATAText(msg.FromUserName),
		FromUserName: CDATAText(msg.ToUserName),
		CreateTime:   goutils.TimestampStr(),
		MsgType:      "text",
		Content:      CDATAText(content),
	}
}

// MsgHandler: notify message handler
type MsgHandler func(msg *NotifyMsg) *ReplyMsg

// Listen: listen notify
func (s *Server) Listen(req *http.Request, resp http.ResponseWriter, msgHandler MsgHandler) error {
	msg, err := s.GetMessage(req)
	if err != nil {
		fmt.Printf("getmsg err:%v", err)
		return err
	}

	if msgHandler == nil {
		return nil
	}

	replyMsg := msgHandler(msg)
	if replyMsg == nil {
		return s.ReplySuccess(resp)
	}

	return s.ReplyMessage(resp, replyMsg)
}
