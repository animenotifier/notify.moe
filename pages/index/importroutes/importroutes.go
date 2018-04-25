package importroutes

import (
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/pages/listimport"
	"github.com/animenotifier/notify.moe/pages/listimport/listimportanilist"
	"github.com/animenotifier/notify.moe/pages/listimport/listimportkitsu"
	"github.com/animenotifier/notify.moe/pages/listimport/listimportmyanimelist"
)

// Register registers the page routes.
func Register(l *layout.Layout) {
	// Import
	l.Page("/import", listimport.Get)
	l.Page("/import/anilist/animelist", listimportanilist.Preview)
	l.Page("/import/anilist/animelist/finish", listimportanilist.Finish)
	l.Page("/import/myanimelist/animelist", listimportmyanimelist.Preview)
	l.Page("/import/myanimelist/animelist/finish", listimportmyanimelist.Finish)
	l.Page("/import/kitsu/animelist", listimportkitsu.Preview)
	l.Page("/import/kitsu/animelist/finish", listimportkitsu.Finish)
}
