package arn

import (
	"errors"
	"fmt"
	"sort"

	"github.com/aerogo/nano"
)

// Person represents a person in real life.
type Person struct {
	Name  PersonName  `json:"name" editable:"true"`
	Image PersonImage `json:"image"`

	hasID
	hasPosts
	hasCreator
	hasEditor
	hasLikes
	hasDraft
}

// NewPerson creates a new person.
func NewPerson() *Person {
	return &Person{
		hasID: hasID{
			ID: GenerateID("Person"),
		},
		hasCreator: hasCreator{
			Created: DateTimeUTC(),
		},
	}
}

// Link ...
func (person *Person) Link() string {
	return "/person/" + person.ID
}

// TitleByUser returns the preferred title for the given user.
func (person *Person) TitleByUser(user *User) string {
	return person.Name.ByUser(user)
}

// String returns the default display name for the person.
func (person *Person) String() string {
	return person.Name.ByUser(nil)
}

// TypeName returns the type name.
func (person *Person) TypeName() string {
	return "Person"
}

// Self returns the object itself.
func (person *Person) Self() Loggable {
	return person
}

// ImageLink ...
func (person *Person) ImageLink(size string) string {
	extension := ".jpg"

	if size == "original" {
		extension = person.Image.Extension
	}

	return fmt.Sprintf("//%s/images/persons/%s/%s%s?%v", MediaHost, size, person.ID, extension, person.Image.LastModified)
}

// Publish publishes the person draft.
func (person *Person) Publish() error {
	// No name
	if person.Name.ByUser(nil) == "" {
		return errors.New("No person name")
	}

	// No image
	if !person.HasImage() {
		return errors.New("No person image")
	}

	return publish(person)
}

// Unpublish turns the person into a draft.
func (person *Person) Unpublish() error {
	return unpublish(person)
}

// HasImage returns true if the person has an image.
func (person *Person) HasImage() bool {
	return person.Image.Extension != "" && person.Image.Width > 0
}

// GetPerson ...
func GetPerson(id string) (*Person, error) {
	obj, err := DB.Get("Person", id)

	if err != nil {
		return nil, err
	}

	return obj.(*Person), nil
}

// DeleteImages deletes all images for the person.
func (person *Person) DeleteImages() {
	deleteImages("persons", person.ID, person.Image.Extension)
}

// SortPersonsByLikes sorts the given slice of persons by the amount of likes.
func SortPersonsByLikes(persons []*Person) {
	sort.Slice(persons, func(i, j int) bool {
		aLikes := len(persons[i].Likes)
		bLikes := len(persons[j].Likes)

		if aLikes == bLikes {
			return persons[i].Name.English.First < persons[j].Name.English.First
		}

		return aLikes > bLikes
	})
}

// StreamPersons returns a stream of all persons.
func StreamPersons() <-chan *Person {
	channel := make(chan *Person, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("Person") {
			channel <- obj.(*Person)
		}

		close(channel)
	}()

	return channel
}

// FilterPersons filters all persons by a custom function.
func FilterPersons(filter func(*Person) bool) []*Person {
	var filtered []*Person

	channel := DB.All("Person")

	for obj := range channel {
		realObject := obj.(*Person)

		if filter(realObject) {
			filtered = append(filtered, realObject)
		}
	}

	return filtered
}

// AllPersons returns a slice of all persons.
func AllPersons() []*Person {
	all := make([]*Person, 0, DB.Collection("Person").Count())

	stream := StreamPersons()

	for obj := range stream {
		all = append(all, obj)
	}

	return all
}
