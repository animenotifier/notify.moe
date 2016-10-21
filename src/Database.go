package main

import as "github.com/aerospike/aerospike-client-go"

// Client ...
var client *as.Client

// InitDatabase ...
func InitDatabase() {
	client, _ = as.NewClient("127.0.0.1", 3000)
}

// GetAnime ...
func GetAnime(id int) *Anime {
	key, _ := as.NewKey("arn", "Anime", id)
	anime := new(Anime)
	client.GetObject(nil, key, anime)
	return anime
}
