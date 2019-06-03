package characterroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/character"
	"github.com/animenotifier/notify.moe/pages/characters"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// Characters
	page.Get(app, "/characters", characters.Latest)
	page.Get(app, "/characters/from/:index", characters.Latest)
	page.Get(app, "/characters/best", characters.Best)
	page.Get(app, "/characters/best/from/:index", characters.Best)

	// Character
	page.Get(app, "/character/:id", character.Get)
	page.Get(app, "/character/:id/edit", character.Edit)
	page.Get(app, "/character/:id/edit/images", character.EditImages)
	page.Get(app, "/character/:id/history", character.History)
}
