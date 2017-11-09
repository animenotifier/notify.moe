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
}

func work(job interface{}) interface{} {
	anime := job.(*arn.Anime)

	if !strings.HasPrefix(anime.Image.Original, "//media.kitsu.io/anime/") {
		return nil
	}

	<-ticker.C
	// resp, body, errs := gorequest.New().Get(anime.Image.Original).End()

	// if len(errs) > 0 {
	// 	color.Red(errs[0].Error())
	// 	return errs[0]
	// }

	// if resp.StatusCode != http.StatusOK {
	// 	color.Red("Status %d", resp.StatusCode)
	// }

	// extension := anime.Image.Original[strings.LastIndex(anime.Image.Original, "."):]
	// fileName := "anime/" + anime.ID + extension
	// fmt.Println(fileName)

	// ioutil.WriteFile(fileName, []byte(body), 0644)

	originals := path.Join(os.Getenv("GOPATH"), "/src/github.com/animenotifier/notify.moe/images/anime/original/")

	system := &ipo.System{
		Inputs: []ipo.Input{
			&inputs.NetworkImage{
				URL: anime.Image.Original,
			},
		},
		Outputs: []ipo.Output{
			&outputs.ImageFile{
				Directory: originals,
				BaseName:  anime.ID,
			},
			&outputs.ImageFile{
				Directory: originals,
				BaseName:  anime.ID,
				Format:    "webp",
				Quality:   85,
			},
		},
		InputProcessor:  ipo.SequentialInputs,
		OutputProcessor: ipo.SequentialOutputs,
	}

	err := system.Run()

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
