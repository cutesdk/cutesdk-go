package request

import (
	"github.com/idoubi/goz"
)

// Client: request client
type Client struct {
	opts *Options
}

// NewClient: new request client
func NewClient(opts *Options) *Client {
	return &Client{opts}
}

// Get: make api request with get method
func (c *Client) Get(uri string, args ...map[string]interface{}) (Result, error) {
	params := map[string]interface{}{}
	headers := map[string]interface{}{}

	if len(args) > 0 {
		params = args[0]
	}
	if len(args) > 1 {
		headers = args[1]
	}

	method := "GET"

	opts := goz.Options{
		Debug:   c.opts.Debug,
		Timeout: float32(c.opts.Timeout.Seconds()),
		Query:   params,
		Headers: headers,
	}

	return c.Request(method, uri, opts)
}

// Post: make api request with post method
func (c *Client) Post(uri string, args ...map[string]interface{}) (Result, error) {
	data := map[string]interface{}{}
	headers := map[string]interface{}{}

	if len(args) > 0 {
		data = args[0]
	}
	if len(args) > 1 {
		headers = args[1]
	}

	method := "POST"

	opts := goz.Options{
		Debug:   c.opts.Debug,
		Timeout: float32(c.opts.Timeout.Seconds()),
		JSON:    data,
		Headers: headers,
	}

	return c.Request(method, uri, opts)
}

// Request: make api request
func (c *Client) Request(method string, uri string, opts goz.Options) (Result, error) {
	cli := goz.NewClient()

	apiUrl := c.opts.BaseUri + uri

	resp, err := cli.Request(method, apiUrl, opts)

	if err != nil {
		return nil, err
	}

	body, err := resp.GetBody()

	return Result(body), err
}
