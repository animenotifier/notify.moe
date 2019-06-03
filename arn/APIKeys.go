package arn

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/animenotifier/anilist"
	"github.com/animenotifier/osu"
	jsoniter "github.com/json-iterator/go"
)

// Root is the full path to the root directory of notify.moe repository.
var Root = os.Getenv("ARN_ROOT")

// APIKeys are global API keys for several services
var APIKeys APIKeysData

// APIKeysData ...
type APIKeysData struct {
	Google struct {
		ID     string `json:"id"`
		Secret string `json:"secret"`
	} `json:"google"`

	Facebook struct {
		ID     string `json:"id"`
		Secret string `json:"secret"`
	} `json:"facebook"`

	Twitter struct {
		ID     string `json:"id"`
		Secret string `json:"secret"`
	} `json:"twitter"`

	Discord struct {
		ID     string `json:"id"`
		Secret string `json:"secret"`
		Token  string `json:"token"`
	} `json:"discord"`

	SoundCloud struct {
		ID     string `json:"id"`
		Secret string `json:"secret"`
	} `json:"soundcloud"`

	GoogleAPI struct {
		Key string `json:"key"`
	} `json:"googleAPI"`

	IPInfoDB struct {
		ID string `json:"id"`
	} `json:"ipInfoDB"`

	AniList struct {
		ID     string `json:"id"`
		Secret string `json:"secret"`
	} `json:"anilist"`

	Osu struct {
		Secret string `json:"secret"`
	} `json:"osu"`

	PayPal struct {
		ID     string `json:"id"`
		Secret string `json:"secret"`
	} `json:"paypal"`

	VAPID struct {
		Subject    string `json:"subject"`
		PublicKey  string `json:"publicKey"`
		PrivateKey string `json:"privateKey"`
	} `json:"vapid"`

	SMTP struct {
		Server   string `json:"server"`
		Address  string `json:"address"`
		Password string `json:"password"`
	} `json:"smtp"`

	S3 struct {
		ID     string `json:"id"`
		Secret string `json:"secret"`
	} `json:"s3"`
}

func init() {
	// Path for API keys
	apiKeysPath := path.Join(Root, "security/api-keys.json")

	// If the API keys file is not available, create a symlink to the default API keys
	if _, err := os.Stat(apiKeysPath); os.IsNotExist(err) {
		defaultAPIKeysPath := path.Join(Root, "security/default/api-keys.json")
		err := os.Link(defaultAPIKeysPath, apiKeysPath)

		if err != nil {
			panic(err)
		}
	}

	// Load API keys
	data, err := ioutil.ReadFile(apiKeysPath)

	if err != nil {
		panic(err)
	}

	// Parse JSON
	err = jsoniter.Unmarshal(data, &APIKeys)

	if err != nil {
		panic(err)
	}

	// Set Osu API key
	osu.APIKey = APIKeys.Osu.Secret

	// Set Anilist API keys
	anilist.APIKeyID = APIKeys.AniList.ID
	anilist.APIKeySecret = APIKeys.AniList.Secret
}
