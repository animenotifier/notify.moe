package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/aerogo/ipo"
	"github.com/aerogo/ipo/inputs"
	"github.com/aerogo/ipo/outputs"
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

var ticker = time.NewTicker(50 * time.Millisecond)

// Shell parameters
var from int
var to int

// Shell flags
func init() {
	flag.IntVar(&from, "from", 0, "From index")
	flag.IntVar(&to, "to", 0, "To index")
	flag.Parse()
}

func main() {
	color.Yellow("Downloading anime images")
	defer arn.Node.Close()

	if from < 0 {
		from = 0
	}

	allAnime := arn.FilterAnime(func(anime *arn.Anime) bool {
		id, _ := strconv.Atoi(anime.ID)
		return id >= from && id <= to
	})

	for index, anime := range allAnime {
		fmt.Printf("%d / %d\n", index+1, len(allAnime))
		work(anime)
	}

	color.Green("Finished downloading anime images.")

	// Give file buffers some time, just to be safe
	time.Sleep(time.Second)
}

func work(anime *arn.Anime) error {
	<-ticker.C

	originals := path.Join(os.Getenv("GOPATH"), "/src/github.com/animenotifier/notify.moe/images/anime/original/")
	large := path.Join(os.Getenv("GOPATH"), "/src/github.com/animenotifier/notify.moe/images/anime/large/")
	medium := path.Join(os.Getenv("GOPATH"), "/src/github.com/animenotifier/notify.moe/images/anime/medium/")
	small := path.Join(os.Getenv("GOPATH"), "/src/github.com/animenotifier/notify.moe/images/anime/small/")

	largeSizeX := 250
	largeSizeY := 0

	mediumSizeX := 142
	mediumSizeY := 200

	smallSizeX := 55
	smallSizeY := 78

	webpQuality := 70
	jpegQuality := 70

	qualityBonusLowDPI := 10
	qualityBonusMedium := 10
	qualityBonusSmall := 10

	kitsuOriginal := fmt.Sprintf("https://media.kitsu.io/anime/poster_images/%s/original", anime.ID)

	system := ipo.System{
		Inputs: []ipo.Input{
			&inputs.FileSystemImage{
				URL: path.Join(originals, anime.ID+".png"),
			},
			&inputs.FileSystemImage{
				URL: path.Join(originals, anime.ID+".jpg"),
			},
			&inputs.FileSystemImage{
				URL: path.Join(originals, anime.ID+".jpeg"),
			},
			&inputs.FileSystemImage{
				URL: path.Join(originals, anime.ID+".gif"),
			},
			&inputs.NetworkImage{
				URL: kitsuOriginal + anime.ImageExtension,
			},
			&inputs.NetworkImage{
				URL: kitsuOriginal + ".png",
			},
			&inputs.NetworkImage{
				URL: kitsuOriginal + ".jpg",
			},
			&inputs.NetworkImage{
				URL: kitsuOriginal + ".jpeg",
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
				Width:     largeSizeX,
				Height:    largeSizeY,
				Quality:   jpegQuality + qualityBonusLowDPI,
			},
			&outputs.ImageFile{
				Directory: large,
				BaseName:  anime.ID + "@2",
				Width:     largeSizeX * 2,
				Height:    largeSizeY * 2,
				Quality:   jpegQuality,
			},
			&outputs.ImageFile{
				Directory: large,
				BaseName:  anime.ID,
				Width:     largeSizeX,
				Height:    largeSizeY,
				Format:    "webp",
				Quality:   webpQuality + qualityBonusLowDPI,
			},
			&outputs.ImageFile{
				Directory: large,
				BaseName:  anime.ID + "@2",
				Width:     largeSizeX * 2,
				Height:    largeSizeY * 2,
				Format:    "webp",
				Quality:   webpQuality,
			},

			// Medium
			&outputs.ImageFile{
				Directory: medium,
				BaseName:  anime.ID,
				Width:     mediumSizeX,
				Height:    mediumSizeY,
				Quality:   jpegQuality + qualityBonusLowDPI + qualityBonusMedium,
			},
			&outputs.ImageFile{
				Directory: medium,
				BaseName:  anime.ID + "@2",
				Width:     mediumSizeX * 2,
				Height:    mediumSizeY * 2,
				Quality:   jpegQuality,
			},
			&outputs.ImageFile{
				Directory: medium,
				BaseName:  anime.ID,
				Width:     mediumSizeX,
				Height:    mediumSizeY,
				Format:    "webp",
				Quality:   webpQuality + qualityBonusLowDPI + qualityBonusMedium,
			},
			&outputs.ImageFile{
				Directory: medium,
				BaseName:  anime.ID + "@2",
				Width:     mediumSizeX * 2,
				Height:    mediumSizeY * 2,
				Format:    "webp",
				Quality:   webpQuality,
			},

			// Small
			&outputs.ImageFile{
				Directory: small,
				BaseName:  anime.ID,
				Width:     smallSizeX,
				Height:    smallSizeY,
				Quality:   jpegQuality + qualityBonusLowDPI + qualityBonusSmall,
			},
			&outputs.ImageFile{
				Directory: small,
				BaseName:  anime.ID + "@2",
				Width:     smallSizeX * 2,
				Height:    smallSizeY * 2,
				Quality:   jpegQuality,
			},
			&outputs.ImageFile{
				Directory: small,
				BaseName:  anime.ID,
				Width:     smallSizeX,
				Height:    smallSizeY,
				Format:    "webp",
				Quality:   webpQuality + qualityBonusLowDPI + qualityBonusSmall,
			},
			&outputs.ImageFile{
				Directory: small,
				BaseName:  anime.ID + "@2",
				Width:     smallSizeX * 2,
				Height:    smallSizeY * 2,
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

	// Try to free up some memory
	system.Inputs = nil
	system.Outputs = nil
	runtime.GC()

	return nil
}
