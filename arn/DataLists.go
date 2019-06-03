package arn

// Option is a selection list item.
type Option struct {
	Value string
	Label string
}

// DataLists maps an ID to a list of keys and values.
// Used for selection lists in UIs.
var DataLists = map[string][]*Option{}
