package main

import (
	"errors"
	"reflect"

	"github.com/aerogo/api"
	"github.com/animenotifier/arn"
)

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

	"/user/:nick/tracks": []string{
		"/+Akyoto/tracks",
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

	"/tracks/:id": []string{
		"/tracks/h0ac8sKkg",
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
	"/auth/google":              nil,
	"/auth/google/callback":     nil,
	"/auth/facebook":            nil,
	"/auth/facebook/callback":   nil,
	"/import":                   nil,
	"/import/anilist/animelist": nil,
	"/anime/:id/edit":           nil,
	"/new/thread":               nil,
	"/new/soundtrack":           nil,
	"/user":                     nil,
	"/settings":                 nil,
	"/extension/embed":          nil,
}

// API interfaces
var creatable = reflect.TypeOf((*api.Creatable)(nil)).Elem()
var updatable = reflect.TypeOf((*api.Updatable)(nil)).Elem()
var collection = reflect.TypeOf((*api.Collection)(nil)).Elem()

// Required interface implementations
var interfaceImplementations = map[string][]reflect.Type{
	"User": []reflect.Type{
		updatable,
	},
	"Thread": []reflect.Type{
		creatable,
	},
	"Post": []reflect.Type{
		creatable,
	},
	"SoundTrack": []reflect.Type{
		creatable,
	},
	"Analytics": []reflect.Type{
		creatable,
	},
	"AnimeList": []reflect.Type{
		collection,
	},
}

func init() {
	// Specify test routes
	for route, examples := range routeTests {
		app.Test(route, examples)
	}

	// Check interface implementations
	for typeName, interfaces := range interfaceImplementations {
		for _, requiredInterface := range interfaces {
			if !reflect.PtrTo(arn.DB.Type(typeName)).Implements(requiredInterface) {
				panic(errors.New(typeName + " does not implement interface " + requiredInterface.Name()))
			}
		}
	}
}
