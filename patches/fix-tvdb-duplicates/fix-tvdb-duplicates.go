package main

import (
	"fmt"
	"strings"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	defer arn.Node.Close()

	for anime := range arn.StreamAnime() {
		existing := map[string]string{}

		for index, mapping := range anime.Mappings {
			if mapping.Service != "thetvdb/anime" {
				continue
			}

			serviceID, exists := existing[mapping.Service]

			if exists {
				fmt.Println("duplicate:", color.YellowString(mapping.ServiceID), "of", color.YellowString(serviceID))

				slashPos := strings.Index(serviceID, "/")

				if slashPos != -1 {
					serviceID = serviceID[:slashPos]
				}

				slashPos = strings.Index(mapping.ServiceID, "/")

				if slashPos != -1 {
					mapping.ServiceID = mapping.ServiceID[:slashPos]
				}

				if serviceID == mapping.ServiceID {
					// Remove duplicate
					anime.Mappings = append(anime.Mappings[:index], anime.Mappings[index+1:]...)
					anime.Save()
					break
				}
			}

			existing[mapping.Service] = mapping.ServiceID
		}
	}

	color.Green("Finished.")
}
