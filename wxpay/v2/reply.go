package wxpay

import (
	"encoding/xml"
	"net/http"
)

type CDATAText string

func (c CDATAText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

// ReplyMsg: notify reply message type
type ReplyMsg struct {
	XMLName    xml.Name  `xml:"xml"`
	ReturnCode CDATAText `xml:"return_code"`
	ReturnMsg  CDATAText `xml:"return_msg"`
}

// ReplySuccess: response success
func (n *NotifyMsg) ReplySuccess() *ReplyMsg {
	return &ReplyMsg{
		ReturnCode: CDATAText("SUCCESS"),
		ReturnMsg:  CDATAText("OK"),
	}
}

// ReplyFail: response fail
func (n *NotifyMsg) ReplyFail(msg string) *ReplyMsg {
	return &ReplyMsg{
		ReturnCode: CDATAText("FAIL"),
		ReturnMsg:  CDATAText(msg),
	}
}

// ReplyMessage 回复消息
func (s *Server) ReplyMessage(resp http.ResponseWriter, msg *ReplyMsg) error {
	reply, err := xml.Marshal(msg)
	if err != nil {
		return err
	}

	resp.Header().Set("Content-Type", "text/xml")
	_, err = resp.Write(reply)

	return err
}
