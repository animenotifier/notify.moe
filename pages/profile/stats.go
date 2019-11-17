package profile

// import (
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/aerogo/aero"
// 	"github.com/animenotifier/notify.moe/arn"
// 	"github.com/animenotifier/notify.moe/components"
// 	"github.com/animenotifier/notify.moe/utils"
// )

// type stats map[string]float64

// // GetStatsByUser shows statistics for a given user.
// func GetStatsByUser(ctx aero.Context) error {
// 	nick := ctx.Get("nick")
// 	viewUser, err := arn.GetUserByNick(nick)
// 	userStats := utils.UserStats{}
// 	ratings := stats{}
// 	status := stats{}
// 	types := stats{}
// 	years := stats{}
// 	studios := stats{}
// 	genres := stats{}
// 	trackTags := stats{}

// 	if err != nil {
// 		return ctx.Error(http.StatusNotFound, "User not found", err)
// 	}

// 	animeList, err := arn.GetAnimeList(viewUser.ID)

// 	if err != nil {
// 		return ctx.Error(http.StatusInternalServerError, "Anime list not found", err)
// 	}

// 	animeList.Lock()
// 	defer animeList.Unlock()

// 	for _, item := range animeList.Items {
// 		status[item.Status]++

// 		if item.Status == arn.AnimeListStatusPlanned {
// 			continue
// 		}

// 		currentWatch := item.Episodes * item.Anime().EpisodeLength
// 		reWatch := item.RewatchCount * item.Anime().EpisodeCount * item.Anime().EpisodeLength
// 		duration := time.Duration(currentWatch + reWatch)
// 		userStats.AnimeWatchingTime += duration * time.Minute

// 		ratings[strconv.Itoa(int(item.Rating.Overall+0.5))]++
// 		types[item.Anime().Type]++

// 		for _, studio := range item.Anime().Studios() {
// 			studios[studio.Name.English]++
// 		}

// 		for _, genre := range item.Anime().Genres {
// 			genres[genre]++
// 		}

// 		if item.Anime().StartDate != "" {
// 			year := item.Anime().StartDate[:4]

// 			if year < "2000" {
// 				year = "Before 2000"
// 			}

// 			years[year]++
// 		}
// 	}

// 	for track := range arn.StreamSoundTracks() {
// 		if !track.LikedBy(viewUser.ID) {
// 			continue
// 		}

// 		for _, tag := range track.Tags {
// 			if strings.Contains(tag, ":") {
// 				continue
// 			}

// 			trackTags[tag]++
// 		}
// 	}

// 	userStats.PieCharts = []*arn.PieChart{
// 		arn.NewPieChart("Genres", genres),
// 		arn.NewPieChart("Studios", studios),
// 		arn.NewPieChart("Years", years),
// 		arn.NewPieChart("Ratings", ratings),
// 		arn.NewPieChart("Types", types),
// 		arn.NewPieChart("Status", status),
// 		arn.NewPieChart("Soundtracks", trackTags),
// 	}

// 	return ctx.HTML(components.ProfileStats(&userStats, viewUser, arn.GetUserFromContext(ctx), ctx.Path()))
// }
