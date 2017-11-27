package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/jikan"
	"github.com/fatih/color"
)

var jikanDB = arn.Node.Namespace("jikan")
var companies = map[string]*arn.Company{}
var now = time.Now()

func main() {
	defer arn.Node.Close()
	color.Yellow("Importing companies")

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
	time.Sleep(2 * time.Second)
}

func importCompanies(anime *arn.Anime, malID string) {
	obj, err := jikanDB.Get("Anime", malID)

	if err != nil {
		fmt.Println(err)
		return
	}

	jikanAnime := obj.(*jikan.Anime)

	for _, studioInfo := range jikanAnime.Studio {
		studioName := studioInfo[1]
		htmlPos := strings.Index(studioName, "<")

		if htmlPos != -1 {
			studioName = studioName[:htmlPos]
		}

		company, exists := companies[studioName]

		if !exists {
			now = now.Add(-time.Second)

			company = &arn.Company{
				ID: arn.GenerateID("Company"),
				Name: arn.CompanyName{
					English: studioName,
				},
				Created:   now.UTC().Format(time.RFC3339),
				CreatedBy: "",
				Mappings:  []*arn.Mapping{},
				Links:     []*arn.Link{},
				Tags:      []string{},
				Likes:     []string{},
			}

			companies[studioName] = company
		}

		anime.StudioIDs = append(anime.StudioIDs, company.ID)
		anime.Save()
	}
}
