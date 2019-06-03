package arn

import (
	"errors"
	"sync"

	"github.com/aerogo/nano"
	"github.com/akyoto/color"
)

// AnimeCharacters is a list of characters for an anime.
type AnimeCharacters struct {
	AnimeID string            `json:"animeId" mainID:"true"`
	Items   []*AnimeCharacter `json:"items" editable:"true"`

	sync.Mutex
}

// Anime returns the anime the characters refer to.
func (characters *AnimeCharacters) Anime() *Anime {
	anime, _ := GetAnime(characters.AnimeID)
	return anime
}

// Add adds an anime character to the list.
func (characters *AnimeCharacters) Add(animeCharacter *AnimeCharacter) error {
	if animeCharacter.CharacterID == "" || animeCharacter.Role == "" {
		return errors.New("Empty ID or role")
	}

	characters.Lock()
	characters.Items = append(characters.Items, animeCharacter)
	characters.Unlock()

	return nil
}

// FindByMapping finds an anime character by the given mapping.
func (characters *AnimeCharacters) FindByMapping(service string, serviceID string) *AnimeCharacter {
	characters.Lock()
	defer characters.Unlock()

	for _, animeCharacter := range characters.Items {
		character := animeCharacter.Character()

		if character == nil {
			color.Red("Anime %s has an incorrect Character ID inside AnimeCharacter: %s", characters.AnimeID, animeCharacter.CharacterID)
			continue
		}

		if character.GetMapping(service) == serviceID {
			return animeCharacter
		}
	}

	return nil
}

// Link returns the link for that object.
func (characters *AnimeCharacters) Link() string {
	return "/anime/" + characters.AnimeID + "/characters"
}

// String implements the default string serialization.
func (characters *AnimeCharacters) String() string {
	return characters.Anime().String()
}

// GetID returns the anime ID.
func (characters *AnimeCharacters) GetID() string {
	return characters.AnimeID
}

// TypeName returns the type name.
func (characters *AnimeCharacters) TypeName() string {
	return "AnimeCharacters"
}

// Self returns the object itself.
func (characters *AnimeCharacters) Self() Loggable {
	return characters
}

// Contains tells you whether the given character ID exists.
func (characters *AnimeCharacters) Contains(characterID string) bool {
	characters.Lock()
	defer characters.Unlock()

	for _, item := range characters.Items {
		if item.CharacterID == characterID {
			return true
		}
	}

	return false
}

// First gives you the first "count" anime characters.
func (characters *AnimeCharacters) First(count int) []*AnimeCharacter {
	characters.Lock()
	defer characters.Unlock()

	if count > len(characters.Items) {
		count = len(characters.Items)
	}

	return characters.Items[:count]
}

// GetAnimeCharacters ...
func GetAnimeCharacters(animeID string) (*AnimeCharacters, error) {
	obj, err := DB.Get("AnimeCharacters", animeID)

	if err != nil {
		return nil, err
	}

	return obj.(*AnimeCharacters), nil
}

// StreamAnimeCharacters returns a stream of all anime characters.
func StreamAnimeCharacters() <-chan *AnimeCharacters {
	channel := make(chan *AnimeCharacters, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("AnimeCharacters") {
			channel <- obj.(*AnimeCharacters)
		}

		close(channel)
	}()

	return channel
}
