package wxpay

import (
	"crypto/x509"
	"errors"

	"github.com/idoubi/goutils"
	"github.com/idoubi/goutils/crypt"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// DecryptData: decrypt data
func (cli *Client) DecryptData(ciphertext, nonce, associatedData string) (string, error) {
	encryptedData := []byte(goutils.Base64Decode(ciphertext))
	res, err := crypt.AesGcmDecrypt([]byte(encryptedData), []byte(associatedData), []byte(nonce), []byte(cli.GetApiKey()))
	if err != nil {
		return "", nil
	}

	return string(res), err
}

// DecryptSensitiveData: decrypt sensitive data
func (cli *Client) DecryptSensitiveData(ciphertext string) (string, error) {
	res, err := utils.DecryptOAEP(ciphertext, cli.opts.privateKey)

	return res, err
}

// SensitiveCryptor: sensitive data cryptor
type SensitiveCryptor struct {
	serialNo    string
	certificate *x509.Certificate
}

// GetSensitiveCryptor: get cryptor with cert
func (cli *Client) GetSensitiveCryptor() (*SensitiveCryptor, error) {
	cv := downloader.MgrInstance().GetCertificateVisitor(cli.opts.MchId)
	// get latest serialNo
	serialNo := cv.GetNewestSerial(cli.ctx)
	certificate, ok := cv.Get(cli.ctx, serialNo)
	if !ok || certificate == nil {
		return nil, errors.New("get certificate failed")
	}

	sc := &SensitiveCryptor{serialNo, certificate}

	return sc, nil
}

// GetSerialNo: get serial no
func (sc *SensitiveCryptor) GetSerialNo() string {
	return sc.serialNo
}

// Encrypt: encrypt data
func (sc *SensitiveCryptor) Encrypt(rawData string) string {
	res, _ := utils.EncryptOAEPWithCertificate(rawData, sc.certificate)

	return res
}
