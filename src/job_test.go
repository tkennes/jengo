package jengo_src

import (
	"testing"
)

func TestPrepareUrl(t *testing.T) {
	expectation := "TEST/I/AM/A/URL"
	test_url := GetBaseURL() + "TEST/I/AM/A/URL"

	result := PrepareUrl(test_url)
	if result != expectation {
		t.Errorf("Expectation: %s\nResult:      %s", expectation, result)
	}
}
