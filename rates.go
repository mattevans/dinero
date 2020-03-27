package dinero

import (
	"errors"
	"fmt"
)

const (
	latestAPIPath = "latest.json"
)

var (
	// ErrRatesNotFound is returned if no rate can be found for a given currency code.
	ErrRatesNotFound = errors.New("no rates found for code")
)

// RatesService handles rate request/responses.
type RatesService struct {
	client       *Client
	baseCurrency string
}

// NewRatesService creates a new handler for this service.
func NewRatesService(
	client *Client,
	baseCurrency string,
) *RatesService {
	return &RatesService{
		client,
		baseCurrency,
	}
}

// RateResponse holds our forex rates for a given base currency
type RateResponse struct {
	Rates     map[string]float64 `json:"rates"`
	Base      string             `json:"base"`
	Timestamp int64              `json:"timestamp"`
}

// List will fetch all the latest rates for the base currency either from
// the the store or the OXR api.
func (s *RatesService) List() (*RateResponse, error) {
	// If we have cached results, use them.
	if results, ok := s.client.Cache.Get(s.baseCurrency); ok {
		return results, nil
	}

	// No cached results, go and fetch them.
	if err := s.Update(s.baseCurrency); err != nil {
		return nil, err
	}

	return s.List()
}

// Get will fetch a single rate for a given currency either from
// the the store or the OXR api.
func (s *RatesService) Get(code string) (*float64, error) {
	// No code passed, let them know!
	if code == "" {
		return nil, errors.New("currency code must be passed")
	}

	// If we have cached results, use them.
	if results, ok := s.client.Cache.Get(s.baseCurrency); ok {
		if single, ok := results.Rates[code]; ok {
			return &single, nil
		}
		return nil, ErrRatesNotFound
	}

	// No cached results, go and fetch them.
	if err := s.Update(s.baseCurrency); err != nil {
		return nil, err
	}

	return s.Get(code)
}

// Update will update the rates for the given currency from OXR.
func (s *RatesService) Update(base string) error {
	if base == "" {
		return errors.New("base currency provided cannot be empty")
	}

	// Build request.
	request, err := s.client.NewRequest(
		"GET",
		fmt.Sprintf("%s?base=%s", latestAPIPath, base),
		nil,
	)
	if err != nil {
		return err
	}

	// Make request
	latest := &RateResponse{}
	if _, err = s.client.Do(request, latest); err != nil {
		return err
	}

	// Store our results.
	s.client.Cache.Store(latest)

	return nil
}

// GetBaseCurrency will return the baseCurrency.
func (s *RatesService) GetBaseCurrency() string {
	return s.baseCurrency
}

// SetBaseCurrency will set the base currency to be used for requests.
func (s *RatesService) SetBaseCurrency(base string) {
	s.baseCurrency = base
}
