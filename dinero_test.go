package dinero

import (
	"net/url"
	"os"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

const (
	// Default currency returned by API calls that do not set a base currency.
	defaultCurrency = "USD"
	// "Free" OXR plans don't allow switching of base currency.
	// > 403 Changing the API `base` currency is available for Developer, Enterprise and Unlimited plan clients.
	setBaseNotAllowedResponsePrefix = "403"
)

var appID = os.Getenv("OPEN_EXCHANGE_APP_ID")

func Test_NewRequest(t *testing.T) {
	// Register the test.
	NewWithT(t)

	// Init dinero client.
	client := NewClient("12345", "", 1*time.Minute)

	var tests = []struct {
		params      url.Values
		expectedURL string
	}{
		{
			params:      url.Values{},
			expectedURL: "https://openexchangerates.org/api/latest.json?app_id=12345",
		},
		{
			params:      url.Values{"base": []string{"AUD"}},
			expectedURL: "https://openexchangerates.org/api/latest.json?app_id=12345&base=AUD",
		},
	}

	for _, test := range tests {
		req, err := client.NewRequest("GET", "latest.json", test.params, nil)
		if err != nil {
			t.Fatal(err)
		}
		actual := req.URL.String()
		if actual != test.expectedURL {
			t.Errorf("expected %s, got %s", test.expectedURL, actual)
		}
	}
}
