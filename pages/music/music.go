package music

import "github.com/aerogo/aero"
import "github.com/animenotifier/notify.moe/components"
import "github.com/animenotifier/arn"

// Get renders the music page.
func Get(ctx *aero.Context) string {
	tracks := []*arn.SoundTrack{}

	tracks = append(tracks, &arn.SoundTrack{
		ID: "1",
		Media: []arn.ExternalMedia{
			arn.ExternalMedia{
				Service:   "Soundcloud",
				ServiceID: "127672476",
			},
		},
		Tags: []string{
			"anime:7622",
		},
	})

	tracks = append(tracks, &arn.SoundTrack{
		ID: "2",
		Media: []arn.ExternalMedia{
			arn.ExternalMedia{
				Service:   "Soundcloud",
				ServiceID: "270777538",
			},
		},
		Tags: []string{
			"anime:11469",
		},
	})

	return ctx.HTML(components.Music(tracks))
}
