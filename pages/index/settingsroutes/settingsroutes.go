package settingsroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/pages/settings"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// Settings
	page.Get(app, "/settings", settings.Get(components.SettingsPersonal))
	page.Get(app, "/settings/accounts", settings.Get(components.SettingsAccounts))
	page.Get(app, "/settings/notifications", settings.Get(components.SettingsNotifications))
	page.Get(app, "/settings/info", settings.Get(components.SettingsInfo))
	page.Get(app, "/settings/style", settings.Get(components.SettingsStyle))
	page.Get(app, "/settings/extras", settings.Get(components.SettingsExtras))
}
