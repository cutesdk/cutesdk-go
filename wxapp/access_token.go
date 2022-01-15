package wxapp

// AccessTokenHandler 自定义 access_token 处理接口
type AccessTokenHandler interface {
	GetAccessToken() string
}
