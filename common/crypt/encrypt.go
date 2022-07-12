package crypt

import (
	"bytes"
	"sort"
	"strings"

	"github.com/idoubi/goutils"
	"github.com/idoubi/goutils/crypt"
)

// GenMsgSignature: gen msg_signature
func GenMsgSignature(token, timestamp, nonce, msgEncrypt string) string {
	arr := []string{token, timestamp, nonce}
	if msgEncrypt != "" {
		arr = append(arr, msgEncrypt)
	}
	sort.Strings(arr)
	sign := crypt.Sha1Encode([]byte(strings.Join(arr, "")))

	return sign
}

// EncryptMsg: encrypt message
func EncryptMsg(aesKey, rawMsg []byte, appid string) (string, error) {
	// part1: nonce_str 16bit
	nonceStr := goutils.NonceStr(16)

	// part2: msg length
	lenBytes := getLenBytes(len(rawMsg))

	// part3: msg content

	// part4: appid

	dataArr := bytes.Join([][]byte{[]byte(nonceStr), lenBytes, rawMsg, []byte(appid)}, nil)

	encryptedData, err := crypt.AesCbcEncrypt(dataArr, aesKey, nil)
	if err != nil {
		return "", err
	}

	encryptedDataB64 := crypt.Base64Encode(encryptedData)

	return encryptedDataB64, nil
}
