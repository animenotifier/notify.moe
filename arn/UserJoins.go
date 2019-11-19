package arn

// Threads ...
func (user *User) Threads() []*Thread {
	threads := GetThreadsByUser(user)
	return threads
}

// // Posts ...
// func (user *User) Posts() []*Post {
// 	posts, _ := GetPostsByUser(user)
// 	return posts
// }

// Settings ...
func (user *User) Settings() *Settings {
	settings, _ := GetSettings(user.ID)
	return settings
}

// Analytics ...
func (user *User) Analytics() *Analytics {
	analytics, _ := GetAnalytics(user.ID)
	return analytics
}

// AnimeList ...
func (user *User) AnimeList() *AnimeList {
	animeList, _ := GetAnimeList(user.ID)
	return animeList
}

// PushSubscriptions ...
func (user *User) PushSubscriptions() *PushSubscriptions {
	subs, _ := GetPushSubscriptions(user.ID)
	return subs
}

// Inventory ...
func (user *User) Inventory() *Inventory {
	inventory, _ := GetInventory(user.ID)
	return inventory
}

// Notifications returns the list of user notifications.
func (user *User) Notifications() *UserNotifications {
	notifications, _ := GetUserNotifications(user.ID)
	return notifications
}

// DraftIndex ...
func (user *User) DraftIndex() *DraftIndex {
	draftIndex, _ := GetDraftIndex(user.ID)
	return draftIndex
}

// SoundTracks returns the soundtracks posted by the user.
func (user *User) SoundTracks() []*SoundTrack {
	tracks := FilterSoundTracks(func(track *SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0 && track.CreatedBy == user.ID
	})
	return tracks
}
