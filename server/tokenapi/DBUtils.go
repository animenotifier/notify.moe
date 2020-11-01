package tokenapi

import (
	"github.com/akyoto/uuid"
	"github.com/animenotifier/notify.moe/arn"
)

func GetUserFromToken(token uuid.UUID) *arn.User {
	user := &arn.User{}

	for user = range arn.StreamUsers() {
		if user.APIToken == token {
			break
		}
	}

	return user
}
