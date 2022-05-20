package wxwork

import "github.com/cutesdk/cutesdk-go/common/request"

// FetchAccessToken: request get_access_token api
func (c *Client) FetchAccessToken() (request.Result, error) {
	uri := "/cgi-bin/gettoken"

	res, err := c.Get(uri, map[string]interface{}{
		"corpid":     c.GetCorpid(),
		"corpsecret": c.GetSecret(),
	})

	return res, err
}
