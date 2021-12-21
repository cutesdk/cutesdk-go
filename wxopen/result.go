package wxopen

import (
	"github.com/idoubi/goutils/convert"
	"github.com/tidwall/gjson"
)

// Result 响应结果
type Result []byte

// Parsed 获取解析后的数据
func (r Result) Parsed() gjson.Result {
	return gjson.ParseBytes(r)
}

// XmlParsed 解析xml数据
func (r Result) XmlParsed() gjson.Result {
	jb, _ := convert.Xml2Json(r)

	return gjson.ParseBytes(jb)
}
