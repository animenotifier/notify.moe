package main

func init() {
	// Uncomment these if you want to enable live profiling via /debug/pprof

	// app.Router.HandlerFunc("GET", "/debug/pprof/", http.HandlerFunc(pprof.Index))
	// app.Router.HandlerFunc("GET", "/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	// app.Router.HandlerFunc("GET", "/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	// app.Router.HandlerFunc("GET", "/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	// app.Router.HandlerFunc("GET", "/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

	// app.Router.Handler("GET", "/debug/pprof/goroutine", pprof.Handler("goroutine"))
	// app.Router.Handler("GET", "/debug/pprof/heap", pprof.Handler("heap"))
	// app.Router.Handler("GET", "/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	// app.Router.Handler("GET", "/debug/pprof/block", pprof.Handler("block"))
}
