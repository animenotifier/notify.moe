package arn

import (
	"sort"
	"sync"

	"github.com/aerogo/nano"
)

// AnimeRelations is a list of relations for an anime.
type AnimeRelations struct {
	AnimeID string           `json:"animeId" mainID:"true"`
	Items   []*AnimeRelation `json:"items" editable:"true"`

	sync.Mutex
}

// Link returns the link for that object.
func (relations *AnimeRelations) Link() string {
	return "/anime/" + relations.AnimeID + "/relations"
}

// SortByStartDate ...
func (relations *AnimeRelations) SortByStartDate() {
	relations.Lock()
	defer relations.Unlock()

	sort.Slice(relations.Items, func(i, j int) bool {
		a := relations.Items[i].Anime()
		b := relations.Items[j].Anime()

		if a == nil {
			return false
		}

		if b == nil {
			return true
		}

		if a.StartDate == b.StartDate {
			return a.Title.Canonical < b.Title.Canonical
		}

		return a.StartDate < b.StartDate
	})
}

// Anime returns the anime the relations list refers to.
func (relations *AnimeRelations) Anime() *Anime {
	anime, _ := GetAnime(relations.AnimeID)
	return anime
}

// String implements the default string serialization.
func (relations *AnimeRelations) String() string {
	return relations.Anime().String()
}

// GetID returns the anime ID.
func (relations *AnimeRelations) GetID() string {
	return relations.AnimeID
}

// TypeName returns the type name.
func (relations *AnimeRelations) TypeName() string {
	return "AnimeRelations"
}

// Self returns the object itself.
func (relations *AnimeRelations) Self() Loggable {
	return relations
}

// Find returns the relation with the specified anime ID, if available.
func (relations *AnimeRelations) Find(animeID string) *AnimeRelation {
	relations.Lock()
	defer relations.Unlock()

	for _, item := range relations.Items {
		if item.AnimeID == animeID {
			return item
		}
	}

	return nil
}

// Remove removes the anime ID from the relations.
func (relations *AnimeRelations) Remove(animeID string) bool {
	relations.Lock()
	defer relations.Unlock()

	for index, item := range relations.Items {
		if item.AnimeID == animeID {
			relations.Items = append(relations.Items[:index], relations.Items[index+1:]...)
			return true
		}
	}

	return false
}

// GetAnimeRelations ...
func GetAnimeRelations(animeID string) (*AnimeRelations, error) {
	obj, err := DB.Get("AnimeRelations", animeID)

	if err != nil {
		return nil, err
	}

	return obj.(*AnimeRelations), nil
}

// StreamAnimeRelations returns a stream of all anime relations.
func StreamAnimeRelations() <-chan *AnimeRelations {
	channel := make(chan *AnimeRelations, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("AnimeRelations") {
			channel <- obj.(*AnimeRelations)
		}

		close(channel)
	}()

	return channel
}
