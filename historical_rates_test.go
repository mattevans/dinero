package dinero

import (
	"reflect"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

// TestHistoricalRates_List will test updating our local store of forex rates from the OXR API expecting rates for defaultCurrency for a historical date.
func TestHistoricalRates_List(t *testing.T) {
	// Register the test.
	NewWithT(t)

	// Init dinero client.
	client := NewClient(appID, defaultCurrency, 1*time.Minute)

	historicalDate := time.Now().AddDate(0, -2, -5)

	// Get latest forex rates.
	response, err := client.HistoricalRates.List(historicalDate)
	if err != nil {
		t.Fatalf("Unexpected error running client.HistoricalRates.List(): %s", err)
	}

	if response.Base != defaultCurrency {
		t.Fatalf("Unexpected base oxr rate: Expecting `%s`.", defaultCurrency)
	}

	if response.Rates == nil {
		t.Fatal("Unexpected nil length of rates")
	}
}

// TestHistoricalRate_Get will test pulling a single rate for defaultCurrency for a historical date.
func TestHistoricalRate_Get(t *testing.T) {
	// Register the test.
	NewWithT(t)

	// Init dinero client.
	client := NewClient(appID, defaultCurrency, 1*time.Minute)

	historicalDate := time.Now().AddDate(0, -2, -5)

	// Get latest forex rates for NZD (using defaultCurrency as a base).
	response, err := client.HistoricalRates.Get("NZD", historicalDate)
	if err != nil {
		t.Fatalf("Unexpected error running client.HistoricalRates.Get('NZD'): %s", err)
	}

	// Did we get a *float64 back?
	if reflect.TypeOf(response).String() != "*float64" {
		t.Fatalf("Unexpected rate datatype, expected float64 got %T", response)
	}
}
