package jengo_src

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Class       string     `yaml:"_class"`
	AbsoluteURL string     `yaml:"absoluteUrl"`
	Description string     `yaml:"description"`
	FullName    string     `yaml:"fullName"`
	ID          string     `yaml:"id"`
	Property    []Property `yaml:"property"`
}
type Property struct {
	Class             string        `yaml:"_class"`
	Address           string        `yaml:"address,omitempty"`
	Triggers          []interface{} `yaml:"triggers,omitempty"`
	InsensitiveSearch bool          `yaml:"insensitiveSearch,omitempty"`
}

func GetUser(username string) (obj User){
	responseData, err := HandleRequest("GET", Kwargs{"name": "user", "user_name": username})
	if err != nil {
		ErrorLog(err)
	}
	fmt.Println(string(responseData))

	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		ErrorLog(err)
	}
	fmt.Println(obj)

	return
}
