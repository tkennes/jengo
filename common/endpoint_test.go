package common

import (
	"fmt"
	"testing"
)

type kwargs map[string]interface{}

func Test_get_endpoint(t *testing.T) {
	input := kwargs{"name": "jobs"}
	expectation := "/api/json?tree=jobs[name,color,buildable,url,description]"
	result := common.get_endpoint(input)
	if result != expectation {
		t.Errorf(fmt.Sprint("Expectation: %s\nResult:     %s", expectation, result))
	}
}
