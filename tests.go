package main

var routeTests = map[string][]string{
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

	"/user/:nick/soundtracks": []string{
		"/+Akyoto/soundtracks",
	},

	"/user/:nick/followers": []string{
		"/+Akyoto/followers",
	},

	"/user/:nick/stats": []string{
		"/+Akyoto/stats",
	},

	"/user/:nick/animelist": []string{
		"/+Akyoto/animelist",
	},

	"/user/:nick/animelist/anime/:id": []string{
		"/+Akyoto/animelist/anime/7929",
	},

	"/user/:nick/animelist/watching": []string{
		"/+Akyoto/animelist/watching",
	},

	"/user/:nick/animelist/completed": []string{
		"/+Akyoto/animelist/completed",
	},

	"/user/:nick/animelist/planned": []string{
		"/+Akyoto/animelist/planned",
	},

	"/user/:nick/animelist/hold": []string{
		"/+Akyoto/animelist/hold",
	},

	"/user/:nick/animelist/dropped": []string{
		"/+Akyoto/animelist/dropped",
	},

	// Pages
	"/anime/:id": []string{
		"/anime/1",
	},

	"/thread/:id": []string{
		"/thread/HJgS7c2K",
	},

	"/post/:id": []string{
		"/post/B1RzshnK",
	},

	"/forum/:tag": []string{
		"/forum/general",
	},

	"/search/:term": []string{
		"/search/Dragon Ball",
	},

	"/soundtrack/:id": []string{
		"/soundtrack/h0ac8sKkg",
	},

	"/character/:id": []string{
		"/character/6556",
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

	"/api/facebooktouser/:id": []string{
		"/api/facebooktouser/10207576239700188",
	},

	"/api/nicktouser/:id": []string{
		"/api/nicktouser/Akyoto",
	},

	"/api/searchindex/:id": []string{
		"/api/searchindex/Anime",
	},

	"/api/analytics/:id": []string{
		"/api/analytics/4J6qpK1ve",
	},

	"/api/soundtrack/:id": []string{
		"/api/soundtrack/h0ac8sKkg",
	},

	"/api/soundcloudtosoundtrack/:id": []string{
		"/api/soundcloudtosoundtrack/145918628",
	},

	"/api/youtubetosoundtrack/:id": []string{
		"/api/youtubetosoundtrack/hU2wqJuOIp4",
	},

	"/api/userfollows/:id": []string{
		"/api/userfollows/4J6qpK1ve",
	},

	"/api/anilisttoanime/:id": []string{
		"/api/anilisttoanime/527",
	},

	"/api/animecharacters/:id": []string{
		"/api/animecharacters/323",
	},

	"/api/animeepisodes/:id": []string{
		"/api/animeepisodes/323",
	},

	"/api/character/:id": []string{
		"/api/character/6556",
	},

	"/api/pushsubscriptions/:id": []string{
		"/api/pushsubscriptions/4J6qpK1ve",
	},

	"/api/myanimelisttoanime/:id": []string{
		"/api/myanimelisttoanime/527",
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

	// Disable these tests because they require authorization
	"/auth/google":                                   nil,
	"/auth/google/callback":                          nil,
	"/auth/facebook":                                 nil,
	"/auth/facebook/callback":                        nil,
	"/dashboard":                                     nil,
	"/import":                                        nil,
	"/import/anilist/animelist":                      nil,
	"/import/anilist/animelist/finish":               nil,
	"/import/myanimelist/animelist":                  nil,
	"/import/myanimelist/animelist/finish":           nil,
	"/import/kitsu/animelist":                        nil,
	"/import/kitsu/animelist/finish":                 nil,
	"/api/test/notification":                         nil,
	"/api/paypal/payment/create":                     nil,
	"/api/userfollows/:id/get/:item":                 nil,
	"/api/userfollows/:id/get/:item/:property":       nil,
	"/api/pushsubscriptions/:id/get/:item":           nil,
	"/api/pushsubscriptions/:id/get/:item/:property": nil,
	"/paypal/success":                                nil,
	"/paypal/cancel":                                 nil,
	"/anime/:id/edit":                                nil,
	"/new/thread":                                    nil,
	"/new/soundtrack":                                nil,
	"/admin/anilist":                                 nil,
	"/admin/shoboi":                                  nil,
	"/user":                                          nil,
	"/settings":                                      nil,
	"/extension/embed":                               nil,
}
