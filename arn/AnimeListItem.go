package arn

// AnimeListStatus values for anime list items
const (
	AnimeListStatusWatching  = "watching"
	AnimeListStatusCompleted = "completed"
	AnimeListStatusPlanned   = "planned"
	AnimeListStatusHold      = "hold"
	AnimeListStatusDropped   = "dropped"
)

// AnimeListItem ...
type AnimeListItem struct {
	AnimeID      string              `json:"animeId"`
	Status       string              `json:"status" editable:"true"`
	Episodes     int                 `json:"episodes" editable:"true"`
	Rating       AnimeListItemRating `json:"rating"`
	Notes        string              `json:"notes" editable:"true"`
	RewatchCount int                 `json:"rewatchCount" editable:"true"`
	Private      bool                `json:"private" editable:"true"`
	Created      string              `json:"created"`
	Edited       string              `json:"edited"`
}

// Anime fetches the associated anime data.
func (item *AnimeListItem) Anime() *Anime {
	anime, _ := GetAnime(item.AnimeID)
	return anime
}

// Link returns the URI for the given item.
func (item *AnimeListItem) Link(userNick string) string {
	return "/+" + userNick + "/animelist/anime/" + item.AnimeID
}

// StatusHumanReadable returns the human readable representation of the status.
func (item *AnimeListItem) StatusHumanReadable() string {
	switch item.Status {
	case AnimeListStatusWatching:
		return "Watching"
	case AnimeListStatusCompleted:
		return "Completed"
	case AnimeListStatusPlanned:
		return "Planned"
	case AnimeListStatusHold:
		return "On Hold"
	case AnimeListStatusDropped:
		return "Dropped"
	default:
		return "Unknown"
	}
}

// OnEpisodesChange is called when the watched episode count changes.
func (item *AnimeListItem) OnEpisodesChange() {
	maxEpisodesKnown := item.Anime().EpisodeCount != 0

	// If we update episodes to the max, set status to completed automatically.
	if item.Anime().Status == "finished" && maxEpisodesKnown && item.Episodes == item.Anime().EpisodeCount {
		// Complete automatically.
		item.Status = AnimeListStatusCompleted
	}

	// We set episodes lower than the max but the status is set as completed.
	if item.Status == AnimeListStatusCompleted && maxEpisodesKnown && item.Episodes < item.Anime().EpisodeCount {
		// Set status back to watching.
		item.Status = AnimeListStatusWatching
	}

	// If we increase the episodes and status is planned, set it to watching.
	if item.Status == AnimeListStatusPlanned && item.Episodes > 0 {
		// Set status to watching.
		item.Status = AnimeListStatusWatching
	}

	// If we set the episodes to 0 and status is not planned or dropped, set it to planned.
	if item.Episodes == 0 && (item.Status != AnimeListStatusPlanned && item.Status != AnimeListStatusDropped) {
		// Set status to planned.
		item.Status = AnimeListStatusPlanned
	}
}

// OnStatusChange is called when the status changes.
func (item *AnimeListItem) OnStatusChange() {
	maxEpisodesKnown := item.Anime().EpisodeCount != 0

	// We just switched to completed status but the episodes aren't max yet.
	if item.Status == AnimeListStatusCompleted && maxEpisodesKnown && item.Episodes < item.Anime().EpisodeCount {
		// Set episodes to max.
		item.Episodes = item.Anime().EpisodeCount
	}

	// We just switched to plan to watch status but the episodes are greater than zero.
	if item.Status == AnimeListStatusPlanned && item.Episodes > 0 {
		// Set episodes back to zero.
		item.Episodes = 0
	}

	// If we have an anime with max episodes watched and we change status to not completed, lower episode count by 1.
	if maxEpisodesKnown && item.Status != AnimeListStatusCompleted && item.Episodes == item.Anime().EpisodeCount {
		// Lower episodes by 1.
		item.Episodes--
	}
}
