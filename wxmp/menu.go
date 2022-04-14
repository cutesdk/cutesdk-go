package wxmp

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// GetCurrentMenu: get current custom menu
func (c *Client) GetCurrentMenu() (request.Result, error) {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get access_token failed: %v", accessToken)
	}

	uri := fmt.Sprintf("/cgi-bin/get_current_selfmenu_info?access_token=%s", accessToken)

	res, err := c.Get(uri)

	return res, err
}
