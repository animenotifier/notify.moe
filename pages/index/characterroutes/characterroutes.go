package characterroutes

import (
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/pages/character"
	"github.com/animenotifier/notify.moe/pages/characters"
)

// Register registers the page routes.
func Register(l *layout.Layout) {
	// Characters
	l.Page("/characters", characters.Latest)
	l.Page("/characters/from/:index", characters.Latest)
	l.Page("/characters/best", characters.Best)
	l.Page("/characters/best/from/:index", characters.Best)

	// Character
	l.Page("/character/:id", character.Get)
	l.Page("/character/:id/edit", character.Edit)
	l.Page("/character/:id/edit/images", character.EditImages)
	l.Page("/character/:id/history", character.History)
}
