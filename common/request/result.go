package request

import (
	"os"

	"github.com/idoubi/goutils/convert"
	"github.com/tidwall/gjson"
)

// Result is the response data from api request
type Result struct {
	raw []byte
	res *gjson.Result
}

// NewResult
func NewResult(data []byte) *Result {
	r := &Result{raw: data}

	// default parsed with json
	r.Parsed()

	return r
}

// Raw get raw data
func (r *Result) Raw() []byte {
	return r.raw
}

// Map convert to map
func (r *Result) Map() map[string]gjson.Result {
	return r.res.Map()
}

// Parsed to json object
func (r *Result) Parsed() gjson.Result {
	res := gjson.ParseBytes(r.raw)
	r.res = &res

	return res
}

// XmlParsed to json object
func (r *Result) XmlParsed() gjson.Result {
	jb, _ := convert.Xml2Json(r.raw)

	res := gjson.ParseBytes(jb).Get("xml")
	r.res = &res

	return res
}

// String print the result as string
func (r *Result) String() string {
	return r.res.String()
}

// Get a sub item from the Parsed data
func (r *Result) Get(key string) gjson.Result {
	return r.res.Get(key)
}

// GetString: get with string format
func (r *Result) GetString(key string) string {
	return r.Get(key).String()
}

// GetInt: get with int64 format
func (r *Result) GetInt(key string) int64 {
	return r.Get(key).Int()
}

// GetBool: get with bool format
func (r *Result) GetBool(key string) bool {
	return r.Get(key).Bool()
}

// GetArray: get item with array format
func (r *Result) GetArray(key string) []gjson.Result {
	return r.Get(key).Array()
}

// GetMap: get item with array format
func (r *Result) GetMap(key string) map[string]gjson.Result {
	return r.Get(key).Map()
}

// SaveAsFile: save as file
func (r *Result) SaveAsFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(r.raw)

	return err
}
