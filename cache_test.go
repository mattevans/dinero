package dinero

import (
	"encoding/json"
	"os"
	"testing"

	. "github.com/onsi/gomega"
)

// TestCache will test that our in-memory cache of forex results is working.
func TestCache(t *testing.T) {
	// Register the test.
	RegisterTestingT(t)

	// Init dinero client.
	client := NewClient(os.Getenv("OPEN_EXCHANGE_APP_ID"), "AUD")

	// Get latest forex rates.
	response1, err := client.Rates.List()
	if err != nil {
		t.Fatalf("Unexpected error running client.Rates.All(): %s", err.Error())
	}

	// Expire the cache
	client.Cache.Expire("AUD")

	// Fetch results again
	response2, err := client.Rates.List()
	if err != nil {
		t.Fatalf("Unexpected error running client.Rates.All(): %s", err.Error())
	}

	// Compare the results, they shouldn't match, as timestamp values will differ.
	first, _ := json.Marshal(response1)
	second, _ := json.Marshal(response2)
	Expect(first).NotTo(MatchJSON(second))
}
