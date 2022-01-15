package wxapp

import "github.com/tidwall/gjson"

// Result 响应数据
type Result []byte

// Parsed 转为解析后的数据
func (r Result) Parsed() gjson.Result {
	return gjson.ParseBytes(r)
}

// String 字符串打印
func (r Result) String() string {
	return string(r)
}

// Get 获取值
func (r Result) Get(key string) gjson.Result {
	return r.Parsed().Get(key)
}
