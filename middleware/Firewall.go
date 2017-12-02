package middleware

// const requestThreshold = 10

// var ipToStats = cache.New(15*time.Minute, 15*time.Minute)

// // IPStats captures the statistics for a single IP.
// type IPStats struct {
// 	Requests []string
// }

// // Firewall middleware detects malicious requests.
// func Firewall() aero.Middleware {
// 	return func(ctx *aero.Context, next func()) {
// 		var stats *IPStats

// 		ip := ctx.RealIP()

// 		// Allow localhost
// 		if ip == "127.0.0.1" {
// 			next()
// 			return
// 		}

// 		statsObj, found := ipToStats.Get(ip)

// 		if found {
// 			stats = statsObj.(*IPStats)
// 		} else {
// 			stats = &IPStats{
// 				Requests: []string{},
// 			}

// 			ipToStats.Set(ip, stats, cache.DefaultExpiration)
// 		}

// 		// Add requested URI to the list of requests
// 		stats.Requests = append(stats.Requests, ctx.URI())

// 		if len(stats.Requests) > requestThreshold {
// 			stats.Requests = stats.Requests[len(stats.Requests)-requestThreshold:]

// 			for _, uri := range stats.Requests {
// 				// Allow request
// 				if strings.Contains(uri, "/_/") || strings.Contains(uri, "/api/") || strings.Contains(uri, "/scripts") || strings.Contains(uri, "/service-worker") || strings.Contains(uri, "/images/") || strings.Contains(uri, "/favicon.ico") || strings.Contains(uri, "/extension/embed") {
// 					next()
// 					return
// 				}
// 			}

// 			// Allow logged in users
// 			if ctx.HasSession() {
// 				user := utils.GetUser(ctx)

// 				if user != nil {
// 					// Allow request
// 					next()
// 					return
// 				}
// 			}

// 			// Disallow request
// 			request.Error("[guest]", ip, "BLOCKED BY FIREWALL", ctx.URI())
// 			return
// 		}

// 		// Allow the request if the number of requests done by the IP is below the threshold
// 		next()
// 	}
// }
