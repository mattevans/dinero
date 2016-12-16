package dinero

import (
	"errors"
	"fmt"
	"time"
)

const (
	latestAPIPath  = "latest.json"
	historyAPIPath = "history.json"
)

// UpdateService handles communication with /latest.json methods of the
// https://docs.openexchangerates.org/docs/latest-json
type UpdateService service

// UpdateResponse is the set of parameters that is used in response to a send
// request
type UpdateResponse struct {
	Timestamp int64
	UpdatedAt time.Time
	Base      string
	Rates     map[string]float64
}

// LatestRates will build and execute request to /latest.json using the base
// currency provided.
func (s *UpdateService) LatestRates(base string) error {
	if base == "" {
		return errors.New("The base currency provided cannot be empty")
	}

	// Append our base currency to request URL.
	rateAPIPath := fmt.Sprintf("%s?base=%s", latestAPIPath, base)

	// Build request.
	request, err := s.client.NewRequest("GET", rateAPIPath, nil)
	if err != nil {
		return err
	}

	// Make request
	latest := &UpdateResponse{}
	_, err = s.client.Do(request, latest)
	if err != nil {
		return err
	}

	// Switch our unix timestamp to time.Time.
	latest.UpdatedAt = time.Unix(latest.Timestamp, 0).UTC()

	// Cache our latest results.
	s.client.Cache.Store(base, latest.Rates)

	return nil
}
