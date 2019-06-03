package arn

import (
	"sort"

	"github.com/aerogo/markdown"
	"github.com/aerogo/nano"
)

// Thread is a forum thread.
type Thread struct {
	Title  string   `json:"title" editable:"true"`
	Sticky int      `json:"sticky" editable:"true"`
	Tags   []string `json:"tags" editable:"true"`
	Edited string   `json:"edited"`

	hasID
	hasText
	hasPosts
	hasCreator
	hasLikes
	hasLocked

	html string
}

// Link returns the relative URL of the thread.
func (thread *Thread) Link() string {
	return "/thread/" + thread.ID
}

// HTML returns the HTML representation of the thread.
func (thread *Thread) HTML() string {
	if thread.html != "" {
		return thread.html
	}

	thread.html = markdown.Render(thread.Text)
	return thread.html
}

// String implements the default string serialization.
func (thread *Thread) String() string {
	return thread.Title
}

// Parent always returns nil for threads.
func (thread *Thread) Parent() PostParent {
	return nil
}

// GetParentID always returns an empty string for threads.
func (thread *Thread) GetParentID() string {
	return ""
}

// TypeName returns the type name.
func (thread *Thread) TypeName() string {
	return "Thread"
}

// Self returns the object itself.
func (thread *Thread) Self() Loggable {
	return thread
}

// OnLike is called when the thread receives a like.
func (thread *Thread) OnLike(likedBy *User) {
	if !thread.Creator().Settings().Notification.ForumLikes {
		return
	}

	go func() {
		thread.Creator().SendNotification(&PushNotification{
			Title:   likedBy.Nick + " liked your thread",
			Message: likedBy.Nick + " liked your thread \"" + thread.Title + "\".",
			Icon:    "https:" + likedBy.AvatarLink("large"),
			Link:    "https://notify.moe" + likedBy.Link(),
			Type:    NotificationTypeLike,
		})
	}()
}

// OnLock is called when the thread is locked.
func (thread *Thread) OnLock(user *User) {
	logEntry := NewEditLogEntry(user.ID, "edit", "Thread", thread.ID, "Locked", "false", "true")
	logEntry.Save()
}

// OnUnlock is called when the thread is unlocked.
func (thread *Thread) OnUnlock(user *User) {
	logEntry := NewEditLogEntry(user.ID, "edit", "Thread", thread.ID, "Locked", "true", "false")
	logEntry.Save()
}

// TitleByUser returns the title of the thread,
// regardless of the user language settings
// because threads are bound to one language.
func (thread *Thread) TitleByUser(user *User) string {
	return thread.Title
}

// GetThread ...
func GetThread(id string) (*Thread, error) {
	obj, err := DB.Get("Thread", id)

	if err != nil {
		return nil, err
	}

	return obj.(*Thread), nil
}

// GetThreadsByTag ...
func GetThreadsByTag(tag string) []*Thread {
	var threads []*Thread
	allTags := (tag == "" || tag == "<nil>")

	for thread := range StreamThreads() {
		if (allTags && !Contains(thread.Tags, "update")) || Contains(thread.Tags, tag) {
			threads = append(threads, thread)
		}
	}

	return threads
}

// GetThreadsByUser ...
func GetThreadsByUser(user *User) []*Thread {
	var threads []*Thread

	for thread := range StreamThreads() {
		if thread.CreatedBy == user.ID {
			threads = append(threads, thread)
		}
	}

	return threads
}

// StreamThreads ...
func StreamThreads() <-chan *Thread {
	channel := make(chan *Thread, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("Thread") {
			channel <- obj.(*Thread)
		}

		close(channel)
	}()

	return channel
}

// AllThreads ...
func AllThreads() []*Thread {
	all := make([]*Thread, 0, DB.Collection("Thread").Count())

	for obj := range StreamThreads() {
		all = append(all, obj)
	}

	return all
}

// SortThreads sorts a slice of threads for the forum view (stickies first).
func SortThreads(threads []*Thread) {
	sort.Slice(threads, func(i, j int) bool {
		a := threads[i]
		b := threads[j]

		if a.Sticky != b.Sticky {
			return a.Sticky > b.Sticky
		}

		return a.Created > b.Created
	})
}

// SortThreadsLatestFirst sorts a slice of threads by creation date.
func SortThreadsLatestFirst(threads []*Thread) {
	sort.Slice(threads, func(i, j int) bool {
		return threads[i].Created > threads[j].Created
	})
}
