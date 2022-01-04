package wxopen

// PushTicket api_start_push_ticket
func (w *WxOpen) PushTicket() (Result, error) {
	apiPath := "/cgi-bin/component/api_start_push_ticket"

	res, err := w.ApiPost(apiPath, nil, map[string]interface{}{
		"component_appid":  w.opts.Appid,
		"component_secret": w.opts.AppSecret,
	})

	return res, err
}
