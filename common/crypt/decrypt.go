package crypt

import (
	"encoding/base64"
	"errors"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/idoubi/goutils/crypt"
)

// DecryptWithSessionKey: decrypt data
func DecryptWithSessionKey(sessionKey, encryptedData, iv string) (request.Result, error) {
	src, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}

	_key, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}

	_iv, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	decryptedData, err := crypt.AesCbcDecrypt(src, _key, _iv)
	if err != nil {
		return nil, err
	}

	return request.Result(decryptedData), nil
}

// DecryptWithAesKey: decrypt with aesKey
func DecryptWithAesKey(aesKey []byte, encryptedData string) ([]byte, string, error) {
	rawData, err := crypt.Base64Decode(encryptedData)
	if err != nil {
		return nil, "", err
	}

	decryptedData, err := crypt.AesCbcDecrypt(rawData, aesKey, nil)
	if err != nil {
		return nil, "", err
	}

	if len(decryptedData) < 20 {
		return nil, "", errors.New("decrypt error: invalid data length")
	}

	contentLen := getBytesLen(decryptedData[16:20])
	if contentLen > len(decryptedData)-20 {
		return nil, "", errors.New("decrypt error: invalid content length")
	}

	contentB := decryptedData[20 : 20+contentLen]

	// parse receiveId
	receiveIdB := decryptedData[20+contentLen:]

	return contentB, string(receiveIdB), nil
}

// get bytes length
func getBytesLen(bytes []byte) int {
	var num = 0
	for i := 0; i < 4; i++ {
		num <<= 8
		num |= (int)(bytes[i] & 0xff)
	}

	return num
}

// get length bytes
func getLenBytes(num int) []byte {
	return []byte{
		(byte)(num >> 24 & 0xFF),
		(byte)(num >> 16 & 0xF),
		(byte)(num >> 8 & 0xFF),
		(byte)(num & 0xFF),
	}
}
