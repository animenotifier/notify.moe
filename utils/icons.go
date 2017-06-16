package utils

import (
	"io/ioutil"
	"strings"
)

// Icon ...
func Icon(name string) string {
	data, _ := ioutil.ReadFile("images/icons/svg/" + name + ".svg")
	return strings.Replace(string(data), "<svg ", "<svg class='icon' ", 1)
}

// RawIcon ...
func RawIcon(name string) string {
	data, _ := ioutil.ReadFile("images/icons/svg/" + name + ".svg")
	return strings.Replace(string(data), "<svg ", "<svg class='raw-icon' ", 1)
}
