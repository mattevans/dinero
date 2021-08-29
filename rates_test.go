package dinero

import (
	"reflect"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

// TestRatesDefaultCurrency_List will test updating our local store of forex rates from the OXR API expecting rates for defaultCurrency.
func TestRatesDefaultCurrency_List(t *testing.T) {
	// Register the test.
	NewWithT(t)

	// Init dinero client.
	client := NewClient(appID, "", 1*time.Minute)

	// Get latest forex rates.
	response, err := client.Rates.List()
	if err != nil {
		t.Fatalf("Unexpected error running client.Rates.List(): %s", err)
	}

	if response.Base != defaultCurrency {
		t.Fatalf("Unexpected base oxr rate: Expecting `%s`.", defaultCurrency)
	}

	if response.Rates == nil {
		t.Fatal("Unexpected nil length of rates")
	}
}

// TestRatesDefaultCurrency_Get will test pulling a single rate for defaultCurrency.
func TestRatesDefaultCurrency_Get(t *testing.T) {
	// Register the test.
	NewWithT(t)

	// Init dinero client.
	client := NewClient(appID, "", 1*time.Minute)

	// Get latest forex rates for NZD (using defaultCurrency as a base).
	response, err := client.Rates.Get("NZD")
	if err != nil {
		t.Fatalf("Unexpected error running client.Rates.Get('NZD'): %s", err)
	}

	// Did we get a *float64 back?
	if reflect.TypeOf(response).String() != "*float64" {
		t.Fatalf("Unexpected rate datatype, expected float64 got %T", response)
	}
}

// TestRates_List will test updating our local store of forex rates from the OXR API.
func TestRates_List(t *testing.T) {
	// Register the test.
	NewWithT(t)

	// Init dinero client.
	client := NewClient(appID, "AUD", 1*time.Minute)

	// Get latest forex rates.
	response, err := client.Rates.List()
	if err != nil {
		if strings.HasPrefix(err.Error(), setBaseNotAllowedResponsePrefix) {
			t.Skipf("skipping test, unsuitable app ID: %s", err)
		}
		t.Fatalf("Unexpected error running client.Rates.List(): %s", err)
	}

	if response.Base != "AUD" {
		t.Fatal("Unexpected base oxr rate. Expecting `AUD`.")
	}

	if response.Rates == nil {
		t.Fatal("Unexpected length of rates")
	}
}

// TestRates_Get will test pulling a single rate.
func TestRates_Get(t *testing.T) {
	// Register the test.
	NewWithT(t)

	// Init dinero client.
	client := NewClient(appID, "AUD", 1*time.Minute)

	// Get latest forex rates for NZD (using AUD as a base).
	response, err := client.Rates.Get("NZD")
	if err != nil {
		if strings.HasPrefix(err.Error(), setBaseNotAllowedResponsePrefix) {
			t.Skipf("skipping test, unsuitable app ID: %s", err)
		}
		t.Fatalf("Unexpected error running client.Rates.Get('NZD'): %s", err)
	}

	// Did we get a *float64 back?
	if reflect.TypeOf(response).String() != "*float64" {
		t.Fatalf("Unexpected rate datatype, expected float64 got %T", response)
	}
}
