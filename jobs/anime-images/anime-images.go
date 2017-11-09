package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/aerogo/flow/jobqueue"
	"github.com/aerogo/ipo"
	"github.com/aerogo/ipo/inputs"
	"github.com/aerogo/ipo/outputs"
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

var ticker = time.NewTicker(50 * time.Millisecond)

func main() {
	color.Yellow("Downloading anime images")
	defer arn.Node.Close()

	jobs := jobqueue.New(work)

	for anime := range arn.StreamAnime() {
		jobs.Queue(anime)
	}

	results := jobs.Wait()
	color.Green("Finished downloading %d anime images.", len(results))

	// Give file buffers some time, just to be safe
	time.Sleep(time.Second)
}

func work(job interface{}) interface{} {
	anime := job.(*arn.Anime)

	if !strings.HasPrefix(anime.Image.Original, "//media.kitsu.io/anime/") {
		return nil
	}

	<-ticker.C

	originals := path.Join(os.Getenv("GOPATH"), "/src/github.com/animenotifier/notify.moe/images/anime/original/")
	large := path.Join(os.Getenv("GOPATH"), "/src/github.com/animenotifier/notify.moe/images/anime/large/")
	medium := path.Join(os.Getenv("GOPATH"), "/src/github.com/animenotifier/notify.moe/images/anime/medium/")
	small := path.Join(os.Getenv("GOPATH"), "/src/github.com/animenotifier/notify.moe/images/anime/small/")

	largeSize := 250
	mediumSize := 142
	smallSize := 55

	webpQuality := 80
	jpegQuality := 80

	system := &ipo.System{
		Inputs: []ipo.Input{
			&inputs.NetworkImage{
				URL: anime.Image.Original,
			},
		},
		Outputs: []ipo.Output{
			// Original
			&outputs.ImageFile{
				Directory: originals,
				BaseName:  anime.ID,
			},

			// Large
			&outputs.ImageFile{
				Directory: large,
				BaseName:  anime.ID,
				Size:      largeSize,
				Quality:   jpegQuality,
			},
			&outputs.ImageFile{
				Directory: large,
				BaseName:  anime.ID + "@2",
				Size:      largeSize * 2,
				Quality:   jpegQuality,
			},
			&outputs.ImageFile{
				Directory: large,
				BaseName:  anime.ID,
				Size:      largeSize,
				Format:    "webp",
				Quality:   webpQuality,
			},
			&outputs.ImageFile{
				Directory: large,
				BaseName:  anime.ID + "@2",
				Size:      largeSize * 2,
				Format:    "webp",
				Quality:   webpQuality,
			},

			// Medium
			&outputs.ImageFile{
				Directory: medium,
				BaseName:  anime.ID,
				Size:      mediumSize,
				Quality:   jpegQuality,
			},
			&outputs.ImageFile{
				Directory: medium,
				BaseName:  anime.ID + "@2",
				Size:      mediumSize * 2,
				Quality:   jpegQuality,
			},
			&outputs.ImageFile{
				Directory: medium,
				BaseName:  anime.ID,
				Size:      mediumSize,
				Format:    "webp",
				Quality:   webpQuality,
			},
			&outputs.ImageFile{
				Directory: medium,
				BaseName:  anime.ID + "@2",
				Size:      mediumSize * 2,
				Format:    "webp",
				Quality:   webpQuality,
			},

			// Small
			&outputs.ImageFile{
				Directory: small,
				BaseName:  anime.ID,
				Size:      smallSize,
				Quality:   jpegQuality,
			},
			&outputs.ImageFile{
				Directory: small,
				BaseName:  anime.ID + "@2",
				Size:      smallSize * 2,
				Quality:   jpegQuality,
			},
			&outputs.ImageFile{
				Directory: small,
				BaseName:  anime.ID,
				Size:      smallSize,
				Format:    "webp",
				Quality:   webpQuality,
			},
			&outputs.ImageFile{
				Directory: small,
				BaseName:  anime.ID + "@2",
				Size:      smallSize * 2,
				Format:    "webp",
				Quality:   webpQuality,
			},
		},
		InputProcessor:  ipo.SequentialInputs,
		OutputProcessor: ipo.ParallelOutputs,
	}

	err := system.Run()

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
