package arn

import (
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/aerogo/nano"
)

// AnimeList is a list of anime list items.
type AnimeList struct {
	UserID UserID           `json:"userId"`
	Items  []*AnimeListItem `json:"items"`

	sync.Mutex
}

// Add adds an anime to the list if it hasn't been added yet.
func (list *AnimeList) Add(animeID string) error {
	if list.Contains(animeID) {
		return errors.New("Anime " + animeID + " has already been added")
	}

	creationDate := DateTimeUTC()

	item := &AnimeListItem{
		AnimeID: animeID,
		Status:  AnimeListStatusPlanned,
		Rating:  AnimeListItemRating{},
		Created: creationDate,
		Edited:  creationDate,
	}

	if item.Anime() == nil {
		return errors.New("Invalid anime ID")
	}

	list.Lock()
	list.Items = append(list.Items, item)
	list.Unlock()

	return nil
}

// Remove removes the anime ID from the list.
func (list *AnimeList) Remove(animeID string) bool {
	list.Lock()
	defer list.Unlock()

	for index, item := range list.Items {
		if item.AnimeID == animeID {
			list.Items = append(list.Items[:index], list.Items[index+1:]...)
			return true
		}
	}

	return false
}

// Contains checks if the list contains the anime ID already.
func (list *AnimeList) Contains(animeID string) bool {
	list.Lock()
	defer list.Unlock()

	for _, item := range list.Items {
		if item.AnimeID == animeID {
			return true
		}
	}

	return false
}

// HasItemsWithStatus checks if the list contains an anime with the given status.
func (list *AnimeList) HasItemsWithStatus(status string) bool {
	list.Lock()
	defer list.Unlock()

	for _, item := range list.Items {
		if item.Status == status {
			return true
		}
	}

	return false
}

// Find returns the list item with the specified anime ID, if available.
func (list *AnimeList) Find(animeID string) *AnimeListItem {
	list.Lock()
	defer list.Unlock()

	for _, item := range list.Items {
		if item.AnimeID == animeID {
			return item
		}
	}

	return nil
}

// Import adds an anime to the list if it hasn't been added yet
// and if it did exist it will update episode, rating and notes.
func (list *AnimeList) Import(item *AnimeListItem) {
	existing := list.Find(item.AnimeID)

	// If it doesn't exist yet: Simply add it.
	if existing == nil {
		list.Lock()
		list.Items = append(list.Items, item)
		list.Unlock()
		return
	}

	// Temporary save it before changing the status
	// because status changes can modify the episode count.
	// This will prevent loss of "episodes watched" data.
	existingEpisodes := existing.Episodes

	// Status
	existing.Status = item.Status
	existing.OnStatusChange()

	// Episodes
	if item.Episodes > existingEpisodes {
		existing.Episodes = item.Episodes
	} else {
		existing.Episodes = existingEpisodes
	}

	existing.OnEpisodesChange()

	// Rating
	if existing.Rating.Overall == 0 {
		existing.Rating.Overall = item.Rating.Overall
		existing.Rating.Clamp()
	}

	if existing.Notes == "" {
		existing.Notes = item.Notes
	}

	if item.RewatchCount > existing.RewatchCount {
		existing.RewatchCount = item.RewatchCount
	}

	// Edited
	existing.Edited = DateTimeUTC()
}

// User returns the user this anime list belongs to.
func (list *AnimeList) User() *User {
	user, _ := GetUser(list.UserID)
	return user
}

// Sort ...
func (list *AnimeList) Sort() {
	list.Lock()
	defer list.Unlock()

	sort.Slice(list.Items, func(i, j int) bool {
		a := list.Items[i]
		b := list.Items[j]

		if (a.Status != AnimeListStatusWatching && a.Status != AnimeListStatusPlanned) && (b.Status != AnimeListStatusWatching && b.Status != AnimeListStatusPlanned) {
			if a.Rating.Overall == b.Rating.Overall {
				return a.Anime().Title.Canonical < b.Anime().Title.Canonical
			}

			return a.Rating.Overall > b.Rating.Overall
		}

		epsA := a.Anime().UpcomingEpisode()
		epsB := b.Anime().UpcomingEpisode()

		if epsA == nil && epsB == nil {
			if a.Rating.Overall == b.Rating.Overall {
				return a.Anime().Title.Canonical < b.Anime().Title.Canonical
			}

			return a.Rating.Overall > b.Rating.Overall
		}

		if epsA == nil {
			return false
		}

		if epsB == nil {
			return true
		}

		return epsA.Episode.AiringDate.Start < epsB.Episode.AiringDate.Start
	})
}

