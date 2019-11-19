package arn

import "github.com/aerogo/api"

// Register data lists.
func init() {
	DataLists["genders"] = []*Option{
		// &Option{"", "Prefer not to say"},
		{"male", "Male"},
		{"female", "Female"},
	}

	// Actions
	API.RegisterActions("User", []*api.Action{
		// Add follow
		FollowAction(),

		// Remove follow
		UnfollowAction(),
	})
}
