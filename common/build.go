package common

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type BuildL struct {
	ID                string `yaml:"id"`
	Number            int    `yaml:"number"`
	Timestamp         int64  `yaml:"timestamp"`
	Result            string `yaml:"result"`
	URL               string `yaml:"url"`
	EstimatedDuration int    `yaml:"estimatedDuration"`
	Duration          int    `yaml:"duration"`
	Description       string `yaml:"description"`
}

type BuildResponse struct {
	ID                string      `yaml:"id"`
	Number            int         `yaml:"number"`
	Timestamp         int64       `yaml:"timestamp"`
	Result            string      `yaml:"result"`
	URL               string      `yaml:"url"`
	EstimatedDuration int         `yaml:"estimatedDuration"`
	Duration          int         `yaml:"duration"`
	Description       interface{} `yaml:"description"`
	FullDisplayName   string      `yaml:"fullDisplayName"`
	Building          bool        `yaml:"building"`
	DisplayName       string      `yaml:"displayName"`
	Executor          interface{} `yaml:"executor"`
	KeepLog           bool        `yaml:"keepLog"`
	QueueID           int         `yaml:"queueId"`
}

type BuildListResponse struct {
	Builds []BuildL `yaml:"builds"`
}

func ListBuilds(job_name string) (res [][]string) {
	responseData, err := HandleRequest("GET", Kwargs{"name": "builds", "job_name": job_name})

	if err != nil {
		log.Fatal(err)
	}

	var obj BuildListResponse
	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		log.Fatal(err)
	}
	for _, build := range obj.Builds {
		// Divide by 1000 for seconds in epoch
		unixTime := time.Unix(build.Timestamp/1000, 0).Format("2006-02-01 15:04:05")

		res = append(res,
			[]string{build.ID,
				strconv.Itoa(build.Number),
				unixTime,
				build.Result,
				build.URL,
				strconv.Itoa(build.Duration),
				strconv.Itoa(build.EstimatedDuration),
				build.Description})
	}
	return
}

func GetBuild(job_name string, build_name string) (obj BuildResponse) {
	responseData, err := HandleRequest("GET", Kwargs{"name": "build_info", "job_name": job_name, "build_name": build_name})
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		log.Fatal(err)
	}
	return
}
