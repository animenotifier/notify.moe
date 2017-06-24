package auth

import (
	"os"

	"github.com/aerogo/log"
)

var authLog = log.New()

func init() {
	authLog.AddOutput(os.Stdout)
	authLog.AddOutput(log.File("logs/auth.log"))
}
