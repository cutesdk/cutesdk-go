package wxwork

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// UploadTempMedia: upload temporary media
func (cli *Client) UploadTempMedia(mediaType, filePath string) (*request.Result, error) {
	uri := fmt.Sprintf("/cgi-bin/media/upload?type=%s", mediaType)
	params := map[string]interface{}{
		"media": fmt.Sprintf("@%s", filePath),
	}

	res, err := cli.PostMultipartWithToken(uri, params)

	return res, err
}

// UploadImage: upload permanent image
func (cli *Client) UploadImage(filePath string) (*request.Result, error) {
	uri := "/cgi-bin/media/uploadimg"
	params := map[string]interface{}{
		"media": fmt.Sprintf("@%s", filePath),
	}

	res, err := cli.PostMultipartWithToken(uri, params)

	return res, err
}
