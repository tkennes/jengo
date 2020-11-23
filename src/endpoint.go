package jengo_src

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Kwargs map[string]interface{}

func GetEndpoint(Kwargs map[string]interface{}) string {
	switch Kwargs["name"] {

	// Jobs`
	case "jobs":
		return "/api/json?tree=jobs[name,color,buildable,url,description]"
	case "job":
		endpoint := fmt.Sprintf("/job/%s/api/json?pretty=true", Kwargs["job_name"])
		return endpoint

	// Builds
	case "builds":
		endpoint := fmt.Sprintf("/job/%s/api/json?tree=builds[number,status,timestamp,id,result,estimatedDuration,duration,executor,description,url]", Kwargs["job_name"])
		return endpoint
	case "build":
		endpoint := fmt.Sprintf("/job/%s/%s/api/json?pretty=true", Kwargs["job_name"], Kwargs["build_name"])
		return endpoint

	// Nodes
	case "nodes":
		return "/computer/api/json?pretty=true"
	case "node":
		endpoint := fmt.Sprintf("/computer/(%s)/api/json?pretty=true", Kwargs["node_name"])
		return endpoint

	// Plugins
	case "plugins":
		return "/pluginManager/api/json?depth=1"

	// Credentials
	case "credentials":
		return "/credentials/store/system/domain/_/api/json?pretty=true"

	// Update Center
	case "update-center":
		return "/updateCenter/site/default/api/json?pretty=true&depth=1"

	// User
	case "user":
		endpoint := fmt.Sprintf("/securityRealm/user/%s/api/json", Kwargs["user_name"])
		return endpoint

	// Default: Endpoint not found
	default:
		panic(errors.New("Not possible endpoint"))
	}
}

func HandleRequest(METHOD string, Kwargs map[string]interface{}) ([]byte, error) {
	req := CreateRequest(METHOD, GetEndpoint(Kwargs))
	// Send req using http Client
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		ErrorLog(err)
	}
	if response.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Erorr: %v", response.StatusCode))
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ErrorLog(err)
	}
	return responseData, nil

}
