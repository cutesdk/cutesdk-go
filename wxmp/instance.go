package wxmp

import (
	"fmt"
	"net/url"
	"time"

	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/common/token"
	"github.com/idoubi/goutils/crypt"
	"github.com/idoubi/goz"
)

// Instance: wxmp instance
type Instance struct {
	opts                *Options
	requestClient       *request.Client
	cacheHandler        cache.ICache
	accessTokenCacheKey string
	accessTokenHandler  token.IToken
	jsapiTicketCacheKey string
	jsapiTicketHandler  token.IToken
}

// New create wxmp instance
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
	if opts.EncodingAesKey != "" {
		if v, _ := crypt.Base64Decode(opts.EncodingAesKey + "="); v != nil {
			opts.aesKey = v
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
	accessTokenCacheKey := fmt.Sprintf("wxmp.access_token.%s", ins.GetAppid())
	ins.SetAccessTokenCacheKey(accessTokenCacheKey)

	// set default access_token handler
	ins.SetAccessTokenHandler(NewAccessToken(ins))

	// set default jsapi_ticket cache key
	jsapiTicketCacheKey := fmt.Sprintf("wxmp.jsapi_ticket.%s", ins.GetAppid())
	ins.SetJsapiTicketCacheKey(jsapiTicketCacheKey)

	// set default jsapi_ticket handler
	ins.SetJsapiTicketHandler(NewJsapiTicket(ins))

	return ins, nil
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

// GetJsapiTicket: get jsapi_ticket from cache or api
func (ins *Instance) GetJsapiTicket() (string, error) {
	return ins.GetJsapiTicketHandler().GetToken()
}

// RefreshJsapiTicket: refresh jsapi_ticket
func (ins *Instance) RefreshJsapiTicket() (string, error) {
	return ins.GetJsapiTicketHandler().RefreshToken()
}

// SetJsapiTicket: set jsapi_ticket
func (ins *Instance) SetJsapiTicket(token string, expire time.Duration) error {
	return ins.GetJsapiTicketHandler().SetToken(token, expire)
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

// AppendAccessTokenToUri: append access_token to request uri
func (ins *Instance) AppendAccessTokenToUri(uri string) (string, error) {
	accessToken, err := ins.GetAccessToken()
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

// GetWithAccessToken: request api with get method, auth get access_token
func (ins *Instance) GetWithAccessToken(uri string, args ...map[string]interface{}) (*request.Result, error) {
	uri, err := ins.AppendAccessTokenToUri(uri)
	if err != nil {
		return nil, err
	}

	return ins.Get(uri, args...)
}

// PostWithAccessToken: request api with post method, auto get access_token
func (ins *Instance) PostWithAccessToken(uri string, args ...map[string]interface{}) (*request.Result, error) {
	uri, err := ins.AppendAccessTokenToUri(uri)
	if err != nil {
		return nil, err
	}

	return ins.Post(uri, args...)
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

// GetJsapiTicketCacheKey: get jsapi_ticket cache key
func (ins *Instance) GetJsapiTicketCacheKey() string {
	return ins.jsapiTicketCacheKey
}

// SetJsapiTicketCacheKey: set jsapi_ticket cache key
func (ins *Instance) SetJsapiTicketCacheKey(cacheKey string) error {
	ins.jsapiTicketCacheKey = cacheKey

	return nil
}

// GetJsapiTicketHandler: get jsapi_ticket handler
func (ins *Instance) GetJsapiTicketHandler() token.IToken {
	return ins.jsapiTicketHandler
}

// SetJsapiTicketHandler: set jsapi_ticket handler
func (ins *Instance) SetJsapiTicketHandler(handler token.IToken) error {
	ins.jsapiTicketHandler = handler

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