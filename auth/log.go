package auth

import (
	"os"

	"github.com/aerogo/log"
)

var authLog = log.New()

func init() {
	authLog.AddWriter(os.Stdout)
	authLog.AddWriter(log.File("logs/auth.log"))
}
