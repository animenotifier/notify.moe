package main

func main() {

}

// import (
// 	"github.com/animenotifier/arn"
// 	"github.com/fatih/color"
// )

// func main() {
// 	// Get a stream of all anime
// 	allAnime, err := arn.AllAnime()

// 	if err != nil {
// 		panic(err)
// 	}

// 	// Iterate over the stream
// 	for _, anime := range allAnime {
// 		for _, trailer := range anime.Trailers {
// 			// trailer.ServiceID = trailer.DeprecatedVideoID
// 			println(trailer.DeprecatedVideoID)
// 			trailer.ServiceID = trailer.DeprecatedVideoID
// 		}

// 		if anime.Trailers == nil {
// 			anime.Trailers = []*arn.ExternalMedia{}
// 		}

// 		err := anime.Save()

// 		if err != nil {
// 			color.Red("Error saving anime: %v", err)
// 		}
// 	}
// }
