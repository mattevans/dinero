package dinero

import (
	"os"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

// TestCurrencies_List will test listing currencies from the OXR api.
func TestCurrencies_List(t *testing.T) {
	// Register the test.
	RegisterTestingT(t)

	// Init dinero client.
	client := NewClient(os.Getenv("OPEN_EXCHANGE_APP_ID"), "AUD", 1*time.Minute)

	// List the currencies
	rsp, err := client.Currencies.List()
	if err != nil {
		t.Fatalf("Unexpected error running client.Currencies.List(): %s", err.Error())
	}

	Expect(err).Should(BeNil())
	Expect(rsp).Should(ContainElement(&CurrencyResponse{
		Code: "AUD",
		Name: "Australian Dollar",
	}))
	Expect(rsp).Should(ContainElement(&CurrencyResponse{
		Code: "NZD",
		Name: "New Zealand Dollar",
	}))
}
