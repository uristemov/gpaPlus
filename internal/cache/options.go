package cache

type Option func(cache *AppCache)

func WithUserCache(user UserCacher) Option {
	return func(cache *AppCache) {
		cache.UserCache = user
	}
}

func WithTokenCache(token TokenCacher) Option {
	return func(cache *AppCache) {
		cache.TokenCache = token
	}
}
