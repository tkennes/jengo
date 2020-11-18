package jengo_src

import (
	"testing"
)

func TestGetEndpoint_jobs(t *testing.T) {

	expectation := "/api/json?tree=jobs[name,color,buildable,url,description]"
	result := GetEndpoint(Kwargs{"name": "jobs"})

	if expectation != result {
		t.Errorf("Expectation: %s\nResult:      %s", expectation, result)
	}
}

func TestGetEndpoint_job(t *testing.T) {

	expectation := "/job/TEST/api/json?pretty=true"
	result := GetEndpoint(Kwargs{"name": "job", "job_name": "TEST"})

	if expectation != result {
		t.Errorf("Expectation: %s\nResult:      %s", expectation, result)
	}
}

func TestGetEndpoint_builds(t *testing.T) {

	expectation := "/job/TEST/api/json?tree=builds[number,status,timestamp,id,result,estimatedDuration,duration,executor,description,url]"
	result := GetEndpoint(Kwargs{"name": "builds", "job_name": "TEST"})

	if expectation != result {
		t.Errorf("Expectation: %s\nResult:      %s", expectation, result)
	}
}
func TestGetEndpoint_build(t *testing.T) {

	expectation := "/job/TEST/TEST/api/json?pretty=true"
	result := GetEndpoint(Kwargs{"name": "build", "job_name": "TEST", "build_name": "TEST"})

	if expectation != result {
		t.Errorf("Expectation: %s\nResult:      %s", expectation, result)
	}
}
