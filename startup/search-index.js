let zlib = require('zlib')

let processTitle = title => title.replace(/[^A-Za-z0-9.:!'"+ ]/g, ' ').replace(/  /g, ' ')

// Create search index
arn.db.ready.then(() => {
	arn.animeCount = 0
	arn.titleToId = {}

	arn.db.forEach('Anime', anime => {
		if(anime.type === 'Music')
			return

		arn.animeCount++

		if(anime.title.romaji)
			arn.titleToId[processTitle(anime.title.romaji)] = anime.id

		if(anime.title.english)
			arn.titleToId[processTitle(anime.title.english)] = anime.id
	}).then(() => {
		arn.titleToIdCount = Object.keys(arn.titleToId).length
		arn.titleToIdJSONString = JSON.stringify(arn.titleToId)

		zlib.gzip(arn.titleToIdJSONString, function(error, gzippedJSON) {
			arn.titleToIdJSONStringGzipped = gzippedJSON
		})
	})
})