// SortByRating sorts the anime list by overall rating.
func (list *AnimeList) SortByRating() {
	list.Lock()
	defer list.Unlock()

	sort.Slice(list.Items, func(i, j int) bool {
		a := list.Items[i]
		b := list.Items[j]

		if a.Rating.Overall == b.Rating.Overall {
			return a.Anime().Title.Canonical < b.Anime().Title.Canonical
		}

		return a.Rating.Overall > b.Rating.Overall
	})
}

// Top returns the top entries.
func (list *AnimeList) Top(count int) []*AnimeListItem {
	list.Lock()
	defer list.Unlock()

	sort.Slice(list.Items, func(i, j int) bool {
		a := list.Items[i]
		b := list.Items[j]

		if a.Rating.Overall == b.Rating.Overall {
			return a.Anime().Title.Canonical < b.Anime().Title.Canonical
		}

		return a.Rating.Overall > b.Rating.Overall
	})

	if count > len(list.Items) {
		count = len(list.Items)
	}

	tmp := make([]*AnimeListItem, count)
	copy(tmp, list.Items[:count])
	return tmp
}

// Watching ...
func (list *AnimeList) Watching() *AnimeList {
	return list.FilterStatus(AnimeListStatusWatching)
}

// FilterStatus ...
func (list *AnimeList) FilterStatus(status string) *AnimeList {
	newList := &AnimeList{
		UserID: list.UserID,
		Items:  []*AnimeListItem{},
	}

	list.Lock()
	defer list.Unlock()

	for _, item := range list.Items {
		if item.Status == status {
			newList.Items = append(newList.Items, item)
		}
	}

	return newList
}

// WithoutPrivateItems returns a new anime list with the private items removed.
func (list *AnimeList) WithoutPrivateItems() *AnimeList {
	list.Lock()
	defer list.Unlock()

	newList := &AnimeList{
		UserID: list.UserID,
		Items:  make([]*AnimeListItem, 0, len(list.Items)),
	}

	for _, item := range list.Items {
		if !item.Private {
			newList.Items = append(newList.Items, item)
		}
	}

	return newList
}

// SplitByStatus splits the anime list into multiple ones by status.
func (list *AnimeList) SplitByStatus() map[string]*AnimeList {
	statusToList := map[string]*AnimeList{}

	statusToList[AnimeListStatusWatching] = &AnimeList{
		UserID: list.UserID,
		Items:  []*AnimeListItem{},
	}

	statusToList[AnimeListStatusCompleted] = &AnimeList{
		UserID: list.UserID,
		Items:  []*AnimeListItem{},
	}

	statusToList[AnimeListStatusPlanned] = &AnimeList{
		UserID: list.UserID,
		Items:  []*AnimeListItem{},
	}

	statusToList[AnimeListStatusHold] = &AnimeList{
		UserID: list.UserID,
		Items:  []*AnimeListItem{},
	}

	statusToList[AnimeListStatusDropped] = &AnimeList{
		UserID: list.UserID,
		Items:  []*AnimeListItem{},
	}

	list.Lock()
	defer list.Unlock()

	for _, item := range list.Items {
		statusList := statusToList[item.Status]
		statusList.Items = append(statusList.Items, item)
	}

	return statusToList
}

