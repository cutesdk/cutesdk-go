package request

import (
	"os"

	"github.com/idoubi/goutils/convert"
	"github.com/tidwall/gjson"
)

// Result is the response data from api request
type Result []byte

// Parsed is the formated data of Result
func (r Result) Parsed() gjson.Result {
	return gjson.ParseBytes(r)
}

// XmlParsed 解析xml数据
func (r Result) XmlParsed() gjson.Result {
	jb, _ := convert.Xml2Json(r)

	return gjson.ParseBytes(jb).Get("xml")
}

// String print the result as string
func (r Result) String() string {
	return string(r)
}

// Get a sub item from the Parsed data
func (r Result) Get(key string) gjson.Result {
	return r.Parsed().Get(key)
}

// GetString: get item with string format
func (r Result) GetString(key string) string {
	return r.Parsed().Get(key).String()
}

// GetInt: get item with int64 format
func (r Result) GetInt(key string) int64 {
	return r.Parsed().Get(key).Int()
}

// SaveAsFile: save as file
func (r Result) SaveAsFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(r)

	return err
}
