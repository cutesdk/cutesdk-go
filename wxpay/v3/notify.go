package wxpay

import (
	"net/http"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
)

// GetNotifyData: get notify data
func (ins *Instance) GetNotifyData(req *http.Request) (*notify.Request, *request.Result, error) {
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(ins.opts.MchId)
	handler := notify.NewNotifyHandler(ins.opts.ApiKey, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))

	content := make(map[string]interface{})
	notifyReq, err := handler.ParseNotifyRequest(ins.ctx, req, content)
	if err != nil {
		return nil, nil, err
	}

	return notifyReq, request.NewResult([]byte(notifyReq.Resource.Plaintext)), nil
}
