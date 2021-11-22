package routetests

var routeTests = map[string][]string{
	// User
	"/user/:nick": {
		"/+Akyoto",
	},

	"/user/:nick/characters/liked": {
		"/+Akyoto/characters/liked",
	},

	// "/user/:nick/forum/threads": {
	// 	"/+Akyoto/forum/threads",
	// },

	// "/user/:nick/forum/posts": {
	// 	"/+Akyoto/forum/posts",
	// },

	// "/user/:nick/soundtracks/added": {
	// 	"/+Akyoto/soundtracks/added",
	// },

	// "/user/:nick/soundtracks/added/from/:index": {
	// 	"/+Akyoto/soundtracks/added/from/3",
	// },

	// "/user/:nick/soundtracks/liked": {
	// 	"/+Akyoto/soundtracks/liked",
	// },

	// "/user/:nick/soundtracks/liked/from/:index": {
	// 	"/+Akyoto/soundtracks/liked/from/3",
	// },

	// "/user/:nick/quotes/added": {
	// 	"/+Scott/quotes/added",
	// },

	// "/user/:nick/quotes/added/from/:index": {
	// 	"/+Scott/quotes/added/from/3",
	// },

	// "/user/:nick/quotes/liked": {
	// 	"/+Scott/quotes/liked",
	// },

	// "/user/:nick/quotes/liked/from/:index": {
	// 	"/+Scott/quotes/liked/from/3",
	// },

	// "/user/:nick/followers": {
	// 	"/+Akyoto/followers",
	// },

	// "/user/:nick/stats": {
	// 	"/+Akyoto/stats",
	// },

	"/user/:nick/animelist/anime/:id": {
		"/+Akyoto/animelist/anime/74y2cFiiR",
	},

	"/user/:nick/animelist/watching": {
		"/+Akyoto/animelist/watching",
	},

	"/user/:nick/animelist/watching/from/:index": {
		"/+Akyoto/animelist/watching/from/1",
	},

	"/user/:nick/animelist/completed": {
		"/+Akyoto/animelist/completed",
	},

	"/user/:nick/animelist/completed/from/:index": {
		"/+Akyoto/animelist/completed/from/3",
	},

	"/user/:nick/animelist/planned": {
		"/+Akyoto/animelist/planned",
	},

	"/user/:nick/animelist/planned/from/:index": {
		"/+Akyoto/animelist/planned/from/3",
	},

	"/user/:nick/animelist/hold": {
		"/+Akyoto/animelist/hold",
	},

	"/user/:nick/animelist/hold/from/:index": {
		"/+Akyoto/animelist/hold/from/3",
	},

	"/user/:nick/animelist/dropped": {
		"/+Akyoto/animelist/dropped",
	},

	"/user/:nick/animelist/dropped/from/:index": {
		"/+Akyoto/animelist/dropped/from/3",
	},

	"/user/:nick/anime/recommended": {
		"/+Akyoto/anime/recommended",
	},

	"/user/:nick/anime/sequels": {
		"/+Akyoto/anime/sequels",
	},

	"/users/country/:country": {
		"/users/country/japan",
	},

	// Pages
	"/anime/:id": {
		"/anime/74y2cFiiR",
	},

	"/anime/:id/characters": {
		"/anime/74y2cFiiR/characters",
	},

	"/anime/:id/episodes": {
		"/anime/74y2cFiiR/episodes",
	},

	"/anime/:id/comments": {
		"/anime/74y2cFiiR/comments",
	},

	"/anime/:id/tracks": {
		"/anime/74y2cFiiR/tracks",
	},

	"/anime/:id/relations": {
		"/anime/74y2cFiiR/relations",
	},

	"/thread/:id": {
		"/thread/HJgS7c2K",
	},

	"/post/:id": {
		"/post/B1RzshnK",
	},

	"/forum/:tag": {
		"/forum/general",
	},

	"/genre/:name": {
		"/genre/action",
	},

	"/company/:id": {
		"/company/xCAUr7UkRaz",
	},

	"/company/:id/history": {
		"/company/xCAUr7UkRaz/history",
	},

	"/companies/from/:index": {
		"/companies/from/3",
	},

	"/explore/color/:color/anime": {
		"/explore/color/hsl:0.050,0.25,0.5/anime",
	},

	"/explore/color/:color/anime/from/:index": {
		"/explore/color/hsl:0.050,0.25,0.5/anime/from/3",
	},

	"/search/:term": {
		"/search/Dragon Ball",
	},

	"/quote/:id": {
		"/quote/gUZugd6zR",
	},

	"/quote/:id/edit": {
		"/quote/gUZugd6zR/edit",
	},

	"/quote/:id/history": {
		"/quote/gUZugd6zR/history",
	},

	"/quotes/from/:index": {
		"/quotes/from/2",
	},

	"/quotes/best/from/:index": {
		"/quotes/best/from/2",
	},

	"/soundtrack/:id": {
		"/soundtrack/h0ac8sKkg",
	},

	"/soundtrack/:id/lyrics": {
		"/soundtrack/vS64GbpzR/lyrics",
	},

	"/soundtrack/:id/edit": {
		"/soundtrack/h0ac8sKkg/edit",
	},

	"/soundtrack/:id/history": {
		"/soundtrack/h0ac8sKkg/history",
	},

	"/soundtracks": {
		"/soundtracks",
	},

	"/soundtracks/from/:index": {
		"/soundtracks/from/12",
	},

	"/soundtracks/best": {
		"/soundtracks/best",
	},

	"/soundtracks/best/from/:index": {
		"/soundtracks/best/from/12",
	},

	"/soundtracks/tag/:tag": {
		"/soundtracks/tag/moe",
	},

	"/soundtracks/tag/:tag/from/:index": {
		"/soundtracks/tag/moe/from/3",
	},

	"/character/:id": {
		"/character/dfrNQrmmg-",
	},

	"/episode/:id": {
		"/episode/oFlqThKWga",
	},

	"/compare/animelist/:nick-1/:nick-2": {
		"/compare/animelist/Akyoto/Scott",
	},

	"/explore/anime/:year/:season/:status/:type": {
		"/explore/anime/2011/any/finished/tv",
	},

	// AMV
	"/amv/:id": {
		"/amv/07scvSWmg",
	},

	"/amv/:id/edit": {
		"/amv/07scvSWmg/edit",
	},

	"/amv/:id/history": {
		"/amv/07scvSWmg/history",
	},

	// AMVs
	"/amvs/from/:index": {
		"/amvs/from/3",
	},

	"/amvs/best/from/:index": {
		"/amvs/best/from/3",
	},

	// Redirects
	"/mal/anime/:id": {
		"/mal/anime/33352",
	},

	"/kitsu/anime/:id": {
		"/kitsu/anime/12230",
	},

	"/anilist/anime/:id": {
		"/anilist/anime/21827",
	},

	// API
	"/api/anime/:id": {
		"/api/anime/74y2cFiiR",
	},

	"/api/thread/:id": {
		"/api/thread/HJgS7c2K",
	},

	"/api/post/:id": {
		"/api/post/B1RzshnK",
	},

	"/api/animelist/:id": {
		"/api/animelist/4J6qpK1ve",
	},

	"/api/settings/:id": {
		"/api/settings/4J6qpK1ve",
	},

	"/api/user/:id": {
		"/api/user/4J6qpK1ve",
	},

	"/api/googletouser/:id": {
		"/api/googletouser/106530160120373282283",
	},

	"/api/nicktouser/:id": {
		"/api/nicktouser/Akyoto",
	},

	"/api/soundtrack/:id": {
		"/api/soundtrack/h0ac8sKkg",
	},

	"/api/animecharacters/:id": {
		"/api/animecharacters/74y2cFiiR",
	},

	"/api/animerelations/:id": {
		"/api/animerelations/74y2cFiiR",
	},

	"/api/amv/:id": {
		"/api/amv/07scvSWmg",
	},

	"/api/character/:id": {
		"/api/character/dfrNQrmmg-",
	},

	"/api/company/:id": {
		"/api/company/xCAUr7UkRaz",
	},

	"/api/draftindex/:id": {
		"/api/draftindex/4J6qpK1ve",
	},

	"/api/inventory/:id": {
		"/api/inventory/4J6qpK1ve",
	},

	"/api/shopitem/:id": {
		"/api/shopitem/pro-account-3",
	},

	"/api/notification/:id": {
		"/api/notification/q6Y6eraig",
	},

	"/api/quote/:id": {
		"/api/quote/GXp675zmR",
	},

	"/api/usernotifications/:id": {
		"/api/usernotifications/4J6qpK1ve",
	},

	"/api/pushsubscriptions/:id": {
		"/api/pushsubscriptions/4J6qpK1ve",
	},

	// Images
	"/images/*file": {
		"/images/elements/no-avatar.svg",
	},

	// Extra tests for higher coverage
	"/_/+Akyoto": {
		"/_/+Akyoto",
	},

	"/_/search/dragon": {
		"/_/search/dragon",
	},
}

// All returns which specific routes to test for a given generic route.
func All() map[string][]string {
	return routeTests
}
