package wxapp

import (
	"fmt"

	"github.com/idoubi/goz"
)

// Code2Session 获取登录凭证
func (w *Wxapp) Code2Session(code string) (Result, error) {
	apiURL := fmt.Sprintf(apiBase+"/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", w.opts.Appid, w.opts.Secret, code)
	resp, err := goz.Get(apiURL, goz.Options{
		Debug: w.opts.Debug,
	})

	if err != nil {
		return nil, err
	}

	body, err := resp.GetBody()

	return Result(body), err
}

// DecryptData 用户信息解密
func (w *Wxapp) DecryptData(sessionKey, encryptedData, iv string) (*UserInfo, error) {
	dc := NewUserDataCrypt(w.opts.Appid, sessionKey)
	userInfo, err := dc.Decrypt(encryptedData, iv)

	return userInfo, err
}
