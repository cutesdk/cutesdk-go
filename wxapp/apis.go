package wxapp

// Code2Session 获取登录凭证
func (w *WxApp) Code2Session(code string) (Result, error) {
	apiPath := "/sns/jscode2session"

	res, err := w.ApiGet(apiPath, map[string]string{
		"appid":      w.opts.Appid,
		"secret":     w.opts.AppSecret,
		"js_code":    code,
		"grant_type": "authorization_code",
	})

	return res, err
}
