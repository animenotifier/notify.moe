package arn

import (
	"errors"

	"github.com/aerogo/nano"
)

// Company represents an anime studio, producer or licensor.
type Company struct {
	Name        CompanyName `json:"name" editable:"true"`
	Description string      `json:"description" editable:"true" type:"textarea"`
	Email       string      `json:"email" editable:"true"`
	Links       []*Link     `json:"links" editable:"true"`

	// Mixins
	hasID
	hasMappings
	hasLikes
	hasDraft

	// Other editable fields
	Location *Location `json:"location" editable:"true"`
	Tags     []string  `json:"tags" editable:"true"`

	// Editing dates
	hasCreator
	hasEditor
}

// NewCompany creates a new company.
func NewCompany() *Company {
	return &Company{
		hasID: hasID{
			ID: GenerateID("Company"),
		},
		Name:  CompanyName{},
		Links: []*Link{},
		Tags:  []string{},
		hasCreator: hasCreator{
			Created: DateTimeUTC(),
		},
		hasMappings: hasMappings{
			Mappings: []*Mapping{},
		},
	}
}

// Link returns a single company.
func (company *Company) Link() string {
	return "/company/" + company.ID
}

// Anime returns the anime connected with this company.
func (company *Company) Anime() (studioAnime []*Anime, producedAnime []*Anime, licensedAnime []*Anime) {
	for anime := range StreamAnime() {
		if Contains(anime.StudioIDs, company.ID) {
			studioAnime = append(studioAnime, anime)
		}

		if Contains(anime.ProducerIDs, company.ID) {
			producedAnime = append(producedAnime, anime)
		}

		if Contains(anime.LicensorIDs, company.ID) {
			licensedAnime = append(licensedAnime, anime)
		}
	}

	SortAnimeByQuality(studioAnime)
	SortAnimeByQuality(producedAnime)
	SortAnimeByQuality(licensedAnime)

	return studioAnime, producedAnime, licensedAnime
}

// Publish publishes the company draft.
func (company *Company) Publish() error {
	// No title
	if company.Name.English == "" {
		return errors.New("No English company name")
	}

	return publish(company)
}

// Unpublish turns the company into a draft.
func (company *Company) Unpublish() error {
	return unpublish(company)
}

// String implements the default string serialization.
func (company *Company) String() string {
	return company.Name.English
}

// TypeName returns the type name.
func (company *Company) TypeName() string {
	return "Company"
}

// Self returns the object itself.
func (company *Company) Self() Loggable {
	return company
}

// GetCompany returns a single company.
func GetCompany(id string) (*Company, error) {
	obj, err := DB.Get("Company", id)

	if err != nil {
		return nil, err
	}

	return obj.(*Company), nil
}

// StreamCompanies returns a stream of all companies.
func StreamCompanies() <-chan *Company {
	channel := make(chan *Company, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("Company") {
			channel <- obj.(*Company)
		}

		close(channel)
	}()

	return channel
}

// FilterCompanies filters all companies by a custom function.
func FilterCompanies(filter func(*Company) bool) []*Company {
	var filtered []*Company

	channel := DB.All("Company")

	for obj := range channel {
		realObject := obj.(*Company)

		if filter(realObject) {
			filtered = append(filtered, realObject)
		}
	}

	return filtered
}

// AllCompanies returns a slice of all companies.
func AllCompanies() []*Company {
	all := make([]*Company, 0, DB.Collection("Company").Count())

	stream := StreamCompanies()

	for obj := range stream {
		all = append(all, obj)
	}

	return all
}
