package main

import (
	"flag"
	"log"
	"net/http"
)

func refreshAvatar(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("ok"))
}

var port = "8001"

func init() {
	flag.StringVar(&port, "port", "", "Port the HTTP server should listen on")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", refreshAvatar)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
