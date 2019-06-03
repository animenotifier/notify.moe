package arn

import (
	"errors"
	"fmt"
	"sort"

	"github.com/aerogo/nano"
)

// Character represents an anime or manga character.
type Character struct {
	Name        CharacterName         `json:"name" editable:"true"`
	Image       CharacterImage        `json:"image"`
	MainQuoteID string                `json:"mainQuoteId" editable:"true"`
	Description string                `json:"description" editable:"true" type:"textarea"`
	Spoilers    []Spoiler             `json:"spoilers" editable:"true"`
	Attributes  []*CharacterAttribute `json:"attributes" editable:"true"`

	hasID
	hasPosts
	hasMappings
	hasCreator
	hasEditor
	hasLikes
	hasDraft
}

// NewCharacter creates a new character.
func NewCharacter() *Character {
	return &Character{
		hasID: hasID{
			ID: GenerateID("Character"),
		},
		hasCreator: hasCreator{
			Created: DateTimeUTC(),
		},
	}
}

// Link ...
func (character *Character) Link() string {
	return "/character/" + character.ID
}

// TitleByUser returns the preferred title for the given user.
func (character *Character) TitleByUser(user *User) string {
	return character.Name.ByUser(user)
}

// String returns the canonical name of the character.
func (character *Character) String() string {
	return character.Name.Canonical
}

// TypeName returns the type name.
func (character *Character) TypeName() string {
	return "Character"
}

// Self returns the object itself.
func (character *Character) Self() Loggable {
	return character
}

// MainQuote ...
func (character *Character) MainQuote() *Quote {
	quote, _ := GetQuote(character.MainQuoteID)
	return quote
}

// AverageColor returns the average color of the image.
func (character *Character) AverageColor() string {
	color := character.Image.AverageColor

	if color.Hue == 0 && color.Saturation == 0 && color.Lightness == 0 {
		return ""
	}

	return color.String()
}

// ImageLink ...
func (character *Character) ImageLink(size string) string {
	extension := ".jpg"

	if size == "original" {
		extension = character.Image.Extension
	}

	return fmt.Sprintf("//%s/images/characters/%s/%s%s?%v", MediaHost, size, character.ID, extension, character.Image.LastModified)
}

// Publish publishes the character draft.
func (character *Character) Publish() error {
	// No name
	if character.Name.Canonical == "" {
		return errors.New("No canonical character name")
	}

	// No image
	if !character.HasImage() {
		return errors.New("No character image")
	}

	return publish(character)
}

// Unpublish turns the character into a draft.
func (character *Character) Unpublish() error {
	return unpublish(character)
}

// Anime returns a list of all anime the character appears in.
func (character *Character) Anime() []*Anime {
	var results []*Anime

	for animeCharacters := range StreamAnimeCharacters() {
		if animeCharacters.Contains(character.ID) {
			anime, err := GetAnime(animeCharacters.AnimeID)

			if err != nil {
				continue
			}

			results = append(results, anime)
		}
	}

	return results
}

// GetCharacter ...
func GetCharacter(id string) (*Character, error) {
	obj, err := DB.Get("Character", id)

	if err != nil {
		return nil, err
	}

	return obj.(*Character), nil
}

// Merge deletes the character and moves all existing references to the new character.
func (character *Character) Merge(target *Character) {
	// Check anime characters
	for list := range StreamAnimeCharacters() {
		for _, animeCharacter := range list.Items {
			if animeCharacter.CharacterID == character.ID {
				animeCharacter.CharacterID = target.ID
				list.Save()
				break
			}
		}
	}

	// Check quotes
	for quote := range StreamQuotes() {
		if quote.CharacterID == character.ID {
			quote.CharacterID = target.ID
			quote.Save()
		}
	}

	// Check log
	for entry := range StreamEditLogEntries() {
		if entry.ObjectType != "Character" {
			continue
		}

		if entry.ObjectID == character.ID {
			// Delete log entries for the old character
			DB.Delete("EditLogEntry", entry.ID)
		}
	}

	// Merge likes
	for _, userID := range character.Likes {
		if !Contains(target.Likes, userID) {
			target.Likes = append(target.Likes, userID)
		}
	}

	target.Save()

	// Delete image files
	character.DeleteImages()

	// Delete character
	DB.Delete("Character", character.ID)
}

// DeleteImages deletes all images for the character.
func (character *Character) DeleteImages() {
	deleteImages("characters", character.ID, character.Image.Extension)
}

// Quotes returns the list of quotes for this character.
func (character *Character) Quotes() []*Quote {
	return FilterQuotes(func(quote *Quote) bool {
		return !quote.IsDraft && quote.CharacterID == character.ID
	})
}

// SortCharactersByLikes sorts the given slice of characters by the amount of likes.
func SortCharactersByLikes(characters []*Character) {
	sort.Slice(characters, func(i, j int) bool {
		aLikes := len(characters[i].Likes)
		bLikes := len(characters[j].Likes)

		if aLikes == bLikes {
			return characters[i].Name.Canonical < characters[j].Name.Canonical
		}

		return aLikes > bLikes
	})
}

// StreamCharacters returns a stream of all characters.
func StreamCharacters() <-chan *Character {
	channel := make(chan *Character, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("Character") {
			channel <- obj.(*Character)
		}

		close(channel)
	}()

	return channel
}

// FilterCharacters filters all characters by a custom function.
func FilterCharacters(filter func(*Character) bool) []*Character {
	var filtered []*Character

	channel := DB.All("Character")

	for obj := range channel {
		realObject := obj.(*Character)

		if filter(realObject) {
			filtered = append(filtered, realObject)
		}
	}

	return filtered
}

// AllCharacters returns a slice of all characters.
func AllCharacters() []*Character {
	all := make([]*Character, 0, DB.Collection("Character").Count())

	stream := StreamCharacters()

	for obj := range stream {
		all = append(all, obj)
	}

	return all
}
