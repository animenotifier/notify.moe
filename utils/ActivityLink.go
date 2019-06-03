package utils

import (
	"fmt"
	"strings"

	"github.com/animenotifier/notify.moe/arn"
)

// ActivityAPILink returns the API link for any activity.
func ActivityAPILink(activity arn.Activity) string {
	id := activity.GetID()
	typeName := activity.TypeName()

	if activity.TypeName() == "ActivityCreate" {
		created := activity.(*arn.ActivityCreate)
		typeName = created.ObjectType
		id = created.ObjectID
	}

	return fmt.Sprintf("/api/%s/%s", strings.ToLower(typeName), id)
}
