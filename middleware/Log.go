package middleware

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aerogo/aero"
	"github.com/aerogo/log"
	"github.com/animenotifier/notify.moe/utils"
)

var request = log.New()
var err = log.New()

// Initialize log files
func init() {
	request.AddOutput(log.File("logs/request.log"))

	err.AddOutput(log.File("logs/error.log"))
	err.AddOutput(os.Stderr)
}

// Log middleware logs every request into logs/request.log and errors into logs/error.log.
func Log() aero.Middleware {
	return func(ctx *aero.Context, next func()) {
		start := time.Now()
		next()
		responseTime := time.Since(start)

		go logRequest(ctx, responseTime)
	}
}

// Logs a single request
func logRequest(ctx *aero.Context, responseTime time.Duration) {
	responseTimeString := strconv.Itoa(int(responseTime.Nanoseconds()/1000000)) + " ms"
	responseTimeString = strings.Repeat(" ", 8-len(responseTimeString)) + responseTimeString

	user := utils.GetUser(ctx)
	ip := ctx.RealIP()

	hostName := "<unknown host>"
	hostNames := GetHostsForIP(ip)

	if len(hostNames) != 0 {
		hostName = hostNames[0]
		hostName = strings.TrimSuffix(hostName, ".")
	}

	// Log every request
	if user != nil {
		request.Info(user.Nick, ip, hostName, responseTimeString, ctx.StatusCode, ctx.URI())
	} else {
		request.Info("[guest]", ip, hostName, responseTimeString, ctx.StatusCode, ctx.URI())
	}

	// Log all requests that failed
	switch ctx.StatusCode {
	case http.StatusOK, http.StatusFound, http.StatusMovedPermanently, http.StatusPermanentRedirect, http.StatusTemporaryRedirect:
		// Ok.

	default:
		err.Error(http.StatusText(ctx.StatusCode), ip, hostName, responseTimeString, ctx.StatusCode, ctx.URI())
	}

	// Notify us about long requests.
	// However ignore requests under /auth/ because those depend on 3rd party servers.
	if responseTime >= 200*time.Millisecond && !strings.HasPrefix(ctx.URI(), "/auth/") {
		err.Error("Long response time", ip, hostName, responseTimeString, ctx.StatusCode, ctx.URI())
	}
}
