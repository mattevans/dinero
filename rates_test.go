package dinero

import (
	"os"
	"reflect"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

// TestListRates will test updating our local store of forex rates from the OXR API.
func TestListRates(t *testing.T) {
	// Register the test.
	RegisterTestingT(t)

	// Init dinero client.
	client := NewClient(os.Getenv("OPEN_EXCHANGE_APP_ID"), "AUD", 1*time.Minute)

	// Get latest forex rates.
	response, err := client.Rates.List()
	if err != nil {
		t.Fatalf("Unexpected error running client.Rates.List(): %s", err.Error())
	}

	if response.Base != "AUD" {
		t.Fatalf("Unexpected base oxr rate: %s. Expecting `AUD`.", err.Error())
	}

	if response.Rates == nil {
		t.Fatalf("Unexpected length of rates: %s.", err.Error())
	}
}

// TestGetRate will test pulling a single rate.
func TestGetRate(t *testing.T) {
	// Register the test.
	RegisterTestingT(t)

	// Init dinero client.
	client := NewClient(os.Getenv("OPEN_EXCHANGE_APP_ID"), "AUD", 1*time.Minute)

	// Get latest forex rates for NZD (using AUD as a base).
	response, err := client.Rates.Get("NZD")
	if err != nil {
		t.Fatalf("Unexpected error running client.Rates.Get('NZD'): %s", err.Error())
	}

	// Did we get a *float64 back?
	if reflect.TypeOf(response).String() != "*float64" {
		t.Fatalf("Unexpected rate datatype, expected float64 got %T", response)
	}
}
