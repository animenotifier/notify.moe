package utils

import "encoding/json"

// ToJSON converts an object to a JSON string, ignoring errors.
func ToJSON(v interface{}) string {
	str, _ := json.Marshal(v)
	return string(str)
}
