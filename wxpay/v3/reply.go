package wxpay

import (
	"encoding/json"
	"net/http"
)

// ReplyMsg: notify reply message type
type ReplyMsg struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ReplySuccess: response success
func (n *NotifyMsg) ReplySuccess() *ReplyMsg {
	return nil
}

// ReplyFail: response fail
func (n *NotifyMsg) ReplyFail(msg string) *ReplyMsg {
	return &ReplyMsg{
		Code:    "FAIL",
		Message: msg,
	}
}

// ReplyMessage 回复消息
func (svr *Server) ReplyMessage(resp http.ResponseWriter, msg *ReplyMsg) error {
	if msg == nil || msg.Code == "SUCCESS" {
		resp.WriteHeader(http.StatusOK)
		resp.Write(nil)

		return nil
	}

	reply, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusInternalServerError)

	_, err = resp.Write(reply)

	return err
}
