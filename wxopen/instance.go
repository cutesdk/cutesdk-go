package wxopen

import (
	"fmt"
	"time"

	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/common/token"
	"github.com/idoubi/goutils/crypt"
	"github.com/idoubi/goz"
)

// Instance: wxopen instance
type Instance struct {
	opts                          *Options
	requestClient                 *request.Client
	cacheHandler                  cache.ICache
	componentAccessTokenCacheKey  string
	componentAccessTokenHandler   token.IToken
	componentVerifyTicketCacheKey string
	componentVerifyTicketHandler  token.IToken
}

// New create wxopen instance
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

	// set request Instance
	ins.requestClient = request.NewClient(opts.Request)

	// set default component_access_token cache key
	componentAccessTokenCacheKey := fmt.Sprintf("wxopen.component_access_token.%s", ins.GetComponentAppid())
	ins.SetComponentAccessTokenCacheKey(componentAccessTokenCacheKey)

	// set default component_access_token handler
	ins.SetComponentAccessTokenHandler(NewComponentAccessToken(ins))

	// set default component_verify_ticket cache key
	componentVerifyTicketCacheKey := fmt.Sprintf("wxopen.component_verify_ticket.%s", ins.GetComponentAppid())
	ins.SetComponentVerifyTicketCacheKey(componentVerifyTicketCacheKey)

	// set default component_verify_ticket handler
	ins.SetComponentVerifyTicketHandler(NewComponentVerifyTicket(ins))

	return ins, nil
}

// GetComponentAccessToken: get component_access_token from cache or api
func (ins *Instance) GetComponentAccessToken() (string, error) {
	return ins.GetComponentAccessTokenHandler().GetToken()
}

// RefreshComponentAccessToken: refresh component_access_token
func (ins *Instance) RefreshComponentAccessToken() (string, error) {
	return ins.GetComponentAccessTokenHandler().RefreshToken()
}

// SetComponentAccessToken: set component_access_token
func (ins *Instance) SetComponentAccessToken(token string, expire time.Duration) error {
	return ins.GetComponentAccessTokenHandler().SetToken(token, expire)
}

// GetComponentVerifyTicket: get component_verify_ticket from cache
func (ins *Instance) GetComponentVerifyTicket() (string, error) {
	return ins.GetComponentVerifyTicketHandler().GetToken()
}

// RefreshComponentVerifyTicket: refresh component_verify_ticket
func (ins *Instance) RefreshComponentVerifyTicket() (string, error) {
	return ins.GetComponentVerifyTicketHandler().RefreshToken()
}

// SetComponentVerifyTicket: set component_verify_ticket
func (ins *Instance) SetComponentVerifyTicket(token string, expire time.Duration) error {
	return ins.GetComponentVerifyTicketHandler().SetToken(token, expire)
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

// GetComponentAppid: get component_appid
func (ins *Instance) GetComponentAppid() string {
	return ins.opts.ComponentAppid
}

// GetComponentAppsecret get component_appsecret
func (ins *Instance) GetComponentAppsecret() string {
	return ins.opts.ComponentAppsecret
}

// GetComponentAccessTokenCacheKey: get component_access_token cache key
func (ins *Instance) GetComponentAccessTokenCacheKey() string {
	return ins.componentAccessTokenCacheKey
}

// SetComponentAccessTokenCacheKey: set component_access_token cache key
func (ins *Instance) SetComponentAccessTokenCacheKey(cacheKey string) error {
	ins.componentAccessTokenCacheKey = cacheKey

	return nil
}

// GetComponentAccessTokenHandler: get component_access_token handler
func (ins *Instance) GetComponentAccessTokenHandler() token.IToken {
	return ins.componentAccessTokenHandler
}

// SetComponentAccessTokenHandler: set component_access_token handler
func (ins *Instance) SetComponentAccessTokenHandler(handler token.IToken) error {
	ins.componentAccessTokenHandler = handler

	return nil
}

// GetComponentVerifyTicketCacheKey: get component_verify_ticket cache key
func (ins *Instance) GetComponentVerifyTicketCacheKey() string {
	return ins.componentVerifyTicketCacheKey
}

// SetComponentVerifyTicketCacheKey: set component_verify_ticket cache key
func (ins *Instance) SetComponentVerifyTicketCacheKey(cacheKey string) error {
	ins.componentVerifyTicketCacheKey = cacheKey

	return nil
}

// GetComponentVerifyTicketHandler: get component_verify_ticket handler
func (ins *Instance) GetComponentVerifyTicketHandler() token.IToken {
	return ins.componentVerifyTicketHandler
}

// SetComponentVerifyTicketHandler: set component_verify_ticket handler
func (ins *Instance) SetComponentVerifyTicketHandler(handler token.IToken) error {
	ins.componentVerifyTicketHandler = handler

	return nil
}