// NormalizeRatings normalizes all ratings so that they are perfectly stretched among the full scale.
func (list *AnimeList) NormalizeRatings() {
	list.Lock()
	defer list.Unlock()

	mapped := map[float64]float64{}
	all := make([]float64, 0, len(list.Items))

	for _, item := range list.Items {
		// Zero rating counts as not rated
		if item.Rating.Overall == 0 {
			continue
		}

		_, found := mapped[item.Rating.Overall]

		if !found {
			mapped[item.Rating.Overall] = item.Rating.Overall
			all = append(all, item.Rating.Overall)
		}
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i] < all[j]
	})

	count := len(all)

	// Prevent division by zero
	if count <= 1 {
		return
	}

	step := 9.9 / float64(count-1)
	currentRating := 0.1

	for _, rating := range all {
		mapped[rating] = currentRating
		currentRating += step
	}

	for _, item := range list.Items {
		item.Rating.Overall = mapped[item.Rating.Overall]
		item.Rating.Clamp()
	}
}

// GetID returns the anime ID.
func (list *AnimeList) GetID() string {
	return list.UserID
}

// TypeName returns the type name.
func (list *AnimeList) TypeName() string {
	return "AnimeList"
}

// Self returns the object itself.
func (list *AnimeList) Self() Loggable {
	return list
}

// Genres returns a map of genre names mapped to the list items that belong to that genre.
func (list *AnimeList) Genres() map[string][]*AnimeListItem {
	genreToListItems := map[string][]*AnimeListItem{}

	for _, item := range list.Items {
		for _, genre := range item.Anime().Genres {
			genreToListItems[genre] = append(genreToListItems[genre], item)
		}
	}

	return genreToListItems
}

// TopGenres returns the most liked genres for the user's anime list.
func (list *AnimeList) TopGenres(count int) []string {
	genreItems := list.Genres()
	genreAffinity := map[string]float64{}
	bestGenres := make([]string, 0, len(genreItems))

	for genre, animeListItems := range genreItems {
		if genre == "Action" || genre == "Comedy" {
			continue
		}

		affinity := 0.0

		for _, item := range animeListItems {
			if item.Status != AnimeListStatusCompleted {
				continue
			}

			if item.Rating.Overall != 0 {
				affinity += item.Rating.Overall - AverageRating
			} else {
				// Add 0.1 to avoid all affinities being 0 when a user doesn't have any rated anime.
				affinity += 0.1
			}
		}

		genreAffinity[genre] = affinity
		bestGenres = append(bestGenres, genre)
	}

	sort.Slice(bestGenres, func(i, j int) bool {
		aAffinity := genreAffinity[bestGenres[i]]
		bAffinity := genreAffinity[bestGenres[j]]

		if aAffinity == bAffinity {
			return bestGenres[i] < bestGenres[j]
		}

		return aAffinity > bAffinity
	})

	if len(bestGenres) > count {
		bestGenres = bestGenres[:count]
	}

	return bestGenres
}

// RemoveDuplicates removes duplicate entries.
func (list *AnimeList) RemoveDuplicates() {
	list.Lock()
	defer list.Unlock()

	existed := map[string]bool{}
	newItems := make([]*AnimeListItem, 0, len(list.Items))

	for _, item := range list.Items {
		_, exists := existed[item.AnimeID]

		if exists {
			fmt.Println(list.User().Nick, "removed anime list item duplicate", item.AnimeID)
			continue
		}

		newItems = append(newItems, item)
		existed[item.AnimeID] = true
	}

	list.Items = newItems
}

// StreamAnimeLists returns a stream of all anime.
func StreamAnimeLists() <-chan *AnimeList {
	channel := make(chan *AnimeList, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("AnimeList") {
			channel <- obj.(*AnimeList)
		}

		close(channel)
	}()

	return channel
}

// AllAnimeLists returns a slice of all anime.
func AllAnimeLists() ([]*AnimeList, error) {
	all := make([]*AnimeList, 0, DB.Collection("AnimeList").Count())

	stream := StreamAnimeLists()

	for obj := range stream {
		all = append(all, obj)
	}

	return all, nil
}

// GetAnimeList ...
func GetAnimeList(userID UserID) (*AnimeList, error) {
	animeList, err := DB.Get("AnimeList", userID)

	if err != nil {
		return nil, err
	}

	return animeList.(*AnimeList), nil
}
