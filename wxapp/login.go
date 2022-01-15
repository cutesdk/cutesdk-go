package wxapp

// DecryptData 用户信息解密
func (w *WxApp) DecryptData(sessionKey, encryptedData, iv string) (*UserInfo, error) {
	dc := NewUserDataCrypt(w.opts.Appid, sessionKey)
	userInfo, err := dc.Decrypt(encryptedData, iv)

	return userInfo, err
}
