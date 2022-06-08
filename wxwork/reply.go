package wxwork

import (
	"net/http"
)

// ReplyMsg: reply message
type ReplyMsg struct {
}

// ReplySuccess 回复字符串success
func (ins *Instance) ReplySuccess(resp http.ResponseWriter) error {
	_, err := resp.Write([]byte("success"))

	return err
}
