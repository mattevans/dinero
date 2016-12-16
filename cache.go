package dinero

import "time"

// cache holds our our cached forex results for varying bases.
var cache map[string]*RatesStore

// cacheTTL stores the time to live of our cache (2 hours).
var cacheTTL = 2 * time.Hour

// CacheService handles in-memory caching of our exchange rates.
type CacheService service

// Get will return our stored in-memory forex rates, if we have them.
func (s *CacheService) Get(base string) *RatesStore {
	// Is our cache expired?
	if s.IsExpired(base) {
		return nil
	}

	// Use stored results.
	return cache[base]
}

// Store will save our forex rates to a RatesStore.
func (s *CacheService) Store(base string, rates map[string]float64) {
	// No cache? Initalize it.
	if cache == nil {
		cache = map[string]*RatesStore{}
	}

	// Store
	tn := time.Now()
	cache[base] = &RatesStore{
		Rates:     rates,
		UpdatedAt: &tn,
		Base:      base,
	}
}

// IsExpired checks if we have stored cache and that it isn't expired.
func (s *CacheService) IsExpired(base string) bool {
	// No cache? bail.
	if cache[base] == nil || (len(cache[base].Rates) <= 0) {
		return true
	}

	// Expired cache? bail.
	lastUpdated := cache[base].UpdatedAt
	if lastUpdated != nil && lastUpdated.Add(cacheTTL).Before(time.Now()) {
		return true
	}

	return false
}

// Expire will expire the cache for a given base currency.
func (s *CacheService) Expire(base string) {
	cache[base] = nil
}
