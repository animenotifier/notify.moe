package anime

import (
	"net/http"

	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Episodes ...
func Episodes(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	id := ctx.Get("id")
	anime, err := arn.GetAnime(id)
	episodeToFriends := map[int][]*arn.User{}

	if user != nil {
		for _, friend := range user.Follows().Users() {
			friendAnimeList := friend.AnimeList()
			friendAnimeListItem := friendAnimeList.Find(anime.ID)

			if friendAnimeListItem != nil && !friendAnimeListItem.Private && len(episodeToFriends[friendAnimeListItem.Episodes]) < maxFriendsPerEpisode {
				episodeToFriends[friendAnimeListItem.Episodes] = append(episodeToFriends[friendAnimeListItem.Episodes], friend)
			}
		}

		ownListItem := user.AnimeList().Find(anime.ID)

		if ownListItem != nil {
			episodeToFriends[ownListItem.Episodes] = append(episodeToFriends[ownListItem.Episodes], user)
		}
	}

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	return ctx.HTML(components.AnimeEpisodes(anime, anime.Episodes().Items, episodeToFriends, user, true))
}
