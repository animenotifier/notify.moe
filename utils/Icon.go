package utils

import (
	"io/ioutil"
	"strings"
)

var svgIcons = make(map[string]string)

func init() {
	files, _ := ioutil.ReadDir("images/icons/")

	for _, file := range files {
		name := strings.TrimSuffix(file.Name(), ".svg")
		data, _ := ioutil.ReadFile("images/icons/" + name + ".svg")
		svgIcons[name] = strings.Replace(string(data), "<svg ", "<svg class='icon icon-"+name+"' ", 1)
	}
}

// Icon ...
func Icon(name string) string {
	return svgIcons[name]
}

// RawIcon ...
func RawIcon(name string) string {
	return strings.Replace(svgIcons[name], "class='icon", "class='raw-icon", 1)
}
