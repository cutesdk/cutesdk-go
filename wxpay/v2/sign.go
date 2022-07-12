package wxpay

import (
	"fmt"
	"sort"
	"strings"

	"github.com/idoubi/goutils"
)

// SignWithHmacSha256: sign with hmac-sha256 method
func SignWithHmacSha256(params map[string]interface{}, key string) string {
	signStr := BuildSignStr(params, key)
	sign := strings.ToUpper(goutils.HmacSha256(signStr, key))

	return sign
}

// SignWithMd5: sign with md5 method
func SignWithMd5(params map[string]interface{}, key string) string {
	signStr := BuildSignStr(params, key)
	sign := strings.ToUpper(goutils.MD5(signStr))

	return sign
}

// BuildSignStr: build sign str use sign params
func BuildSignStr(params map[string]interface{}, key string) string {
	keys := []string{}
	for k := range params {
		keys = append(keys, k)
	}

	// sort param keys
	sort.Strings(keys)

	arr := []string{}
	str := ""

	for _, k := range keys {
		if k == "sign" {
			continue
		}
		if params[k] == nil {
			continue
		}
		if v, ok := params[k].(string); ok && v == "" {
			continue
		}
		arr = append(arr, fmt.Sprintf("%s=%v", k, params[k]))
	}

	str = strings.Join(arr, "&")
	str += fmt.Sprintf("&%s=%s", "key", key)

	return str
}
