package arn

// AnimeFinder holds an internal map of ID to anime mappings
// and is therefore very efficient to use when trying to find
// anime by a given service and ID.
type AnimeFinder struct {
	idToAnime map[string]*Anime
}

// NewAnimeFinder creates a new finder for external anime.
func NewAnimeFinder(mappingName string) *AnimeFinder {
	finder := &AnimeFinder{
		idToAnime: map[string]*Anime{},
	}

	for anime := range StreamAnime() {
		id := anime.GetMapping(mappingName)

		if id != "" {
			finder.idToAnime[id] = anime
		}
	}

	return finder
}

// GetAnime tries to find an external anime in our anime database.
func (finder *AnimeFinder) GetAnime(id string) *Anime {
	return finder.idToAnime[id]
}
