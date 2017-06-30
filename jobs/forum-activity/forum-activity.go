package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

const maxEntries = 5

func main() {
	color.Yellow("Caching list of forum activities")

	posts, err := arn.AllPosts()
	arn.PanicOnError(err)

	threads, err := arn.AllThreads()
	arn.PanicOnError(err)

	arn.SortPostsLatestFirst(posts)
	arn.SortThreadsLatestFirst(threads)

	posts = arn.FilterPostsWithUniqueThreads(posts, maxEntries)

	postPostables := arn.ToPostables(posts)
	threadPostables := arn.ToPostables(threads)

	allPostables := append(postPostables, threadPostables...)

	arn.SortPostablesLatestFirst(allPostables)
	cachedPostables := arn.FilterPostablesWithUniqueThreads(allPostables, maxEntries)

	cache := &arn.ListOfMappedIDs{}

	for _, postable := range cachedPostables {
		cache.Append(postable.Type(), postable.ID())
	}

	// // Debug log
	// arn.PrettyPrint(cache)

	// // Try to resolve
	// for _, r := range arn.ToPostables(cache.Resolve()) {
	// 	color.Green(r.Title())
	// }

	arn.PanicOnError(arn.DB.Set("Cache", "forum activity", cache))

	color.Green("Finished.")
}
