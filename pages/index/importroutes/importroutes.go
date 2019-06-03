package importroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/listimport"
	"github.com/animenotifier/notify.moe/pages/listimport/listimportanilist"
	"github.com/animenotifier/notify.moe/pages/listimport/listimportkitsu"
	"github.com/animenotifier/notify.moe/pages/listimport/listimportmyanimelist"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// Import
	page.Get(app, "/import", listimport.Get)
	page.Get(app, "/import/anilist/animelist", listimportanilist.Preview)
	page.Get(app, "/import/anilist/animelist/finish", listimportanilist.Finish)
	page.Get(app, "/import/myanimelist/animelist", listimportmyanimelist.Preview)
	page.Get(app, "/import/myanimelist/animelist/finish", listimportmyanimelist.Finish)
	page.Get(app, "/import/kitsu/animelist", listimportkitsu.Preview)
	page.Get(app, "/import/kitsu/animelist/finish", listimportkitsu.Finish)
}
