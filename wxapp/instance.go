package wxapp

import (
	"fmt"
	"time"

	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/common/token"
	"github.com/idoubi/goz"
)

// Instance: wxapp instance
type Instance struct {
	opts                *Options
	requestClient       *request.Client
	cacheHandler        cache.ICache
	accessTokenCacheKey string
	accessTokenHandler  token.IToken
}

// New create wxapp instance
func New(opts *Options) (*Instance, error) {
	// set default options
	if opts.Request == nil {
		opts.Request = &request.Options{
			Debug: opts.Debug,
		}
	}
	if opts.Cache == nil {
		opts.Cache = &cache.Options{}
	}
	if opts.Request.BaseUri == "" {
		opts.Request.BaseUri = "https://api.weixin.qq.com"
	}
	if opts.Request.Timeout == 0 {
		opts.Request.Timeout = 5 * time.Second
	}
	if opts.Cache.Driver == "" {
		opts.Cache.Driver = "file"
		opts.Cache.Conf = map[string]interface{}{
			"dir": "./cache",
		}
	}

	// new instance
	ins := &Instance{opts: opts}

	// set cache handler
	if cache, err := cache.NewCache(opts.Cache); err != nil {
		return nil, fmt.Errorf("new cache handler failed: " + err.Error())
	} else {
		ins.SetCacheHandler(cache)
	}

	// set request client
	ins.requestClient = request.NewClient(opts.Request)

	// set default access_token cache key
	accessTokenCacheKey := fmt.Sprintf("wxapp.access_token.%s", ins.opts.Appid)
	ins.SetAccessTokenCacheKey(accessTokenCacheKey)

	// set default access_token handler
	ins.SetAccessTokenHandler(NewAccessToken(ins))

	return ins, nil
}

// GetRequestClient: get request handler
func (ins *Instance) GetRequestClient() *request.Client {
	return ins.requestClient
}

// GetCacheHandler: get cache handler
func (ins *Instance) GetCacheHandler() cache.ICache {
	return ins.cacheHandler
}

// SetCacheHandler: set cache handler
func (ins *Instance) SetCacheHandler(handler cache.ICache) error {
	ins.cacheHandler = handler

	return nil
}

// GetOptions: get options
func (ins *Instance) GetOptions() *Options {
	return ins.opts
}

// GetAppid: get appid
func (ins *Instance) GetAppid() string {
	return ins.opts.Appid
}

// GetSecret: get secret
func (ins *Instance) GetSecret() string {
	return ins.opts.Secret
}

// GetAccessTokenCacheKey: get access_token cache key
func (ins *Instance) GetAccessTokenCacheKey() string {
	return ins.accessTokenCacheKey
}

// SetAccessTokenCacheKey: set access_token cache key
func (ins *Instance) SetAccessTokenCacheKey(cacheKey string) error {
	ins.accessTokenCacheKey = cacheKey

	return nil
}

// GetAccessTokenHandler: get access_token handler
func (ins *Instance) GetAccessTokenHandler() token.IToken {
	return ins.accessTokenHandler
}

// SetAccessTokenHandler: set access_token handler
func (ins *Instance) SetAccessTokenHandler(handler token.IToken) error {
	ins.accessTokenHandler = handler

	return nil
}

// GetAccessToken: get access_token from cache or api
func (ins *Instance) GetAccessToken() (string, error) {
	return ins.GetAccessTokenHandler().GetToken()
}

// RefreshAccessToken: refresh access_token
func (ins *Instance) RefreshAccessToken() (string, error) {
	return ins.GetAccessTokenHandler().RefreshToken()
}

// SetAccessToken: set access_token
func (ins *Instance) SetAccessToken(token string, expire time.Duration) error {
	return ins.GetAccessTokenHandler().SetToken(token, expire)
}

// Get: request api with get method
func (ins *Instance) Get(uri string, args ...map[string]interface{}) (*request.Result, error) {
	return ins.GetRequestClient().Get(uri, args...)
}

// Post: request api with post method
func (ins *Instance) Post(uri string, args ...map[string]interface{}) (*request.Result, error) {
	return ins.GetRequestClient().Post(uri, args...)
}

// Request: request api
func (ins *Instance) Request(method, uri string, opts goz.Options) (*request.Result, error) {
	return ins.GetRequestClient().Request(method, uri, opts)
}
