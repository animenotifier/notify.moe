package arn

import "github.com/animenotifier/japanese/client"

// JapaneseTokenizer tokenizes a sentence via the HTTP API.
var JapaneseTokenizer = &client.Tokenizer{
	Endpoint: "http://127.0.0.1:6000/",
}
