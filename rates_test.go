package dinero

import (
	"os"
	"reflect"
	"strings"
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

// TestListRatesDefaultCurrency will test updating our local store of forex rates from the OXR API expecting rates for defaultCurrency.
func TestListRatesDefaultCurrency(t *testing.T) {
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

// TestGetRateDefaultCurrency will test pulling a single rate for defaultCurrency.
func TestGetRateDefaultCurrency(t *testing.T) {
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

// TestListRates will test updating our local store of forex rates from the OXR API.
func TestListRates(t *testing.T) {
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

// TestListRates will test rates for the base currency for the given time.Time.
func TestListHistoricalRates(t *testing.T) {
	// Register the test.
	NewWithT(t)

	// Init dinero client.
	client := NewClient(appID, "USD", 1*time.Minute)

	historicalDateStr := "2012-07-10"
	historicalDateTime, _ := time.Parse("2006-01-02", historicalDateStr)

	// Get yesterdays forex rates.
	response, err := client.Rates.ListHistorical(historicalDateTime)
	if err != nil {
		if strings.HasPrefix(err.Error(), setBaseNotAllowedResponsePrefix) {
			t.Skipf("skipping test, unsuitable app ID: %s", err)
		}
		t.Fatalf("Unexpected error running client.Rates.ListHistorical(): %s", err)
	}

	if response.Base != "USD" {
		t.Fatal("Unexpected base oxr rate. Expecting `USD`.")
	}

	if response.Rates == nil {
		t.Fatal("Unexpected length of historical rates")
	}
}

// TestGetRate will test pulling a single rate.
func TestGetRate(t *testing.T) {
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
