const zlib = require('zlib')
const processTitle = title => title.replace(/[^A-Za-z0-9.:!'"+ ]/g, ' ').replace(/  /g, ' ')
const bestCompressionOptions = {
	level: zlib.Z_BEST_COMPRESSION
}

function updateSearchIndex() {
	console.log(chalk.cyan('↻'), 'Updating search index...')

	let animeCount = 0
	let titleToId = {}

	arn.db.forEach('Anime', anime => {
		if(anime.type === 'Music')
			return

		animeCount++

		if(anime.title.romaji)
			titleToId[processTitle(anime.title.romaji)] = anime.id

		if(anime.title.english)
			titleToId[processTitle(anime.title.english)] = anime.id
	}).then(() => {
		let titleCount = Object.keys(titleToId).length
		let titleToIdString = JSON.stringify(titleToId)

		zlib.gzip(titleToIdString, bestCompressionOptions, function(error, titleToIdStringGzipped) {
			arn.db.set('Cache', 'animeStats', {
				animeCount,
				titleCount
			})

			arn.db.set('Cache', 'animeTitleToId', {
				raw: titleToId,
				compressed: titleToIdStringGzipped
			})

			console.log(chalk.green('✔'), 'Updated search index.')
		})
	})
}

arn.repeatedly(1 * hours, updateSearchIndex)