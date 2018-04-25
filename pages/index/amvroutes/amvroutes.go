package amvroutes

import (
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/pages/amv"
	"github.com/animenotifier/notify.moe/pages/amvs"
)

// Register registers the page routes.
func Register(l *layout.Layout) {
	// AMV
	l.Page("/amv/:id", amv.Get)
	l.Page("/amv/:id/edit", amv.Edit)
	l.Page("/amv/:id/history", amv.History)

	// AMVs
	l.Page("/amvs", amvs.Latest)
	l.Page("/amvs/from/:index", amvs.Latest)
	l.Page("/amvs/best", amvs.Best)
	l.Page("/amvs/best/from/:index", amvs.Best)
}
