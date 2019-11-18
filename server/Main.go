package server

// Main runs the main loop of the web server.
func Main() {
	app := New()
	app.Run()
}
