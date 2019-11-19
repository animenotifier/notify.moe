package anime

import (
	"net/http"

	"github.com/animenotifier/notify.moe/components"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Episodes ...
func Episodes(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	id := ctx.Get("id")
	anime, err := arn.GetAnime(id)
	episodeToFriends := map[int][]*arn.User{}

	if user != nil {
		ownListItem := user.AnimeList().Find(anime.ID)

		if ownListItem != nil {
			episodeToFriends[ownListItem.Episodes] = append(episodeToFriends[ownListItem.Episodes], user)
		}

		for _, friend := range user.Follows() {
			friendAnimeList := friend.AnimeList()
			friendAnimeListItem := friendAnimeList.Find(anime.ID)

			if friendAnimeListItem != nil && !friendAnimeListItem.Private && len(episodeToFriends[friendAnimeListItem.Episodes]) < maxFriendsPerEpisode {
				episodeToFriends[friendAnimeListItem.Episodes] = append(episodeToFriends[friendAnimeListItem.Episodes], friend)
			}
		}
	}

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	return ctx.HTML(components.AnimeEpisodes(anime, anime.Episodes(), episodeToFriends, user, true))
}
