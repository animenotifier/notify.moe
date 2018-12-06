package routetests

var routeTests = map[string][]string{
	// User
	"/user/:nick": []string{
		"/+Akyoto",
	},

	"/user/:nick/characters/liked": []string{
		"/+Akyoto/characters/liked",
	},

	"/user/:nick/forum/threads": []string{
		"/+Akyoto/forum/threads",
	},

	"/user/:nick/forum/posts": []string{
		"/+Akyoto/forum/posts",
	},

	"/user/:nick/soundtracks/added": []string{
		"/+Akyoto/soundtracks/added",
	},

	"/user/:nick/soundtracks/added/from/:index": []string{
		"/+Akyoto/soundtracks/added/from/3",
	},

	"/user/:nick/soundtracks/liked": []string{
		"/+Akyoto/soundtracks/liked",
	},

	"/user/:nick/soundtracks/liked/from/:index": []string{
		"/+Akyoto/soundtracks/liked/from/3",
	},

	"/user/:nick/quotes/added": []string{
		"/+Scott/quotes/added",
	},

	"/user/:nick/quotes/added/from/:index": []string{
		"/+Scott/quotes/added/from/3",
	},

	"/user/:nick/quotes/liked": []string{
		"/+Scott/quotes/liked",
	},

	"/user/:nick/quotes/liked/from/:index": []string{
		"/+Scott/quotes/liked/from/3",
	},

	"/user/:nick/followers": []string{
		"/+Akyoto/followers",
	},

	"/user/:nick/stats": []string{
		"/+Akyoto/stats",
	},

	"/user/:nick/animelist/anime/:id": []string{
		"/+Akyoto/animelist/anime/74y2cFiiR",
	},

	"/user/:nick/animelist/watching": []string{
		"/+Akyoto/animelist/watching",
	},

	"/user/:nick/animelist/watching/from/:index": []string{
		"/+Akyoto/animelist/watching/from/1",
	},

	"/user/:nick/animelist/completed": []string{
		"/+Akyoto/animelist/completed",
	},

	"/user/:nick/animelist/completed/from/:index": []string{
		"/+Akyoto/animelist/completed/from/3",
	},

	"/user/:nick/animelist/planned": []string{
		"/+Akyoto/animelist/planned",
	},

	"/user/:nick/animelist/planned/from/:index": []string{
		"/+Akyoto/animelist/planned/from/3",
	},

	"/user/:nick/animelist/hold": []string{
		"/+Akyoto/animelist/hold",
	},

	"/user/:nick/animelist/hold/from/:index": []string{
		"/+Akyoto/animelist/hold/from/3",
	},

	"/user/:nick/animelist/dropped": []string{
		"/+Akyoto/animelist/dropped",
	},

	"/user/:nick/animelist/dropped/from/:index": []string{
		"/+Akyoto/animelist/dropped/from/3",
	},

	"/user/:nick/recommended/anime": []string{
		"/+Akyoto/recommended/anime",
	},

	"/users/country/:country": []string{
		"/users/country/japan",
	},

	// Pages
	"/anime/:id": []string{
		"/anime/74y2cFiiR",
	},

	"/anime/:id/characters": []string{
		"/anime/74y2cFiiR/characters",
	},

	"/anime/:id/episodes": []string{
		"/anime/74y2cFiiR/episodes",
	},

	"/anime/:id/comments": []string{
		"/anime/74y2cFiiR/comments",
	},

	"/anime/:id/tracks": []string{
		"/anime/74y2cFiiR/tracks",
	},

	"/anime/:id/relations": []string{
		"/anime/74y2cFiiR/relations",
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

	"/genre/:name": []string{
		"/genre/action",
	},

	"/company/:id": []string{
		"/company/xCAUr7UkRaz",
	},

	"/company/:id/history": []string{
		"/company/xCAUr7UkRaz/history",
	},

	"/companies/from/:index": []string{
		"/companies/from/3",
	},

	"/explore/color/:color/anime": []string{
		"/explore/color/hsl:0.050,0.25,0.5/anime",
	},

	"/explore/color/:color/anime/from/:index": []string{
		"/explore/color/hsl:0.050,0.25,0.5/anime/from/3",
	},

	"/search/:term": []string{
		"/search/Dragon Ball",
	},

	"/quote/:id": []string{
		"/quote/gUZugd6zR",
	},

	"/quote/:id/edit": []string{
		"/quote/gUZugd6zR/edit",
	},

	"/quote/:id/history": []string{
		"/quote/gUZugd6zR/history",
	},

	"/quotes/from/:index": []string{
		"/quotes/from/2",
	},

	"/quotes/best/from/:index": []string{
		"/quotes/best/from/2",
	},

	"/soundtrack/:id": []string{
		"/soundtrack/h0ac8sKkg",
	},

	"/soundtrack/:id/lyrics": []string{
		"/soundtrack/vS64GbpzR/lyrics",
	},

	"/soundtrack/:id/edit": []string{
		"/soundtrack/h0ac8sKkg/edit",
	},

	"/soundtrack/:id/history": []string{
		"/soundtrack/h0ac8sKkg/history",
	},

	"/soundtracks": []string{
		"/soundtracks",
	},

	"/soundtracks/from/:index": []string{
		"/soundtracks/from/12",
	},

	"/soundtracks/best": []string{
		"/soundtracks/best",
	},

	"/soundtracks/best/from/:index": []string{
		"/soundtracks/best/from/12",
	},

	"/soundtracks/tag/:tag": []string{
		"/soundtracks/tag/moe",
	},

	"/soundtracks/tag/:tag/from/:index": []string{
		"/soundtracks/tag/moe/from/3",
	},

	"/character/:id": []string{
		"/character/dfrNQrmmg-",
	},

	// "/kitsu/character/:id": []string{
	// 	"/kitsu/character/6556",
	// },

	// "/mal/character/:id": []string{
	// 	"/mal/character/498",
	// },

	"/compare/animelist/:nick-1/:nick-2": []string{
		"/compare/animelist/Akyoto/Scott",
	},

	"/explore/anime/:year/:season/:status/:type": []string{
		"/explore/anime/2011/any/finished/tv",
	},

	// AMV
	"/amv/:id": []string{
		"/amv/07scvSWmg",
	},

	"/amv/:id/edit": []string{
		"/amv/07scvSWmg/edit",
	},

	"/amv/:id/history": []string{
		"/amv/07scvSWmg/history",
	},

	// AMVs
	"/amvs/from/:index": []string{
		"/amvs/from/3",
	},

	"/amvs/best/from/:index": []string{
		"/amvs/best/from/3",
	},

	// Redirects
	"/mal/anime/:id": []string{
		"/mal/anime/33352",
	},

	"/kitsu/anime/:id": []string{
		"/kitsu/anime/12230",
	},

	"/anilist/anime/:id": []string{
		"/anilist/anime/21827",
	},

	// API
	"/api/anime/:id": []string{
		"/api/anime/74y2cFiiR",
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

	"/api/settings/:id": []string{
		"/api/settings/4J6qpK1ve",
	},

	"/api/user/:id": []string{
		"/api/user/4J6qpK1ve",
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

	"/api/analytics/:id": []string{
		"/api/analytics/4J6qpK1ve",
	},

	"/api/soundtrack/:id": []string{
		"/api/soundtrack/h0ac8sKkg",
	},

	"/api/userfollows/:id": []string{
		"/api/userfollows/4J6qpK1ve",
	},

	"/api/animecharacters/:id": []string{
		"/api/animecharacters/74y2cFiiR",
	},

	"/api/animerelations/:id": []string{
		"/api/animerelations/74y2cFiiR",
	},

	"/api/animeepisodes/:id": []string{
		"/api/animeepisodes/74y2cFiiR",
	},

	"/anime/:id/episode/:episode-number": []string{
		"/anime/74y2cFiiR/episode/5",
	},

	"/api/amv/:id": []string{
		"/api/amv/07scvSWmg",
	},

	"/api/character/:id": []string{
		"/api/character/dfrNQrmmg-",
	},

	"/api/company/:id": []string{
		"/api/company/xCAUr7UkRaz",
	},

	"/api/draftindex/:id": []string{
		"/api/draftindex/4J6qpK1ve",
	},

	"/api/inventory/:id": []string{
		"/api/inventory/4J6qpK1ve",
	},

	"/api/shopitem/:id": []string{
		"/api/shopitem/pro-account-3",
	},

	"/api/notification/:id": []string{
		"/api/notification/q6Y6eraig",
	},

	"/api/quote/:id": []string{
		"/api/quote/GXp675zmR",
	},

	"/api/usernotifications/:id": []string{
		"/api/usernotifications/4J6qpK1ve",
	},

	"/api/pushsubscriptions/:id": []string{
		"/api/pushsubscriptions/4J6qpK1ve",
	},

	// Images
	"/images/*file": []string{
		"/images/elements/no-avatar.svg",
	},

	// Extra tests for higher coverage
	"/_/+Akyoto": []string{
		"/_/+Akyoto",
	},

	"/_/search/dragon": []string{
		"/_/search/dragon",
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
	"/animelist/watching":                            nil,
	"/animelist/completed":                           nil,
	"/animelist/planned":                             nil,
	"/animelist/hold":                                nil,
	"/animelist/dropped":                             nil,
	"/notifications":                                 nil,
	"/user/:nick/notifications":                      nil,
	"/user/:nick/edit":                               nil,
	"/user/:nick/log":                                nil,
	"/user/:nick/log/from/:index":                    nil,
	"/editor/soundtracks/file":                       nil,
	"/editor/soundtracks/links":                      nil,
	"/editor/soundtracks/lyrics/missing":             nil,
	"/editor/soundtracks/lyrics/unaligned":           nil,
	"/editor/soundtracks/tags":                       nil,
	"/api/test/notification":                         nil,
	"/api/paypal/payment/create":                     nil,
	"/api/emailtouser/:id":                           nil,
	"/api/userfollows/:id/get/:item":                 nil,
	"/api/userfollows/:id/get/:item/:property":       nil,
	"/api/pushsubscriptions/:id/get/:item":           nil,
	"/api/pushsubscriptions/:id/get/:item/:property": nil,
	"/api/count/notifications/unseen":                nil,
	"/api/mark/notifications/seen":                   nil,
	"/api/sse/events":                                nil,
	"/editor/kitsu/new/anime":                        nil,
	"/paypal/success":                                nil,
	"/paypal/cancel":                                 nil,
	"/anime/:id/edit":                                nil,
	"/anime/:id/edit/images":                         nil,
	"/anime/:id/edit/characters":                     nil,
	"/anime/:id/edit/relations":                      nil,
	"/anime/:id/edit/episodes":                       nil,
	"/anime/:id/edit/history":                        nil,
	"/new/thread":                                    nil,
	"/thread/:id/edit":                               nil,
	"/post/:id/edit":                                 nil,
	"/company/:id/edit":                              nil,
	"/admin/purchases":                               nil,
	"/admin/registrations":                           nil,
	"/admin/payments":                                nil,
	"/editor/anilist":                                nil,
	"/editor/shoboi":                                 nil,
	"/dark-flame-master":                             nil,
	"/groups/joined":                                 nil,
	"/user":                                          nil,
	"/settings":                                      nil,
	"/settings/accounts":                             nil,
	"/settings/notifications":                        nil,
	"/settings/info":                                 nil,
	"/settings/avatar":                               nil,
	"/settings/style":                                nil,
	"/settings/extras":                               nil,
	"/shop":                                          nil,
	"/shop/history":                                  nil,
	"/support":                                       nil,
	"/charge":                                        nil,
	"/log":                                           nil,
	"/log/from/:index":                               nil,
	"/inventory":                                     nil,
	"/extension/embed":                               nil,
	"/welcome":                                       nil,
}

// All returns which specific routes to test for a given generic route.
func All() map[string][]string {
	return routeTests
}
