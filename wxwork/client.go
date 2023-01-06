package wxwork

import (
	"fmt"
	"net/url"
	"time"

	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/common/token"
	"github.com/idoubi/goz"
)

// Client: wxwork client
type Client struct {
	opts                *Options
	requestClient       *request.Client
	cacheHandler        cache.ICache
	accessTokenCacheKey string
	accessTokenHandler  token.IToken
}

// NewClient: new wxmp client
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
	cli := &Client{opts: opts}

	// new cache handler
	cacheHandler, err := cache.NewCache(opts.Cache)
	if err != nil {
		return nil, fmt.Errorf("new cache handler failed: %v", err)
	}

	// set cache handler
	cli.SetCacheHandler(cacheHandler)

	// set request client
	cli.requestClient = request.NewClient(&request.Options{
		BaseUri: "https://qyapi.weixin.qq.com",
		Debug:   opts.Debug,
		Timeout: opts.Timeout,
	})

	// set default access_token cache key
	accessTokenCacheKey := fmt.Sprintf("wxwork.access_token.%s.%s", cli.GetCorpid(), cli.GetAppid())
	cli.SetAccessTokenCacheKey(accessTokenCacheKey)

	// set default access_token handler
	cli.SetAccessTokenHandler(NewAccessToken(cli))

	return cli, nil
}

// GetAccessToken: get access_token from cache or api
func (ins *Client) GetAccessToken() (string, error) {
	return ins.GetAccessTokenHandler().GetToken()
}

// RefreshAccessToken: refresh access_token
func (ins *Client) RefreshAccessToken() (string, error) {
	return ins.GetAccessTokenHandler().RefreshToken()
}

// SetAccessToken: set access_token
func (ins *Client) SetAccessToken(token string, expire time.Duration) error {
	return ins.GetAccessTokenHandler().SetToken(token, expire)
}

// GetOptions: get options
func (ins *Client) GetOptions() *Options {
	return ins.opts
}

// GetCorpid: get corpid
func (ins *Client) GetCorpid() string {
	return ins.opts.Corpid
}

// GetAppid: get appid
func (ins *Client) GetAppid() string {
	return ins.opts.Appid
}

// Get: request api with get method
func (ins *Client) Get(uri string, args ...map[string]interface{}) (*request.Result, error) {
	return ins.GetRequestClient().Get(uri, args...)
}

// Post: request api with post method
func (ins *Client) Post(uri string, args ...map[string]interface{}) (*request.Result, error) {
	return ins.GetRequestClient().Post(uri, args...)
}

// PostMultipart: request api with post method
func (cli *Client) PostMultipart(uri string, args ...map[string]interface{}) (*request.Result, error) {
	return cli.GetRequestClient().PostMultipart(uri, args...)
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

// GetWithToken: request api with get method, auth get access_token
func (cli *Client) GetWithToken(uri string, args ...map[string]interface{}) (*request.Result, error) {
	uri, err := cli.AppendAccessTokenToUri(uri)
	if err != nil {
		return nil, err
	}

	return cli.Get(uri, args...)
}

// PostWithToken: request api with post method, auto get access_token
func (cli *Client) PostWithToken(uri string, args ...map[string]interface{}) (*request.Result, error) {
	uri, err := cli.AppendAccessTokenToUri(uri)
	if err != nil {
		return nil, err
	}

	return cli.Post(uri, args...)
}

// PostMultipartWithToken: request api with post method, auto get access_token
func (cli *Client) PostMultipartWithToken(uri string, args ...map[string]interface{}) (*request.Result, error) {
	uri, err := cli.AppendAccessTokenToUri(uri)
	if err != nil {
		return nil, err
	}

	return cli.PostMultipart(uri, args...)
}

// Request: request api
func (ins *Client) Request(method, uri string, opts goz.Options) (*request.Result, error) {
	return ins.GetRequestClient().Request(method, uri, opts)
}

// GetRequestClient: get request handler
func (ins *Client) GetRequestClient() *request.Client {
	return ins.requestClient
}

// GetCacheHandler: get cache handler
func (ins *Client) GetCacheHandler() cache.ICache {
	return ins.cacheHandler
}

// SetCacheHandler: set cache handler
func (ins *Client) SetCacheHandler(handler cache.ICache) error {
	ins.cacheHandler = handler

	return nil
}

// GetAccessTokenCacheKey: get access_token cache key
func (ins *Client) GetAccessTokenCacheKey() string {
	return ins.accessTokenCacheKey
}

// SetAccessTokenCacheKey: set access_token cache key
func (ins *Client) SetAccessTokenCacheKey(cacheKey string) error {
	ins.accessTokenCacheKey = cacheKey

	return nil
}

// GetAccessTokenHandler: get access_token handler
func (ins *Client) GetAccessTokenHandler() token.IToken {
	return ins.accessTokenHandler
}

// SetAccessTokenHandler: set access_token handler
func (ins *Client) SetAccessTokenHandler(handler token.IToken) error {
	ins.accessTokenHandler = handler

	return nil
}
