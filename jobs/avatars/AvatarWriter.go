package main

// AvatarOutput represents a system that saves an avatar locally (in database or as a file, e.g.)
type AvatarOutput interface {
	SaveAvatar(*Avatar) error
}
