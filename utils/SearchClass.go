package utils

import "reflect"

// SearchClass returns the class for a search section
// depending on the length of the search results.
func SearchClass(results interface{}) string {
	if reflect.ValueOf(results).Len() == 0 {
		return "search-section-disabled"
	}

	return ""
}
