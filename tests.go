package main

func init() {
	// User
	app.Test("/user/:nick", []string{
		"/+Akyoto",
	})

	app.Test("/user/:nick/threads", []string{
		"/+Akyoto/threads",
	})

	app.Test("/user/:nick/avatar", []string{
		"/+Akyoto/avatar",
	})

	app.Test("/user/:nick/avatar/small", []string{
		"/+Akyoto/avatar/small",
	})

	// Pages
	app.Test("/anime/:id", []string{
		"/anime/1",
	})

	app.Test("/threads/:id", []string{
		"/threads/HJgS7c2K",
	})

	app.Test("/posts/:id", []string{
		"/posts/B1RzshnK",
	})

	app.Test("/forum/:tag", []string{
		"/forum/general",
	})

	// API
	app.Test("/api/anime/:id", []string{
		"/api/anime/1",
	})

	app.Test("/api/thread/:id", []string{
		"/api/thread/HJgS7c2K",
	})

	app.Test("/api/post/:id", []string{
		"/api/post/B1RzshnK",
	})

	app.Test("/api/animelist/:id", []string{
		"/api/animelist/4J6qpK1ve",
	})

	app.Test("/api/settings/:id", []string{
		"/api/settings/4J6qpK1ve",
	})

	app.Test("/api/user/:id", []string{
		"/api/user/4J6qpK1ve",
	})

	// Others
	app.Test("/images/cover/:file", []string{
		"/images/cover/default",
	})

	app.Test("/images/elements/:file", []string{
		"/images/elements/no-gravatar.svg",
	})

	// Disable
	app.Test("/auth/google", nil)
	app.Test("/auth/google/callback", nil)
}
