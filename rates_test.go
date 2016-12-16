package dinero

import (
	"os"
	"reflect"
	"testing"

	. "github.com/onsi/gomega"
)

// TestAllRates will test updating our local store of forex rates from the OXR API.
func TestAllRates(t *testing.T) {
	// Register the test.
	RegisterTestingT(t)

	// Init dinero client.
	client := NewClient(os.Getenv("OPEN_EXCHANGE_APP_ID"))

	// Set a base currency to work with.
	client.Rates.SetBaseCurrency("AUD")

	// Get latest forex rates.
	response, err := client.Rates.All()
	if err != nil {
		t.Fatalf("Unexpected error running client.Rates.All(): %s", err.Error())
	}

	if response.Base != "AUD" {
		t.Fatalf("Unexpected base oxr rate: %s. Expecting `AUD`.", err.Error())
	}

	if response.UpdatedAt.IsZero() {
		t.Fatalf("Unexpected response timestamp: %s.", err.Error())
	}

	if response.Rates == nil {
		t.Fatalf("Unexpected length of rates: %s.", err.Error())
	}
}

// TestSingleRate will test pulling a single rate.
func TestSingleRate(t *testing.T) {
	// Register the test.
	RegisterTestingT(t)

	// Init dinero client.
	client := NewClient(os.Getenv("OPEN_EXCHANGE_APP_ID"))

	// Set a base currency to work with.
	client.Rates.SetBaseCurrency("AUD")

	// Get latest forex rates for NZD (using AUD as a base).
	response, err := client.Rates.Single("NZD")
	if err != nil {
		t.Fatalf("Unexpected error running client.Rates.Single('NZD'): %s", err.Error())
	}

	// Did we get a *float64 back?
	if reflect.TypeOf(response).String() != "*float64" {
		t.Fatalf("Unexpected rate datatype, expected float64 got %T", response)
	}
}
