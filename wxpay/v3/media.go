package wxpay

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/fileuploader"
)

// UploadImage: upload permanent image
func (cli *Client) UploadImage(filePath string) (*request.Result, error) {
	file, err := parseFile(filePath)
	if err != nil {
		return nil, err
	}

	svc := fileuploader.ImageUploader{Client: cli.payClient}
	_, res, err := svc.Upload(cli.ctx, file.content, file.fileName, file.contentType)
	if err != nil {
		return nil, err
	}

	return parseApiResult(res)
}

// UploadMarketingImage: upload marketing image
func (cli *Client) UploadMarketingImage(filePath string) (*request.Result, error) {
	file, err := parseFile(filePath)
	if err != nil {
		return nil, err
	}

	svc := fileuploader.MarketingImageUploader{Client: cli.payClient}
	_, res, err := svc.Upload(cli.ctx, file.content, file.fileName, file.contentType)
	if err != nil {
		return nil, err
	}

	return parseApiResult(res)
}

// UploadMerchantImage: upload merchant-service image
func (cli *Client) UploadMerchantImage(filePath string) (*request.Result, error) {
	file, err := parseFile(filePath)
	if err != nil {
		return nil, err
	}

	svc := fileuploader.MchBizUploader{Client: cli.payClient}
	_, res, err := svc.Upload(cli.ctx, file.content, file.fileName, file.contentType)
	if err != nil {
		return nil, err
	}

	return parseApiResult(res)
}

// UploadVideo: upload video
func (cli *Client) UploadVideo(filePath string) (*request.Result, error) {
	file, err := parseFile(filePath)
	if err != nil {
		return nil, err
	}

	svc := fileuploader.VideoUploader{Client: cli.payClient}
	_, res, err := svc.Upload(cli.ctx, file.content, file.fileName, file.contentType)
	if err != nil {
		return nil, err
	}

	return parseApiResult(res)
}

// File: file to be uploaded
type File struct {
	fileName    string
	contentType string
	content     *bytes.Buffer
}

func parseFile(filePath string) (*File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fs, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fd := make([]byte, fs.Size())
	file.Read(fd)

	fileName := filepath.Base(filePath)
	contentType := http.DetectContentType(fd)
	content := bytes.NewBuffer(fd)

	f := &File{fileName, contentType, content}

	return f, nil
}

func parseApiResult(res *core.APIResult) (*request.Result, error) {
	defer res.Response.Body.Close()

	body, err := ioutil.ReadAll(res.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("parse response body failed: %v", err)
	}

	return request.NewResult(body), nil
}
