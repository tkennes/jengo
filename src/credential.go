package jengo_src

import (
	"encoding/json"
)

type CredentialResponse struct {
	Credentials     []interface{} `yaml:"credentials"`
	Description     string        `yaml:"description"`
	DisplayName     string        `yaml:"displayName"`
	FullDisplayName string        `yaml:"fullDisplayName"`
	FullName        string        `yaml:"fullName"`
	Global          bool          `yaml:"global"`
	URLName         string        `yaml:"urlName"`
}

func ListCredentials() (obj CredentialResponse) {
	responseData, err := HandleRequest("GET", Kwargs{"name": "credentials"})
	if err != nil {
		ErrorLog(err)
	}

	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		ErrorLog(err)
	}

	return
}
