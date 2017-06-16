package middleware

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aerogo/aero"
	"github.com/aerogo/log"
)

// Log middleware logs every request into logs/request.log and errors into logs/error.log.
func Log() aero.Middleware {
	request := log.NewLog()
	request.AddOutput(log.File("logs/request.log"))

	err := log.NewLog()
	err.AddOutput(log.File("logs/error.log"))
	err.AddOutput(os.Stderr)

	return func(ctx *aero.Context, next func()) {
		start := time.Now()
		next()
		responseTime := time.Since(start)
		responseTimeString := strconv.Itoa(int(responseTime.Nanoseconds()/1000000)) + " ms"
		responseTimeString = strings.Repeat(" ", 8-len(responseTimeString)) + responseTimeString

		// Log every request
		request.Info(ctx.RealIP(), ctx.StatusCode, responseTimeString, ctx.URI())

		// Log all requests that failed
		switch ctx.StatusCode {
		case http.StatusOK, http.StatusFound, http.StatusMovedPermanently, http.StatusPermanentRedirect, http.StatusTemporaryRedirect:
			// Ok.

		default:
			err.Error(http.StatusText(ctx.StatusCode), ctx.RealIP(), ctx.StatusCode, responseTimeString, ctx.URI())
		}

		// Notify us about long requests
		if responseTime >= 200*time.Millisecond {
			err.Error("Long response time", ctx.RealIP(), ctx.StatusCode, responseTimeString, ctx.URI())
		}
	}
}
