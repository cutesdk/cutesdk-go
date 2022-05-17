package wxopen

import (
	"bytes"
	"errors"
	"sort"
	"strings"

	"github.com/idoubi/goutils"
	"github.com/idoubi/goutils/crypt"
)

// DecryptMsg: decrypt message
func (s *Server) DecryptMsg(encryptedMsg string) ([]byte, error) {
	rawData, err := crypt.Base64Decode(encryptedMsg)
	if err != nil {
		return nil, err
	}

	decryptedData, err := crypt.AesCbcDecrypt(rawData, s.opts.aesKey, nil)
	if err != nil {
		return nil, err
	}

	if len(decryptedData) < 20 {
		return nil, errors.New("decrypt error: invalid data length")
	}

	contentLen := getBytesLen(decryptedData[16:20])
	if contentLen > len(decryptedData)-20 {
		return nil, errors.New("decrypt error: invalid content length")
	}

	contentB := decryptedData[20 : 20+contentLen]

	// parse appid
	appidB := decryptedData[20+contentLen:]

	if string(appidB) != s.GetComponentAppid() {
		return nil, errors.New("decrypt error: invalid appid")
	}

	return contentB, nil
}

// EncryptMsg: encrypt message
func (s *Server) EncryptMsg(rawMsg []byte) (string, error) {
	// part1: nonce_str 16bit
	nonceStr := goutils.NonceStr(16)

	// part2: msg length
	lenBytes := getLenBytes(len(rawMsg))

	// part3: msg content

	// part4: wxopen appid

	dataArr := bytes.Join([][]byte{[]byte(nonceStr), lenBytes, rawMsg, []byte(s.GetComponentAppid())}, nil)

	encryptedData, err := crypt.AesCbcEncrypt(dataArr, s.opts.aesKey, nil)
	if err != nil {
		return "", err
	}

	encryptedDataB64 := crypt.Base64Encode(encryptedData)

	return encryptedDataB64, nil
}

// VerifyMsg: verify notify data
// signature=sha1(sort(Token、timestamp、nonce, msg_encrypt))
func (s *Server) VerifyMsg(timestamp, nonce, msgSignature, msgEncrypt string) error {
	if s.GenSign(timestamp, nonce, msgEncrypt) != msgSignature {
		return errors.New("invalid signature")
	}

	return nil
}

// GenSign: gen signature
func (s *Server) GenSign(timestamp, nonce, encrypt string) string {
	arr := []string{s.opts.VerifyToken, timestamp, nonce, encrypt}
	sort.Strings(arr)
	sign := crypt.Sha1Encode([]byte(strings.Join(arr, "")))

	return sign
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

// getLenBytes
func getLenBytes(num int) []byte {
	return []byte{
		(byte)(num >> 24 & 0xFF),
		(byte)(num >> 16 & 0xF),
		(byte)(num >> 8 & 0xFF),
		(byte)(num & 0xFF),
	}
}
