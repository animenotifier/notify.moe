package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/animenotifier/arn"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/animenotifier/avatar/lib"
)

var port = "8001"

func init() {
	flag.StringVar(&port, "port", "", "Port the HTTP server should listen on")
	flag.Parse()
}

func main() {
	// Switch to main directory
	exe, err := os.Executable()

	if err != nil {
		panic(err)
	}

	root := path.Dir(exe)
	os.Chdir(path.Join(root, "../../"))

	// Start server
	http.HandleFunc("/", onRequest)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// onRequest handles requests and refreshes the requested avatar.
func onRequest(w http.ResponseWriter, req *http.Request) {
	// User ID is simply the path without the slash
	userID := strings.TrimPrefix(req.URL.Path, "/")

	// Get user from database
	user, err := arn.GetUser(userID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	// Refresh
	lib.RefreshAvatar(user)

	// Send JSON response
	buffer, err := json.Marshal(user.Avatar)

	if err != nil {
		io.WriteString(w, err.Error())
	}

	w.Write(buffer)
}
