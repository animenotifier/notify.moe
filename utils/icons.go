package utils

import (
	"io/ioutil"
	"strings"
)

var svgIcons = make(map[string]string)

func init() {
	files, _ := ioutil.ReadDir("images/icons/svg/")

	for _, file := range files {
		name := strings.Replace(file.Name(), ".svg", "", 1)
		data, _ := ioutil.ReadFile("images/icons/svg/" + name + ".svg")
		svgIcons[name] = strings.Replace(string(data), "<svg ", "<svg class='icon' ", 1)
	}
}

// Icon ...
func Icon(name string) string {
	return svgIcons[name]
}

// RawIcon ...
func RawIcon(name string) string {
	return strings.Replace(svgIcons[name], "class='icon'", "class='raw-icon'", 1)
}
