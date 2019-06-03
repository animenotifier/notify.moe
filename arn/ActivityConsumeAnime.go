package arn

import "sort"

// ActivityConsumeAnime is a user activity that consumes anime.
type ActivityConsumeAnime struct {
	AnimeID     string `json:"animeId"`
	FromEpisode int    `json:"fromEpisode"`
	ToEpisode   int    `json:"toEpisode"`

	hasID
	hasCreator
	hasLikes
}

// NewActivityConsumeAnime creates a new activity.
func NewActivityConsumeAnime(animeID string, fromEpisode int, toEpisode int, userID UserID) *ActivityConsumeAnime {
	return &ActivityConsumeAnime{
		hasID: hasID{
			ID: GenerateID("ActivityConsumeAnime"),
		},
		hasCreator: hasCreator{
			Created:   DateTimeUTC(),
			CreatedBy: userID,
		},
		AnimeID:     animeID,
		FromEpisode: fromEpisode,
		ToEpisode:   toEpisode,
	}
}

// Anime returns the anime.
func (activity *ActivityConsumeAnime) Anime() *Anime {
	anime, _ := GetAnime(activity.AnimeID)
	return anime
}

// TypeName returns the type name.
func (activity *ActivityConsumeAnime) TypeName() string {
	return "ActivityConsumeAnime"
}

// Self returns the object itself.
func (activity *ActivityConsumeAnime) Self() Loggable {
	return activity
}

// LastActivityConsumeAnime returns the last activity for the given anime.
func (user *User) LastActivityConsumeAnime(animeID string) *ActivityConsumeAnime {
	activities := FilterActivitiesConsumeAnime(func(activity *ActivityConsumeAnime) bool {
		return activity.AnimeID == animeID && activity.CreatedBy == user.ID
	})

	if len(activities) == 0 {
		return nil
	}

	sort.Slice(activities, func(i, j int) bool {
		return activities[i].Created > activities[j].Created
	})

	return activities[0]
}

// FilterActivitiesConsumeAnime filters all anime consumption activities by a custom function.
func FilterActivitiesConsumeAnime(filter func(*ActivityConsumeAnime) bool) []*ActivityConsumeAnime {
	var filtered []*ActivityConsumeAnime

	for obj := range DB.All("ActivityConsumeAnime") {
		realObject := obj.(*ActivityConsumeAnime)

		if filter(realObject) {
			filtered = append(filtered, realObject)
		}
	}

	return filtered
}

// // OnLike is called when the activity receives a like.
// func (activity *Activity) OnLike(likedBy *User) {
// 	if likedBy.ID == activity.CreatedBy {
// 		return
// 	}

// 	go func() {
// 		notifyUser := activity.Creator()

// 		notifyUser.SendNotification(&PushNotification{
// 			Title:   likedBy.Nick + " liked your activity",
// 			Message: activity.TextByUser(notifyUser),
// 			Icon:    "https:" + likedBy.AvatarLink("large"),
// 			Link:    "https://notify.moe" + activity.Link(),
// 			Type:    NotificationTypeLike,
// 		})
// 	}()
// }
