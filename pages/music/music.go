package music

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxTracks = 10

// Get renders the music page.
func Get(ctx *aero.Context) string {
	tracks := []*arn.SoundTrack{}

	tracks = append(tracks, &arn.SoundTrack{
		ID: arn.GenerateID("SoundTrack"),
		Media: []arn.ExternalMedia{
			arn.ExternalMedia{
				Service:   "Soundcloud",
				ServiceID: "145918628",
			},
		},
		Tags: []string{
			"anime:2357",
		},
		Created:   arn.DateTimeUTC(),
		CreatedBy: "4J6qpK1ve",
	})

	tracks = append(tracks, &arn.SoundTrack{
		ID: arn.GenerateID("SoundTrack"),
		Media: []arn.ExternalMedia{
			arn.ExternalMedia{
				Service:   "Soundcloud",
				ServiceID: "127672476",
			},
		},
		Tags: []string{
			"anime:7622",
		},
		Created:   arn.DateTimeUTC(),
		CreatedBy: "4J6qpK1ve",
	})

	tracks = append(tracks, &arn.SoundTrack{
		ID: arn.GenerateID("SoundTrack"),
		Media: []arn.ExternalMedia{
			arn.ExternalMedia{
				Service:   "Soundcloud",
				ServiceID: "270777538",
			},
		},
		Tags: []string{
			"anime:11469",
		},
		Created:   arn.DateTimeUTC(),
		CreatedBy: "4J6qpK1ve",
	})

	tracks = append(tracks, &arn.SoundTrack{
		ID: arn.GenerateID("SoundTrack"),
		Media: []arn.ExternalMedia{
			arn.ExternalMedia{
				Service:   "Soundcloud",
				ServiceID: "243839100",
			},
		},
		Tags: []string{
			"anime:9962",
		},
		Created:   arn.DateTimeUTC(),
		CreatedBy: "4J6qpK1ve",
	})

	tracks = append(tracks, &arn.SoundTrack{
		ID: arn.GenerateID("SoundTrack"),
		Media: []arn.ExternalMedia{
			arn.ExternalMedia{
				Service:   "Soundcloud",
				ServiceID: "207355237",
			},
		},
		Tags: []string{
			"anime:6589",
		},
		Created:   arn.DateTimeUTC(),
		CreatedBy: "4J6qpK1ve",
	})

	tracks = append(tracks, &arn.SoundTrack{
		ID: arn.GenerateID("SoundTrack"),
		Media: []arn.ExternalMedia{
			arn.ExternalMedia{
				Service:   "Soundcloud",
				ServiceID: "242172944",
			},
		},
		Tags: []string{
			"anime:10740",
		},
		Created:   arn.DateTimeUTC(),
		CreatedBy: "4J6qpK1ve",
	})

	if len(tracks) > maxTracks {
		tracks = tracks[:maxTracks]
	}

	sort.Slice(tracks, func(i, j int) bool {
		return tracks[i].Created > tracks[j].Created
	})

	return ctx.HTML(components.Music(tracks))
}
