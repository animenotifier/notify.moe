package tokenapi

import (
	"os"

	"github.com/aerogo/aero"
	"github.com/aerogo/log"
)

// Install installs all authentication routes in the application.
func Install(app *aero.Application) {
	authLog := log.New()
	authLog.AddWriter(os.Stdout)
	authLog.AddWriter(log.File("logs/tokenapi.log"))

	// Login
	Main(app, authLog)
}
