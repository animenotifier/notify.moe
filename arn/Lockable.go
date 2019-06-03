package arn

import (
	"errors"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Lockable ...
type Lockable interface {
	Lock(userID UserID)
	Unlock(userID UserID)
	IsLocked() bool
	Save()
}

// LockEventReceiver ...
type LockEventReceiver interface {
	OnLock(user *User)
	OnUnlock(user *User)
}

// LockAction ...
func LockAction() *api.Action {
	return &api.Action{
		Name:  "lock",
		Route: "/lock",
		Run: func(obj interface{}, ctx aero.Context) error {
			lockable := obj.(Lockable)
			user := GetUserFromContext(ctx)

			if user == nil {
				return errors.New("Not logged in")
			}

			lockable.Lock(user.ID)

			// Call OnLock if the object implements it
			receiver, ok := lockable.(LockEventReceiver)

			if ok {
				receiver.OnLock(user)
			}

			lockable.Save()
			return nil
		},
	}
}

// UnlockAction ...
func UnlockAction() *api.Action {
	return &api.Action{
		Name:  "unlock",
		Route: "/unlock",
		Run: func(obj interface{}, ctx aero.Context) error {
			lockable := obj.(Lockable)
			user := GetUserFromContext(ctx)

			if user == nil {
				return errors.New("Not logged in")
			}

			lockable.Unlock(user.ID)

			// Call OnUnlock if the object implements it
			receiver, ok := lockable.(LockEventReceiver)

			if ok {
				receiver.OnUnlock(user)
			}

			lockable.Save()
			return nil
		},
	}
}

// IsLocked returns true if the given object is locked.
func IsLocked(obj interface{}) bool {
	lockable, isLockable := obj.(Lockable)
	return isLockable && lockable.IsLocked()
}
