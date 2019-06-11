package dinero

const (
	currenciesAPIPath = "currencies.json"
)

// CurrenciesService handles currency request/responses.
type CurrenciesService struct {
	client *Client
}

// NewCurrenciesService creates a new handler for this service.
func NewCurrenciesService(
	client *Client,
) *CurrenciesService {
	return &CurrenciesService{
		client,
	}
}

// CurrencyResponse represents a currency from OXR.
type CurrencyResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// List will fetch all list of all currencies available via the OXR api.
func (s *CurrenciesService) List() ([]*CurrencyResponse, error) {
	// Build request.
	req, err := s.client.NewUnauthedRequest(
		"GET",
		currenciesAPIPath,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// Make request.
	rsp := map[string]string{}
	if _, err = s.client.Do(req, &rsp); err != nil {
		return nil, err
	}

	// Parse rsp into slice of *CurrencyResponse's.
	latest := []*CurrencyResponse{}
	for code, name := range rsp {
		latest = append(latest, &CurrencyResponse{
			Code: code,
			Name: name,
		})
	}
	return latest, nil
}
