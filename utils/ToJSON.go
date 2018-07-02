package utils

import jsoniter "github.com/json-iterator/go"

// ToJSON converts an object to a JSON string, ignoring errors.
func ToJSON(v interface{}) string {
	str, _ := jsoniter.Marshal(v)
	return string(str)
}
