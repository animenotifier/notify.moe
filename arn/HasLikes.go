package arn

// HasLikes implements common like and unlike methods.
type hasLikes struct {
	Likes []string `json:"likes"`
}

// Like makes the given user ID like the object.
func (obj *hasLikes) Like(userID UserID) {
	for _, id := range obj.Likes {
		if id == userID {
			return
		}
	}

	obj.Likes = append(obj.Likes, userID)
}

// Unlike makes the given user ID unlike the object.
func (obj *hasLikes) Unlike(userID UserID) {
	for index, id := range obj.Likes {
		if id == userID {
			obj.Likes = append(obj.Likes[:index], obj.Likes[index+1:]...)
			return
		}
	}
}

// LikedBy checks to see if the user has liked the object.
func (obj *hasLikes) LikedBy(userID UserID) bool {
	for _, id := range obj.Likes {
		if id == userID {
			return true
		}
	}

	return false
}

// CountLikes returns the number of likes the object has received.
func (obj *hasLikes) CountLikes() int {
	return len(obj.Likes)
}
