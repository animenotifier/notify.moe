package main

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aerogo/aero"
	"github.com/aerogo/log"
)

func init() {
	err := log.NewChannel("error")
	err.AddOutput(log.File("logs/error.log"))
	err.AddOutput(os.Stderr)

	web := log.NewChannel("web")
	web.AddOutput(log.File("logs/request.log"))

	app.Use(func(ctx *aero.Context, next func()) {
		start := time.Now()
		next()
		responseTime := time.Since(start)
		responseTimeString := strconv.Itoa(int(responseTime.Nanoseconds()/1000000)) + " ms"
		responseTimeString = strings.Repeat(" ", 8-len(responseTimeString)) + responseTimeString

		// Log every request
		web.Info(ctx.RealIP(), ctx.StatusCode, responseTimeString, ctx.URI())

		// Notify us about long requests
		if responseTime >= 100*time.Millisecond {
			err.Error("Unusually long response time", ctx.RealIP(), ctx.StatusCode, responseTimeString, ctx.URI())
		}
	})
}
