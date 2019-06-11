package dinero

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

// CacheService handles in-memory caching of our rates.
type CacheService struct {
	client *Client
	store  *cache.Cache
}

// NewCacheService creates a new handler for this service.
func NewCacheService(
	client *Client,
	store *cache.Cache,
) *CacheService {
	return &CacheService{
		client,
		store,
	}
}

// Get will return our in-memory stored currency/rates.
func (s *CacheService) Get(base string) (*RateResposne, bool) {
	if x, found := s.store.Get(base); found {
		return x.(*RateResposne), found
	}
	return nil, false
}

// Store will store our currency/rates in-memory.
func (s *CacheService) Store(rsp *RateResposne) {
	// Set a stored timestamp.
	rsp.Timestamp = time.Now().Unix()

	s.store.Set(
		rsp.Base,
		rsp,
		cache.DefaultExpiration,
	)
}

// IsExpired checks whether or not rate stored is expired.
func (s *CacheService) IsExpired(base string) bool {
	if _, found := s.store.Get(base); found {
		return false
	}
	return true
}

// Expire will expire the cache for a given base currency.
func (s *CacheService) Expire(base string) {
	s.store.Delete(base)
}
