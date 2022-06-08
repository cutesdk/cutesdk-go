package wxapp

import (
	"github.com/cutesdk/cutesdk-go/common/crypt"
	"github.com/cutesdk/cutesdk-go/common/request"
)

// DecryptUserInfo: decrypt userinfo
func (ins *Instance) DecryptUserInfo(sessionKey, encryptedData, iv string) (*request.Result, error) {
	return crypt.DecryptWithSessionKey(sessionKey, encryptedData, iv)
}

// DecryptPhone: decrypt user phone
func (ins *Instance) DecryptPhone(sessionKey, encryptedData, iv string) (*request.Result, error) {
	return crypt.DecryptWithSessionKey(sessionKey, encryptedData, iv)
}
