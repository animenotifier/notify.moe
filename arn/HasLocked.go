package arn

// HasLocked implements common like and unlike methods.
type hasLocked struct {
	Locked bool `json:"locked"`
}

// Lock locks the object.
func (obj *hasLocked) Lock(userID UserID) {
	obj.Locked = true
}

// Unlock unlocks the object.
func (obj *hasLocked) Unlock(userID UserID) {
	obj.Locked = false
}

// IsLocked implements the Lockable interface.
func (obj *hasLocked) IsLocked() bool {
	return obj.Locked
}
