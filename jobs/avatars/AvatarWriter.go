package main

// AvatarWriter represents a system that saves an avatar locally (in database or as a file, e.g.)
type AvatarWriter interface {
	SaveAvatar(*Avatar) error
}
