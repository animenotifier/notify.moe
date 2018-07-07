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
	repeatSpaceCount := 8 - len(responseTimeString)

	if repeatSpaceCount < 0 {
		repeatSpaceCount = 0
	}

	responseTimeString = strings.Repeat(" ", repeatSpaceCount) + responseTimeString

	user := utils.GetUser(ctx)
	ip := ctx.RealIP()

	hostName := "<unknown host>"
	hostNames := GetHostsForIP(ip)

	if len(hostNames) != 0 {
		hostName = hostNames[0]
		hostName = strings.TrimSuffix(hostName, ".")
	}

	// Log every request
	id := "[id]"
	nick := "[guest]"

	if user != nil {
		id = user.ID
		nick = user.Nick
	}

	request.Info(nick, id, ip, hostName, responseTimeString, ctx.StatusCode, ctx.URI())

	// Log all requests that failed
	switch ctx.StatusCode {
	case http.StatusOK, http.StatusFound, http.StatusMovedPermanently, http.StatusPermanentRedirect, http.StatusTemporaryRedirect:
		// Ok.

	default:
		err.Error(nick, id, ip, hostName, responseTimeString, ctx.StatusCode, ctx.URI(), ctx.ErrorMessage)
	}

	// Notify us about long requests.
	// However ignore requests under /auth/ because those depend on 3rd party servers.
	if responseTime >= 300*time.Millisecond && !strings.HasPrefix(ctx.URI(), "/auth/") && !strings.HasPrefix(ctx.URI(), "/sitemap/") {
		err.Error("Long response time", nick, id, ip, hostName, responseTimeString, ctx.StatusCode, ctx.URI())
	}
}
