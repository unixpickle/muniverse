package chrome

import (
	"os"
	"testing"
)

func testingHost(t *testing.T) string {
	host := os.Getenv("CHROME_DEVTOOLS_HOST")
	if host == "" {
		t.Fatal("no CHROME_DEVTOOLS_HOST environment variable")
	}
	return host
}
