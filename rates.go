package dinero

import (
	"errors"
	"time"
)

// RatesService handles retrieving forex rates for a given currency, either from
// in-memory cache or by fetching fresh results.
type RatesService service

// RatesStore holds our forex rates for a given base currency
type RatesStore struct {
	Rates     map[string]float64 `json:"rates"`
	UpdatedAt *time.Time         `json:"updated_at"`
	Base      string             `json:"base"`
}

var baseCurrency string

// GetBaseCurrency will return the baseCurrency.
func (s *RatesService) GetBaseCurrency() string {
	return baseCurrency
}

// SetBaseCurrency will set the base currency to be used for requests.
func (s *RatesService) SetBaseCurrency(base string) {
	baseCurrency = base
}

// All will build and execute request to fetch the latest rates for given base
// currency either from the in-memory cache or OXR API.
func (s *RatesService) All() (*RatesStore, error) {
	// No base currency provided, let them know!
	if baseCurrency == "" {
		return nil, errors.New("please set a base currency")
	}

	// If we have cached results, use them.
	results := s.client.Cache.Get(baseCurrency)
	if results != nil {
		return results, nil
	}

	// No cached results, go and fetch them.
	err := s.client.Update.LatestRates(baseCurrency)
	if err != nil {
		return nil, err
	}

	return s.All()
}

// Single will return forex rate for given base/code.
func (s *RatesService) Single(code string) (*float64, error) {
	// No base currency provided, let them know!
	if baseCurrency == "" || code == "" {
		return nil, errors.New("both the base currency and requested currency values must be set")
	}

	// If we have cached results, use them.
	results := s.client.Cache.Get(baseCurrency)
	if results != nil {
		single := results.Rates[code]
		return &single, nil
	}

	// No cached results, go and fetch them.
	err := s.client.Update.LatestRates(baseCurrency)
	if err != nil {
		return nil, err
	}

	return s.Single(code)
}
