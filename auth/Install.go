package auth

import (
	"os"

	"github.com/aerogo/aero"
	"github.com/aerogo/log"
)

const newUserStartRoute = "/welcome"

// Install installs all authentication routes in the application.
func Install(app *aero.Application) {
	authLog := log.New()
	authLog.AddWriter(os.Stdout)
	authLog.AddWriter(log.File("logs/auth.log"))

	// Login
	Google(app, authLog)
	Facebook(app, authLog)
	Twitter(app, authLog)

	// Logout
	Logout(app, authLog)
}
