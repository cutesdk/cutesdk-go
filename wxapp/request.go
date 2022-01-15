package wxapp

import (
	"errors"

	"github.com/idoubi/goz"
)

// ApiGet Make api request with get method
func (w *WxApp) ApiGet(apiPath string, queryParams map[string]string) (Result, error) {
	// init query params
	if queryParams == nil {
		queryParams = make(map[string]string)
	}

	// replace token
	if len(queryParams) > 0 {
		for k, v := range queryParams {
			if v == AccessToken {
				token, err := w.getAccessToken()
				if err != nil {
					return nil, errors.New("get access_token failed: " + err.Error())
				}
				queryParams[k] = token
			}
		}
	}

	apiUrl := apiBase + apiPath

	var resp *goz.Response
	var err error

	// make request
	resp, err = goz.Get(apiUrl, goz.Options{
		Debug: w.opts.Debug,
		Query: queryParams,
	})

	// return when request error
	if err != nil {
		return nil, err
	}

	// return response data
	body, err := resp.GetBody()

	return Result(body), err
}

// ApiPost Make api request with post method
func (w *WxApp) ApiPost(apiPath string, queryParams map[string]string, postData map[string]interface{}) (Result, error) {
	// init query params
	if queryParams == nil {
		queryParams = make(map[string]string)
	}

	// replace token
	if len(queryParams) > 0 {
		for k, v := range queryParams {
			if v == AccessToken {
				token, err := w.getAccessToken()
				if err != nil {
					return nil, errors.New("get access_token failed: " + err.Error())
				}
				queryParams[k] = token
			}
		}
	}

	apiUrl := apiBase + apiPath

	var resp *goz.Response
	var err error

	// make request
	resp, err = goz.Post(apiUrl, goz.Options{
		Debug: w.opts.Debug,
		Query: queryParams,
		JSON:  postData,
	})

	// return when request error
	if err != nil {
		return nil, err
	}

	// return response data
	body, err := resp.GetBody()

	return Result(body), err
}
