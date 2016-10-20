let zlib = require('zlib')

let processTitle = title => title.replace(/[^A-Za-z0-9.:!'"+ ]/g, ' ').replace(/  /g, ' ')

// Create search index
arn.db.ready.then(() => {
	arn.animeCount = 0
	arn.animeToId = {}

	arn.db.forEach('Anime', anime => {
		if(anime.type === 'Music')
			return

		arn.animeCount++

		if(anime.title.romaji)
			arn.animeToId[processTitle(anime.title.romaji)] = anime.id

		if(anime.title.english)
			arn.animeToId[processTitle(anime.title.english)] = anime.id
	}).then(() => {
		arn.animeToIdCount = Object.keys(arn.animeToId).length
		arn.animeToIdJSONString = JSON.stringify(arn.animeToId)

		zlib.gzip(arn.animeToIdJSONString, function(error, gzippedJSON) {
			arn.animeToIdJSONStringGzipped = gzippedJSON
		})
	})
})