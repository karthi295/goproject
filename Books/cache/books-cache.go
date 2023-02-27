package cache

type BookCache interface {
	Set(key string, value string)
	Get(key string) string
}
