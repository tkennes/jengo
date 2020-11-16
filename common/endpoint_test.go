package common

import (
	"testing"
)

func TestGetEndpoint(t *testing.T) {

	expectation := "/api/json?tree=jobs[name,color,buildable,url,description]"
	output := GetEndpoint(Kwargs{"name": "job"})

	if expectation != output {
		t.Errorf("Error")
	}
}
