package common

import (
	"testing"
)

var _ = func() bool {
	testing.Init()
	return true
}()

func TestGetEndpoint(t *testing.T) {

	expectation := "/api/json?tree=jobs[name,color,buildable,url,description]"
	output := GetEndpoint(Kwargs{"name": "job"})

	if expectation != output {
		t.Errorf("Error")
	}
}
