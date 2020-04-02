package dinero

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

// TestCache will test that our in-memory cache of forex results is working.
func TestCache(t *testing.T) {
	// Register the test.
	NewWithT(t)

	// Init dinero client.
	client := NewClient(appID, "AUD", 1*time.Minute)

	// Get latest forex rates.
	response1, err := client.Rates.List()
	if err != nil {
		if strings.HasPrefix(err.Error(), setBaseNotAllowedResponsePrefix) {
			t.Skipf("skipping test, unsuitable app ID: %s", err)
		}
		t.Fatalf("Unexpected error running client.Rates.List(): %s", err.Error())
	}

	// Fetch results again
	response2, ok := client.Cache.Get("AUD")
	if !ok {
		t.Fatalf("Expected response when fetching from cache for base currency AUD, got: %v", response2)
	}

	first, _ := json.Marshal(response1)
	second, _ := json.Marshal(response2)
	Expect(first).To(MatchJSON(second))

	// Expire the cache
	client.Cache.Expire("AUD")

	// Fetch results again (from the cache), now it's cleared.
	response2, _ = client.Cache.Get("AUD")

	// Should be nothing.
	Expect(response2).Should(BeNil())
}
