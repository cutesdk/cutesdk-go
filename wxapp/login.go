package wxapp

import (
	"fmt"

	"github.com/idoubi/goz"
	"github.com/tidwall/gjson"
)

// Code2Session 获取登录凭证
func (w *Wxapp) Code2Session(code string) (*gjson.Result, error) {
	uri := fmt.Sprintf(apiBase+"/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", w.opts.Appid, w.opts.Appsecret, code)
	res, err := goz.Get(uri)

	if err != nil {
		return nil, err
	}

	return res.GetParsedBody()
}

// DecryptData 用户信息解密
func (w *Wxapp) DecryptData(sessionKey, encryptedData, iv string) (*UserInfo, error) {
	dc := NewUserDataCrypt(w.opts.Appid, sessionKey)
	userInfo, err := dc.Decrypt(encryptedData, iv)

	return userInfo, err
}
