package jengo_src

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

type JobL struct {
	Name        string `yaml:"name"`
	Color       string `yaml:"color"`
	Buildable   bool   `yaml:"buildable"`
	URL         string `yaml:"url"`
	Description string `yaml:"description"`
}

type JobRawResponse struct {
	Class       string `yaml:"_class"`
	Description string `yaml:"description"`
	FullName    string `yaml:"fullName"`
	Name        string `yaml:"name"`
	URL         string `yaml:"url"`
	Buildable   bool   `yaml:"buildable"`
	Builds      []struct {
		Class  string `yaml:"_class"`
		Number int    `yaml:"number"`
		URL    string `yaml:"url"`
	} `yaml:"builds"`
	Color        string `yaml:"color"`
	HealthReport []struct {
		Description string `yaml:"description"`
		Score       int    `yaml:"score"`
	} `yaml:"healthReport"`
	InQueue         bool `yaml:"inQueue"`
	NextBuildNumber int  `yaml:"nextBuildNumber"`
	ConcurrentBuild bool `yaml:"concurrentBuild"`
	ResumeBlocked   bool `yaml:"resumeBlocked"`
}

type JobResponse struct {
	FullName         string `yaml:"fullName"`
	Name             string `yaml:"name"`
	Description      string `yaml:"description"`
	URL              string `yaml:"url"`
	Buildable        bool   `yaml:"buildable"`
	Color            string `yaml:"color"`
	Builds           int    `yaml:"builds"`
	LastHealthReport struct {
		Description string `yaml:"description"`
		Score       int    `yaml:"score"`
	} `yaml:"healthReport"`
	InQueue         bool `yaml:"inQueue"`
	NextBuildNumber int  `yaml:"nextBuildNumber"`
	ConcurrentBuild bool `yaml:"concurrentBuild"`
	ResumeBlocked   bool `yaml:"resumeBlocked"`
}

type JobListResponse struct {
	Jobs []JobL `yaml:"jobs"`
}

func ListJobs() (res [][]string) {
	responseData, err := HandleRequest("GET", Kwargs{"name": "jobs"})
	if err != nil {
		log.Fatal(err)
	}

	var obj JobListResponse
	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		log.Fatal(err)
	}
	for _, job := range obj.Jobs {
		res = append(res, []string{job.Name,
			job.Color, strconv.FormatBool(job.Buildable),
			PrepareUrl(job.URL), job.Description})
	}
	return
}

func GetJob(job_name string) (out_obj JobResponse) {
	var obj JobRawResponse
	responseData, err := HandleRequest("GET", Kwargs{"name": "job_info", "job_name": job_name})
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		log.Fatal(err)
	}
	out_obj = prepare_job_response(obj)
	return
}

func prepare_job_response(obj JobRawResponse) (out_obj JobResponse) {
	out_obj.Description = obj.Description
	out_obj.FullName = obj.FullName
	out_obj.Name = obj.Name
	out_obj.URL = obj.URL
	out_obj.Builds = len(obj.Builds)
	out_obj.Buildable = obj.Buildable
	out_obj.Color = obj.Color
	out_obj.InQueue = obj.InQueue
	if len(obj.HealthReport) > 0 {
		out_obj.LastHealthReport = obj.HealthReport[0]
	}
	out_obj.NextBuildNumber = obj.NextBuildNumber
	out_obj.ConcurrentBuild = obj.ConcurrentBuild
	out_obj.ResumeBlocked = obj.ResumeBlocked
	return
}

func PrepareUrl(url string) string {
	base_url := GetBaseURL()
	return strings.ReplaceAll(url, base_url, "")
}
