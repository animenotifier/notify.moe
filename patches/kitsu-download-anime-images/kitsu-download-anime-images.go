package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/aerogo/http/client"
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
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

	defer color.Green("Finished.")
	defer arn.Node.Close()

	if from < 0 {
		from = 0
	}

	allAnime := arn.FilterAnime(func(anime *arn.Anime) bool {
		id, _ := strconv.Atoi(anime.GetMapping("kitsu/anime"))
		return id >= from && id <= to
	})

	for index, anime := range allAnime {
		fmt.Printf("%d / %d\n", index+1, len(allAnime))
		err := work(anime)

		if err != nil {
			color.Red(err.Error())
		}
	}

	// Give file buffers some time, just to be safe
	time.Sleep(time.Second)
}

func work(anime *arn.Anime) error {
	<-ticker.C

	kitsuOriginal := fmt.Sprintf("https://media.kitsu.io/anime/poster_images/%s/original%s", anime.GetMapping("kitsu/anime"), anime.Image.Extension)

	// Download kitsu image
	response, err := client.Get(kitsuOriginal).End()

	if err != nil {
		return fmt.Errorf("%s (%s)", err.Error(), kitsuOriginal)
	}

	if response.StatusCode() != http.StatusOK {
		return fmt.Errorf("Status %d (%s)", response.StatusCode(), kitsuOriginal)
	}

	err = anime.SetImageBytes(response.Bytes())

	if err != nil {
		return err
	}

	// Try to free up some memory
	runtime.GC()

	return nil
}
