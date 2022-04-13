package request

import (
	"os"

	"github.com/tidwall/gjson"
)

// Result is the response data from api request
type Result []byte

// Parsed is the formated data of Result
func (r Result) Parsed() gjson.Result {
	return gjson.ParseBytes(r)
}

// String print the result as string
func (r Result) String() string {
	return string(r)
}

// Get a sub item from the Parsed data
func (r Result) Get(key string) gjson.Result {
	return r.Parsed().Get(key)
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
