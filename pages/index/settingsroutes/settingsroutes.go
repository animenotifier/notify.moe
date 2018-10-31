package settingsroutes

import (
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/pages/settings"
)

// Register registers the page routes.
func Register(l *layout.Layout) {
	// Settings
	l.Page("/settings", settings.Get(components.SettingsPersonal))
	l.Page("/settings/accounts", settings.Get(components.SettingsAccounts))
	l.Page("/settings/notifications", settings.Get(components.SettingsNotifications))
	l.Page("/settings/apps", settings.Get(components.SettingsApps))
	l.Page("/settings/formatting", settings.Get(components.SettingsFormatting))
	l.Page("/settings/extras", settings.Get(components.SettingsExtras))
}
