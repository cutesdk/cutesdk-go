package wxpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
)

// NotifyMsg: notify message type
type NotifyMsg struct {
	*request.Result
	reqData *request.Result
}

// GetReqData: get request data
func (m *NotifyMsg) GetReqData() *request.Result {
	return m.reqData
}

// NotifyHandler: notify handler
type NotifyHandler func(*NotifyMsg) *ReplyMsg

// Listen: listen notify
func (svr *Server) Listen(req *http.Request, resp http.ResponseWriter, notifyHandler NotifyHandler) error {
	msg, err := svr.GetNotifyMsg(req)
	if err != nil {
		return fmt.Errorf("get notify message failed: %v", err)
	}

	if notifyHandler == nil {
		return nil
	}

	replyMsg := notifyHandler(msg)

	return svr.ReplyMessage(resp, replyMsg)
}

// GetReqBody: get request data in body
func (svr *Server) GetReqBody(req *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("invalid notify data: %v", err)
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}

// GetNotifyMsg: get notify message
func (svr *Server) GetNotifyMsg(req *http.Request) (*NotifyMsg, error) {
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(svr.opts.MchId)
	handler := notify.NewNotifyHandler(svr.opts.ApiKey, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))

	content := make(map[string]interface{})
	notifyReq, err := handler.ParseNotifyRequest(svr.ctx, req, content)
	if err != nil {
		return nil, err
	}

	reqData, err := json.Marshal(notifyReq)
	if err != nil {
		return nil, err
	}

	res := request.NewResult([]byte(notifyReq.Resource.Plaintext))
	data := request.NewResult(reqData)

	notifyMsg := &NotifyMsg{
		res,
		data,
	}

	return notifyMsg, nil
}
