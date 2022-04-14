package ttapp

import (
	"github.com/cutesdk/cutesdk-go/common/crypt"
	"github.com/cutesdk/cutesdk-go/common/request"
)

// DecryptUserInfo: decrypt userinfo
func (c *Client) DecryptUserInfo(sessionKey, encryptedData, iv string) (request.Result, error) {
	return crypt.DecryptData(sessionKey, encryptedData, iv)
}
