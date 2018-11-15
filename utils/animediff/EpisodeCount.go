package animediff

import "strconv"

// EpisodeCount ...
type EpisodeCount struct {
	EpisodesA   int
	EpisodesB   int
	NumericHash uint64
}

// TypeName returns the diff type.
func (diff *EpisodeCount) TypeName() string {
	return "EpisodeCount"
}

// Explanation returns the description.
func (diff *EpisodeCount) Explanation() string {
	return "Episode counts are different"
}

// DetailsA shows the details for the first anime.
func (diff *EpisodeCount) DetailsA() string {
	return strconv.Itoa(diff.EpisodesA)
}

// DetailsB shows the details for the second anime.
func (diff *EpisodeCount) DetailsB() string {
	return strconv.Itoa(diff.EpisodesB)
}

// Hash returns the hash for the suggested value (from anime B).
func (diff *EpisodeCount) Hash() uint64 {
	return diff.NumericHash
}
