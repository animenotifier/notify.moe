package arn

// GroupMember ...
type GroupMember struct {
	UserID UserID `json:"userId"`
	Role   string `json:"role"`
	Joined string `json:"joined"`

	user *User
}

// User returns the user.
func (member *GroupMember) User() *User {
	if member.user != nil {
		return member.user
	}

	member.user, _ = GetUser(member.UserID)
	return member.user
}
