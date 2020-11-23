package jengo_src

import (
	"net/http"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

func GetClient() Context {
	home, _ := homedir.Dir()
	jenkins := GetCurrentContext(filepath.Join(home, ConfigFile))

	return jenkins
}

func CreateRequest(METHOD string, URL string) *http.Request {
	jenkins := GetClient()
	req, err := http.NewRequest(METHOD, jenkins.URL+URL, nil)
	req.SetBasicAuth(jenkins.Username, jenkins.Token)
	if err != nil {
		panic(err)
	}
	return req
}
