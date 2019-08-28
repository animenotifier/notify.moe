package arn

import (
	"sort"
	"strconv"
	"strings"
)

// EpisodeList is a list of episodes.
type EpisodeList []*Episode

// Sort sorts the episodes by episode number.
func (episodes EpisodeList) Sort() {
	sort.Slice(episodes, func(i, j int) bool {
		return episodes[i].Number < episodes[j].Number
	})
}

// Find finds the given episode number.
func (episodes EpisodeList) Find(episodeNumber int) (*Episode, int) {
	for index, episode := range episodes {
		if episode.Number == episodeNumber {
			return episode, index
		}
	}

	return nil, -1
}

// Merge combines the data of both episode lists to one.
func (episodes EpisodeList) Merge(b EpisodeList) EpisodeList {
	for index, episode := range b {
		if index >= len(episodes) {
			episodes = append(episodes, episode)
		} else {
			episodes[index].Merge(episode)
		}
	}

	return episodes
}

// HumanReadable returns a text representation of the anime episodes.
func (episodes EpisodeList) HumanReadable() string {
	b := strings.Builder{}

	for _, episode := range episodes {
		b.WriteString(strconv.Itoa(episode.Number))
		b.WriteString(" | ")
		b.WriteString(episode.Title.Japanese)
		b.WriteString(" | ")
		b.WriteString(episode.AiringDate.StartDateHuman())
		b.WriteByte('\n')
	}

	return strings.TrimRight(b.String(), "\n")
}

// AvailableCount counts the number of available episodes.
func (episodes EpisodeList) AvailableCount() int {
	available := 0

	for _, episode := range episodes {
		if episode.Available() {
			available++
		}
	}

	return available
}
