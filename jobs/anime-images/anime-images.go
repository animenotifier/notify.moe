package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/aerogo/flow/jobqueue"
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
	"github.com/parnurzeal/gorequest"
)

var ticker = time.NewTicker(50 * time.Millisecond)

func main() {
	color.Yellow("Downloading anime images")
	jobs := jobqueue.New(work)
	allAnime, _ := arn.AllAnime()

	for _, anime := range allAnime {
		jobs.Queue(anime)
	}

	results := jobs.Wait()
	color.Green("Finished downloading %d anime images.", len(results))
}

func work(job interface{}) interface{} {
	anime := job.(*arn.Anime)

	if !strings.HasPrefix(anime.Image.Original, "//media.kitsu.io/anime/") {
		return nil
	}

	<-ticker.C
	resp, body, errs := gorequest.New().Get(anime.Image.Original).End()

	if len(errs) > 0 {
		color.Red(errs[0].Error())
		return errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		color.Red("Status %d", resp.StatusCode)
	}

	extension := anime.Image.Original[strings.LastIndex(anime.Image.Original, "."):]
	fileName := "anime/" + anime.ID + extension
	fmt.Println(fileName)

	ioutil.WriteFile(fileName, []byte(body), 0644)

	return nil
}
