package middleware

import (
	"net"
	"time"

	"github.com/aerogo/aero"
	cache "github.com/patrickmn/go-cache"
)

var ipToHosts = cache.New(60*time.Minute, 30*time.Minute)

// IPToHost middleware tries to find domain names for the given IP.
func IPToHost() aero.Middleware {
	return func(ctx *aero.Context, next func()) {
		next()

		go findHostsForIP(ctx.RealIP())
	}
}

// GetHostsForIP returns all host names for the given IP (if cached).
func GetHostsForIP(ip string) []string {
	hosts, found := ipToHosts.Get(ip)

	if !found || hosts == nil {
		return nil
	}

	return hosts.([]string)
}

// Finds all host names for the given IP
func findHostsForIP(ip string) {
	hosts, err := net.LookupAddr(ip)

	if err != nil {
		return
	}

	if len(hosts) == 0 {
		return
	}

	// Cache host names
	ipToHosts.Set(ip, hosts, cache.DefaultExpiration)
}
