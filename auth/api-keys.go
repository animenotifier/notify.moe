package auth

import (
	"encoding/json"
	"io/ioutil"

	"github.com/animenotifier/arn"
)

var apiKeys arn.APIKeys

func init() {
	data, _ := ioutil.ReadFile("security/api-keys.json")
	err := json.Unmarshal(data, &apiKeys)

	if err != nil {
		panic(err)
	}
}
