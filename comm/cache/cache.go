package cache

// ICache 缓存接口
type ICache interface {
	Set(key string, value interface{}) error
	Get(key string) interface{}
}

//
type Cache struct {
}
