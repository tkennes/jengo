package common

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	jclient "github.com/tomkennes/jengo/client"
)

type kwargs map[string]interface{}

func get_endpoint(kwargs map[string]interface{}) string {
	if kwargs["name"] == "jobs" {
		return "/api/json?tree=jobs[name,color,buildable,url,description]"
	} else if kwargs["name"] == "job_info" {
		endpoint := fmt.Sprintf("/job/%s/api/json?pretty=true", kwargs["job_name"])
		return endpoint
	} else if kwargs["name"] == "builds" {
		endpoint := fmt.Sprintf("/job/%s/api/json?tree=builds[number,status,timestamp,id,result,estimatedDuration,duration,executor,description,url]", kwargs["job_name"])
		return endpoint
	} else if kwargs["name"] == "build" {
		endpoint := fmt.Sprintf("/job/%s/%s/api/json?pretty=true", kwargs["job_name"], kwargs["build_name"])
		return endpoint
	} else {
		panic(errors.New("Not possible endpoint"))
	}
}

func HandleRequest(METHOD string, kwargs map[string]interface{}) ([]byte, error) {
	req := jclient.CreateRequest(METHOD, get_endpoint(kwargs))
	// Send req using http Client
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Erorr: %v", response.StatusCode))
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData, nil

}
