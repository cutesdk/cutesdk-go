package wxopen

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/idoubi/goutils/crypt"
)

// VerifyNotifyData 验证通知数据
// signature=sha1(sort(Token、timestamp、nonce, msg_encrypt))
func (w *WxOpen) VerifyNotifyData(timestamp, nonce, msgSignature, msgEncrypt string) error {
	arr := []string{w.opts.VerifyToken, timestamp, nonce, msgEncrypt}
	// 字典序排列
	sort.Strings(arr)
	// sha1 加密
	signature := crypt.Sha1Encode([]byte(strings.Join(arr, "")))
	fmt.Println(signature)
	if signature != msgSignature {
		return errors.New("invalid signature")
	}

	return nil
}

// GetNotifyData 获取解密后的通知数据
func (w *WxOpen) GetNotifyData(encryptedData string) (Result, error) {
	rawData, err := crypt.Base64Decode(encryptedData)
	if err != nil {
		return nil, err
	}

	decryptedData, err := crypt.AesCbcDecrypt(rawData, w.opts.aesKey, nil)
	if err != nil {
		return nil, err
	}

	if len(decryptedData) < 20 {
		return nil, errors.New("decrypt error: invalid data length")
	}

	// 读取有效内容长度
	contentLen := getBytesLen(decryptedData[16:20])
	if contentLen > len(decryptedData)-20 {
		return nil, errors.New("decrypt error: invalid content length")
	}

	// 有效内容
	contentB := decryptedData[20 : 20+contentLen]

	// 尾部的开放平台appid
	appidB := decryptedData[20+contentLen:]

	if string(appidB) != w.opts.Appid {
		return nil, errors.New("decrypt error: invalid appid")
	}

	return Result(contentB), nil
}

// 获取网络字节序
func getBytesLen(bytes []byte) int {
	var num = 0
	for i := 0; i < 4; i++ {
		num <<= 8
		num |= (int)(bytes[i] & 0xff)
	}

	return num
}
