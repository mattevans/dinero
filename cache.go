package dinero

import (
	"fmt"
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
func (s *CacheService) Get(base string, date time.Time) (*RateResponse, bool) {
	if x, found := s.store.Get(getCacheKey(base, date)); found {
		return x.(*RateResponse), found
	}
	return nil, false
}

// Store will store our currency/rates in-memory.
func (s *CacheService) Store(rsp *RateResponse, date time.Time) {
	// Set a stored timestamp.
	rsp.Timestamp = time.Now().Unix()

	s.store.Set(
		getCacheKey(rsp.Base, date),
		rsp,
		cache.DefaultExpiration,
	)
}

// IsExpired checks whether the rate stored is expired.
func (s *CacheService) IsExpired(base string, date time.Time) bool {
	if _, found := s.store.Get(getCacheKey(base, date)); found {
		return false
	}
	return true
}

// Expire will expire the cache for a given base currency.
func (s *CacheService) Expire(base string, date time.Time) {
	s.store.Delete(getCacheKey(base, date))
}

func getCacheKey(base string, date time.Time) string {
	return fmt.Sprintf("%s_%s", base, date.Format("2006-01-02"))
}
