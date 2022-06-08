package wxwork

import "github.com/cutesdk/cutesdk-go/common/request"

// FetchAccessToken: request get_access_token api
func (ins *Instance) FetchAccessToken() (*request.Result, error) {
	uri := "/cgi-bin/gettoken"

	res, err := ins.Get(uri, map[string]interface{}{
		"corpid":     ins.opts.Corpid,
		"corpsecret": ins.opts.Secret,
	})

	return res, err
}
