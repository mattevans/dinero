package dinero

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

const (
	historicalAPIPath = "historical/%s.json"
)

// HistoricalRatesService handles historical rate request/responses.
type HistoricalRatesService struct {
	client       *Client
	baseCurrency string
}

// NewHistoricalRatesService creates a new handler for this service.
func NewHistoricalRatesService(
	client *Client,
	baseCurrency string,
) *HistoricalRatesService {
	return &HistoricalRatesService{
		client:       client,
		baseCurrency: baseCurrency,
	}
}

// HistoricalRatesResponse holds our forex rates for a given base currency
type HistoricalRatesResponse struct {
	Rates     map[string]float64 `json:"rates"`
	Base      string             `json:"base"`
	Timestamp int64              `json:"timestamp"`
}

// List will fetch all the latest rates for the base currency either from the store or the OXR api.
func (s *HistoricalRatesService) List(date time.Time) (*RateResponse, error) {
	// If we have cached results, use them.
	if results, ok := s.client.Cache.Get(s.baseCurrency, date); ok {
		return results, nil
	}

	// No cached results, go and fetch them.
	if err := s.fetch(date); err != nil {
		return nil, err
	}

	return s.List(date)
}

// Get will fetch a single rate for a given currency either from the store or the OXR api.
func (s *HistoricalRatesService) Get(code string, date time.Time) (*float64, error) {
	// No code passed, let them know!
	if code == "" {
		return nil, errors.New("currency code must be passed")
	}

	// If we have cached results, use them.
	if results, ok := s.client.Cache.Get(s.baseCurrency, date); ok {
		if single, ok := results.Rates[code]; ok {
			return &single, nil
		}
		return nil, ErrRatesNotFound
	}

	// No cached results, go and fetch them.
	if err := s.fetch(date); err != nil {
		return nil, err
	}

	return s.Get(code, date)
}

// GetBaseCurrency will return the baseCurrency.
func (s *HistoricalRatesService) GetBaseCurrency() string {
	return s.baseCurrency
}

// SetBaseCurrency will set the base currency to be used for requests.
func (s *HistoricalRatesService) SetBaseCurrency(base string) {
	s.baseCurrency = base
}

func (s *HistoricalRatesService) fetch(date time.Time) error {
	// Build request.
	// add `base` query param if it is not empty
	params := url.Values{}
	if s.baseCurrency != "" {
		params.Set("base", s.baseCurrency)
	}
	request, err := s.client.NewRequest(
		"GET",
		fmt.Sprintf(historicalAPIPath, date.Format("2006-01-02")),
		params,
		nil,
	)
	if err != nil {
		return err
	}

	// Make request
	var latest *RateResponse
	if _, err := s.client.Do(request, &latest); err != nil {
		return err
	}

	s.SetBaseCurrency(latest.Base)

	// Store our results.
	s.client.Cache.Store(latest, date)

	return nil
}
