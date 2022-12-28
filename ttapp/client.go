package ttapp

import (
	"fmt"
	"net/url"
	"time"

	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/common/token"
	"github.com/idoubi/goz"
)

// Client: ttapp client
type Client struct {
	opts                *Options
	requestClient       *request.Client
	cacheHandler        cache.ICache
	accessTokenCacheKey string
	accessTokenHandler  token.IToken
}

// NewClient: new ttapp client
func NewClient(opts *Options) (*Client, error) {
	// set default options
	if opts.Timeout <= 0 {
		opts.Timeout = 10 * time.Second
	}

	if opts.Cache == nil {
		opts.Cache = &cache.FileOptions{
			Dir: "./cache",
		}
	}

	// new client
	client := &Client{opts: opts}

	// new cache handler
	cacheHandler, err := cache.NewCache(opts.Cache)
	if err != nil {
		return nil, fmt.Errorf("new cache handler failed: %v", err)
	}

	// set cache handler
	client.SetCacheHandler(cacheHandler)

	// set request client
	client.requestClient = request.NewClient(&request.Options{
		BaseUri: "https://developer.toutiao.com",
		Debug:   opts.Debug,
		Timeout: opts.Timeout,
	})

	// set default access_token cache key
	accessTokenCacheKey := fmt.Sprintf("ttapp.access_token.%s", client.GetAppid())
	client.SetAccessTokenCacheKey(accessTokenCacheKey)

	// set default access_token handler
	client.SetAccessTokenHandler(NewAccessToken(client))

	return client, nil
}

// GetAccessToken: get access_token from cache or api
func (cli *Client) GetAccessToken() (string, error) {
	return cli.GetAccessTokenHandler().GetToken()
}

// RefreshAccessToken: refresh access_token
func (cli *Client) RefreshAccessToken() (string, error) {
	return cli.GetAccessTokenHandler().RefreshToken()
}

// SetAccessToken: set access_token
func (cli *Client) SetAccessToken(token string, expire time.Duration) error {
	return cli.GetAccessTokenHandler().SetToken(token, expire)
}

// Get: request api with get method
func (cli *Client) Get(uri string, args ...map[string]interface{}) (*request.Result, error) {
	return cli.GetRequestClient().Get(uri, args...)
}

// Post: request api with post method
func (cli *Client) Post(uri string, args ...map[string]interface{}) (*request.Result, error) {
	return cli.GetRequestClient().Post(uri, args...)
}

// Request: request api
func (cli *Client) Request(method, uri string, opts goz.Options) (*request.Result, error) {
	return cli.GetRequestClient().Request(method, uri, opts)
}

// AppendAccessTokenToUri: append access_token to request uri
func (cli *Client) AppendAccessTokenToUri(uri string) (string, error) {
	accessToken, err := cli.GetAccessToken()
	if err != nil {
		return uri, fmt.Errorf("%w: %v", token.ErrGetTokenFailed, err)
	}

	u, err := url.Parse(uri)
	if err != nil {
		return uri, err
	}

	q := u.Query()
	q.Add("access_token", accessToken)

	u.RawQuery = q.Encode()
	uri = u.String()

	return uri, nil
}

// AppendAccessTokenToParams: append access_token to request params
func (cli *Client) AppendAccessTokenToParams(params map[string]interface{}) (map[string]interface{}, error) {
	accessToken, err := cli.GetAccessToken()
	if err != nil {
		return params, fmt.Errorf("%w: %v", token.ErrGetTokenFailed, err)
	}

	params["access_token"] = accessToken

	return params, nil
}

// GetWithToken: request api with get method, auto get access_token
func (cli *Client) GetWithToken(uri string, args ...map[string]interface{}) (*request.Result, error) {
	uri, err := cli.AppendAccessTokenToUri(uri)
	if err != nil {
		return nil, err
	}

	return cli.Get(uri, args...)
}

// PostWithToken: request api with post method, auto get access_token
func (cli *Client) PostWithToken(uri string, args ...map[string]interface{}) (*request.Result, error) {
	params := map[string]interface{}{}
	if len(args) > 0 {
		params = args[0]
	}

	newparams, err := cli.AppendAccessTokenToParams(params)
	if err != nil {
		return nil, err
	}
	args[0] = newparams

	return cli.Post(uri, args...)
}

// GetRequestClient: get request handler
func (cli *Client) GetRequestClient() *request.Client {
	return cli.requestClient
}

// GetCacheHandler: get cache handler
func (cli *Client) GetCacheHandler() cache.ICache {
	return cli.cacheHandler
}

// SetCacheHandler: set cache handler
func (cli *Client) SetCacheHandler(handler cache.ICache) error {
	cli.cacheHandler = handler

	return nil
}

// GetAccessTokenCacheKey: get access_token cache key
func (cli *Client) GetAccessTokenCacheKey() string {
	return cli.accessTokenCacheKey
}

// SetAccessTokenCacheKey: set access_token cache key
func (cli *Client) SetAccessTokenCacheKey(cacheKey string) error {
	cli.accessTokenCacheKey = cacheKey

	return nil
}

// GetAccessTokenHandler: get access_token handler
func (cli *Client) GetAccessTokenHandler() token.IToken {
	return cli.accessTokenHandler
}

// SetAccessTokenHandler: set access_token handler
func (cli *Client) SetAccessTokenHandler(handler token.IToken) error {
	cli.accessTokenHandler = handler

	return nil
}

// GetOptions: get options
func (cli *Client) GetOptions() *Options {
	return cli.opts
}

// GetAppid: get appid
func (cli *Client) GetAppid() string {
	return cli.opts.Appid
}

// GetSecret: get secret
func (cli *Client) GetSecret() string {
	return cli.opts.Secret
}
