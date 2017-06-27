package main

var tests = map[string][]string{
	// User
	"/user/:nick": []string{
		"/+Akyoto",
	},

	"/user/:nick/threads": []string{
		"/+Akyoto/threads",
	},

	"/user/:nick/posts": []string{
		"/+Akyoto/posts",
	},

	"/user/:nick/animelist": []string{
		"/+Akyoto/animelist",
	},

	"/user/:nick/animelist/:id": []string{
		"/+Akyoto/animelist/7929",
	},

	// Pages
	"/anime/:id": []string{
		"/anime/1",
	},

	"/threads/:id": []string{
		"/threads/HJgS7c2K",
	},

	"/posts/:id": []string{
		"/posts/B1RzshnK",
	},

	"/forum/:tag": []string{
		"/forum/general",
	},

	"/search/:term": []string{
		"/search/Dragon Ball",
	},

	// API
	"/api/anime/:id": []string{
		"/api/anime/1",
	},

	"/api/thread/:id": []string{
		"/api/thread/HJgS7c2K",
	},

	"/api/post/:id": []string{
		"/api/post/B1RzshnK",
	},

	"/api/animelist/:id": []string{
		"/api/animelist/4J6qpK1ve",
	},

	"/api/animelist/:id/get/:item": []string{
		"/api/animelist/4J6qpK1ve/get/7929",
	},

	"/api/animelist/:id/get/:item/:property": []string{
		"/api/animelist/4J6qpK1ve/get/7929/Episodes",
	},

	"/api/settings/:id": []string{
		"/api/settings/4J6qpK1ve",
	},

	"/api/user/:id": []string{
		"/api/user/4J6qpK1ve",
	},

	"/api/emailtouser/:id": []string{
		"/api/emailtouser/e.urbach@gmail.com",
	},

	"/api/googletouser/:id": []string{
		"/api/googletouser/106530160120373282283",
	},

	"/api/nicktouser/:id": []string{
		"/api/nicktouser/Akyoto",
	},

	"/api/searchindex/:id": []string{
		"/api/searchindex/Anime",
	},

	// Images
	"/images/avatars/large/:file": []string{
		"/images/avatars/large/4J6qpK1ve.webp",
	},

	"/images/avatars/small/:file": []string{
		"/images/avatars/small/4J6qpK1ve.webp",
	},

	"/images/brand/:file": []string{
		"/images/brand/64.webp",
	},

	"/images/login/:file": []string{
		"/images/login/google",
	},

	"/images/cover/:file": []string{
		"/images/cover/default",
	},

	"/images/elements/:file": []string{
		"/images/elements/no-avatar.svg",
	},

	// Disable
	"/auth/google":          nil,
	"/auth/google/callback": nil,
	"/new/thread":           nil,
	"/user":                 nil,
	"/settings":             nil,
	"/extension/embed":      nil,
}

func init() {
	// Specify test routes
	for route, examples := range tests {
		app.Test(route, examples)
	}
}
