package middleware

import (
	"net"
	"time"

	"github.com/akyoto/cache"
)

var ipToHosts = cache.New(60 * time.Minute)

// GetHostsForIP returns all host names for the given IP (if cached).
func GetHostsForIP(ip string) ([]string, bool) {
	hosts, found := ipToHosts.Get(ip)

	if !found {
		hosts = findHostsForIP(ip)
	}

	if hosts == nil {
		return nil, found
	}

	return hosts.([]string), found
}

// Finds all host names for the given IP
func findHostsForIP(ip string) []string {
	hosts, err := net.LookupAddr(ip)

	if err != nil {
		return nil
	}

	if len(hosts) == 0 {
		return nil
	}

	// Cache host names
	ipToHosts.Set(ip, hosts, 60*time.Minute)

	return hosts
}
