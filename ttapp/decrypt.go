package ttapp

import (
	"github.com/cutesdk/cutesdk-go/common/crypt"
	"github.com/cutesdk/cutesdk-go/common/request"
)

// DecryptUserInfo: decrypt userinfo
func (cli *Client) DecryptUserInfo(sessionKey, encryptedData, iv string) (*request.Result, error) {
	return crypt.DecryptWithSessionKey(sessionKey, encryptedData, iv)
}
