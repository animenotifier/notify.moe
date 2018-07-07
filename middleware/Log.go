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

var requestLog = log.New()
var errorLog = log.New()
var ipLog = log.New()

// Initialize log files
func init() {
	requestLog.AddOutput(log.File("logs/request.log"))
	errorLog.AddOutput(log.File("logs/error.log"))
	errorLog.AddOutput(os.Stderr)
	ipLog.AddOutput(log.File("logs/ip.log"))
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
	hostNames, cached := GetHostsForIP(ip)

	if !cached && len(hostNames) > 0 {
		ipLog.Info(ip, strings.Join(hostNames, ", "))
	}

	// Log every request
	id := "[id]"
	nick := "[guest]"

	if user != nil {
		id = user.ID
		nick = user.Nick
	}

	requestLog.Info(nick, id, ip, responseTimeString, ctx.StatusCode, ctx.URI())

	// Log all requests that failed
	switch ctx.StatusCode {
	case http.StatusOK, http.StatusFound, http.StatusMovedPermanently, http.StatusPermanentRedirect, http.StatusTemporaryRedirect:
		// Ok.

	default:
		errorLog.Error(nick, id, ip, responseTimeString, ctx.StatusCode, ctx.URI(), ctx.ErrorMessage)
	}

	// Notify us about long requests.
	// However ignore requests under /auth/ because those depend on 3rd party servers.
	if responseTime >= 300*time.Millisecond && !strings.HasPrefix(ctx.URI(), "/auth/") && !strings.HasPrefix(ctx.URI(), "/sitemap/") {
		errorLog.Error("Long response time", nick, id, ip, responseTimeString, ctx.StatusCode, ctx.URI())
	}
}
