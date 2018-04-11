package main

func main() {
	// color.Yellow("Generating character IDs")

	// defer color.Green("Finished")
	// defer arn.Node.Close()

	// sort.Slice(allCharacters, func(i, j int) bool {
	// 	aID, _ := strconv.Atoi(allCharacters[i].ID)
	// 	bID, _ := strconv.Atoi(allCharacters[j].ID)

	// 	return aID < bID
	// })

	// // Create map of old IDs to new IDs
	// idMap := map[string]string{}

	// for counter, character := range allCharacters {
	// 	newID := arn.GenerateID("Character")
	// 	fmt.Printf("[%d / %d] Old [%s] New [%s] %s\n", counter+1, len(allCharacters), color.YellowString(character.ID), color.GreenString(newID), character)
	// 	arn.DB.Delete("Character", character.ID)
	// 	idMap[character.ID] = newID
	// 	character.ID = newID
	// 	character.Save()
	// }

	// // Update quotes
	// for quote := range arn.StreamQuotes() {
	// 	newID, exists := idMap[quote.CharacterID]

	// 	if exists {
	// 		quote.CharacterID = newID
	// 		quote.Save()
	// 	}
	// }

	// // Update log
	// for entry := range arn.StreamEditLogEntries() {
	// 	if entry.ObjectType != "Character" {
	// 		continue
	// 	}

	// 	newID, exists := idMap[entry.ObjectID]

	// 	if exists {
	// 		entry.ObjectID = newID
	// 		entry.Save()
	// 	}
	// }

	// // Update anime characters
	// for list := range arn.StreamAnimeCharacters() {
	// 	modified := false

	// 	for _, animeCharacter := range list.Items {
	// 		newID, exists := idMap[animeCharacter.CharacterID]

	// 		if exists {
	// 			animeCharacter.CharacterID = newID
	// 			modified = true
	// 		}
	// 	}

	// 	if modified {
	// 		list.Save()
	// 	}
	// }
}
