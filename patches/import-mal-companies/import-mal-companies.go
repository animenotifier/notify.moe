package main

import (
	"fmt"
	"time"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/mal"
	"github.com/fatih/color"
)

var malDB = arn.Node.Namespace("mal").RegisterTypes((*mal.Anime)(nil))
var companies = map[string]*arn.Company{}
var now = time.Now()

func main() {
	defer arn.Node.Close()
	color.Yellow("Importing companies")

	for company := range arn.StreamCompanies() {
		companies[company.Name.English] = company
	}

	for anime := range arn.StreamAnime() {
		malID := anime.GetMapping("myanimelist/anime")

		if malID == "" {
			continue
		}

		importCompanies(anime, malID)
	}

	for name, company := range companies {
		fmt.Println(name)
		company.Save()
	}

	color.Green("Finished importing %d companies", len(companies))
}

func importCompanies(anime *arn.Anime, malID string) {
	obj, err := malDB.Get("Anime", malID)

	if err != nil {
		fmt.Println(err)
		return
	}

	malAnime := obj.(*mal.Anime)

	for _, producer := range malAnime.Studios {
		importByName(anime, "studio", producer)
	}

	for _, producer := range malAnime.Producers {
		importByName(anime, "producer", producer)
	}

	for _, producer := range malAnime.Licensors {
		importByName(anime, "licensor", producer)
	}
}

func importByName(anime *arn.Anime, companyType string, producer *mal.Producer) {
	company, exists := companies[producer.Name]

	if !exists {
		now = now.Add(-time.Second)

		company = &arn.Company{
			ID: arn.GenerateID("Company"),
			Name: arn.CompanyName{
				English: producer.Name,
			},
			Created:   now.UTC().Format(time.RFC3339),
			CreatedBy: "",
			Mappings: []*arn.Mapping{
				&arn.Mapping{
					Service:   "myanimelist/producer",
					ServiceID: producer.ID,
					Created:   arn.DateTimeUTC(),
					CreatedBy: "",
				},
			},
			Links: []*arn.Link{},
			Tags:  []string{},
			Likes: []string{},
		}

		companies[producer.Name] = company
	}

	switch companyType {
	case "studio":
		anime.StudioIDs = append(anime.StudioIDs, company.ID)
	case "producer":
		anime.ProducerIDs = append(anime.ProducerIDs, company.ID)
	case "licensor":
		anime.LicensorIDs = append(anime.LicensorIDs, company.ID)
	}

	anime.Save()
}
