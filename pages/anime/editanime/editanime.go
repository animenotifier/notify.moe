package editanime

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/editform"
)

// Main anime edit page.
func Main(ctx aero.Context) error {
	id := ctx.Get("id")
	user := arn.GetUserFromContext(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not logged in or not auhorized to edit this anime")
	}

	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	return ctx.HTML(components.EditAnimeTabs(anime) + editform.Render(anime, "Edit anime", user))
}

// Images anime images edit page.
func Images(ctx aero.Context) error {
	id := ctx.Get("id")
	user := arn.GetUserFromContext(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not logged in or not auhorized to edit this anime")
	}

	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	return ctx.HTML(components.EditAnimeImages(anime))
}

// Characters anime characters edit page.
func Characters(ctx aero.Context) error {
	id := ctx.Get("id")
	user := arn.GetUserFromContext(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not logged in or not auhorized to edit")
	}

	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	animeCharacters, err := arn.GetAnimeCharacters(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime characters not found", err)
	}

	return ctx.HTML(components.EditAnimeTabs(anime) + editform.Render(animeCharacters, "Edit anime characters", user))
}

// Relations anime relations edit page.
func Relations(ctx aero.Context) error {
	id := ctx.Get("id")
	user := arn.GetUserFromContext(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not logged in or not auhorized to edit")
	}

	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	animeRelations, err := arn.GetAnimeRelations(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime relations not found", err)
	}

	return ctx.HTML(components.EditAnimeTabs(anime) + editform.Render(animeRelations, "Edit anime relations", user))
}

// Episodes anime episodes edit page.
func Episodes(ctx aero.Context) error {
	id := ctx.Get("id")
	user := arn.GetUserFromContext(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not logged in or not auhorized to edit")
	}

	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	// episodes := anime.Episodes()

	// if err != nil {
	// 	return ctx.Error(http.StatusNotFound, "Anime episodes not found", err)
	// }

	// editform.Render(episodes, "Edit anime episodes", user)
	return ctx.HTML(components.EditAnimeTabs(anime) + "<p class='no-data mountable'>Temporarily disabled.</p>")
}
