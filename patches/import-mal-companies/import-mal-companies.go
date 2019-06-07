package main

import (
	"fmt"
	"time"

	"github.com/akyoto/color"
	"github.com/animenotifier/mal"
	"github.com/animenotifier/notify.moe/arn"
)

var malDB = arn.Node.Namespace("mal").RegisterTypes((*mal.Anime)(nil))
var companies = map[string]*arn.Company{}
var now = time.Now()

func main() {
	defer arn.Node.Close()
	color.Yellow("Importing companies")

	for company := range arn.StreamCompanies() {
		malID := company.GetMapping("myanimelist/producer")

		if malID == "" {
			continue
		}

		companies[malID] = company
	}

	for anime := range arn.StreamAnime() {
		malID := anime.GetMapping("myanimelist/anime")

		if malID == "" {
			continue
		}

		importCompanies(anime, malID)
	}

	color.Green("Finished importing %d companies", len(companies))
}

func importCompanies(anime *arn.Anime, malID string) {
	obj, err := malDB.Get("Anime", malID)

	if err != nil {
		fmt.Printf("%s: %s (invalid MAL ID)\n", color.YellowString(anime.ID), color.RedString(malID))
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
	company, exists := companies[producer.ID]

	// If we encounter a company that has not been added yet, save the new company
	if !exists {
		fmt.Println("Adding new company:", producer.Name)

		// Subtract one second every time we create a new company
		// so that they don't all end up with the same creation date.
		now = now.Add(-time.Second)

		// Create new company
		company = arn.NewCompany()
		company.Name.English = producer.Name
		company.Created = now.UTC().Format(time.RFC3339)
		company.SetMapping("myanimelist/producer", producer.ID)
		company.Save()

		// Add company to the global ID map
		companies[producer.ID] = company
	}

	switch companyType {
	case "studio":
		anime.AddStudio(company.ID)
	case "producer":
		anime.AddProducer(company.ID)
	case "licensor":
		anime.AddLicensor(company.ID)
	}

	anime.Save()
}
