package jengo_src

import (
	"testing"
)

func TestGetEndpoint_jobs(t *testing.T) {

	expectation := "/api/json?tree=jobs[name,color,buildable,url,description]"
	output := GetEndpoint(Kwargs{"name": "jobs"})

	if expectation != output {
		t.Errorf("Expectation: %s\nResult:      %s", expectation, output)
	}
}

func TestGetEndpoint_job(t *testing.T) {

	expectation := "/job/TEST/api/json?pretty=true"
	output := GetEndpoint(Kwargs{"name": "job_info", "job_name": "TEST"})

	if expectation != output {
		t.Errorf("Expectation: %s\nResult:      %s", expectation, output)
	}
}

func TestGetEndpoint_builds(t *testing.T) {

	expectation := "/job/TEST/api/json?tree=builds[number,status,timestamp,id,result,estimatedDuration,duration,executor,description,url]"
	output := GetEndpoint(Kwargs{"name": "builds", "job_name": "TEST"})

	if expectation != output {
		t.Errorf("Expectation: %s\nResult:      %s", expectation, output)
	}
}
func TestGetEndpoint_build(t *testing.T) {

	expectation := "/job/TEST/TEST/api/json?pretty=true"
	output := GetEndpoint(Kwargs{"name": "build_info", "job_name": "TEST", "build_name": "TEST"})

	if expectation != output {
		t.Errorf("Expectation: %s\nResult:      %s", expectation, output)
	}
}
