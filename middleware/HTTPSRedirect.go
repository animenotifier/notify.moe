package middleware

// // HTTPSRedirect middleware redirects to HTTPS if needed.
// func HTTPSRedirect() aero.Middleware {
// 	return func(ctx *aero.Context, next func()) {
// 		request := ctx.Request()
// 		userAgent := request.Header().Get("User-Agent")
// 		isBrowser := strings.Contains(userAgent, "Mozilla/") || strings.Contains(userAgent, "Chrome/") || strings.Contains(userAgent, "AppleWebKit/")

// 		if !strings.HasPrefix(request.Protocol(), "HTTP/2") && isBrowser {
// 			fmt.Println("Redirect to HTTPS")
// 			ctx.Redirect("https://" + request.Host() + request.URL().Path)
// 			ctx.Response().WriteHeader(ctx.StatusCode)
// 			return
// 		}

// 		// Handle the request
// 		next()
// 	}
// }
