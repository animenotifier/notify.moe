package arn

import "fmt"

// Name is the combination of a first and last name.
type Name struct {
	First string `json:"first" editable:"true"`
	Last  string `json:"last" editable:"true"`
}

// String returns the default visualization of the name.
func (name Name) String() string {
	return fmt.Sprintf("%s %s", name.First, name.Last)
}
