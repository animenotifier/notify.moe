let request = require('request-promise')
let xml2js = require('xml2js')

const COMPLETED = 2
const WATCHING = 1

class MyAnimeList {
	constructor() {
		this.headers = {
			'User-Agent': 'Anime Release Notifier'
		}

		this.xmlParser = new xml2js.Parser({
			explicitArray: false,	// Don't put single nodes into an array
			ignoreAttrs: true,		// Ignore attributes and only create text nodes
			trim: true,
			normalize: true,
			explicitRoot: false
		})
	}

	getAnimeListUrl(userName) {
		return `http://myanimelist.net/animelist/${userName}&status=${WATCHING}`
	}

	getAnimeList(userName, callback) {
		return request({
			uri: `http://myanimelist.net/malappinfo.php?u=${userName}&status=all&type=anime`,
			method: 'GET',
			headers: this.headers
		}).then(body => {
			this.xmlParser.parseString(body, (error, json) => {
				if(error) {
					callback(error, [])
					return
				}

				let watching = []

				if(!Array.isArray(json.anime)) {
					return callback(undefined, watching)
				}

				watching = json.anime
				.filter(entry => parseInt(entry.my_status) === WATCHING)
				.map(entry => {
					let episodesWatched = parseInt(entry.my_watched_episodes)
					let nextEpisodeToWatch = episodesWatched + 1
					let episodesOffset = 0

					return {
						title: entry.series_title,
						image: entry.series_image,
						url: 'http://myanimelist.net/anime/' + entry.series_animedb_id,
						providerId: parseInt(entry.series_animedb_id),
						airingDate: null,
						episodes: {
							watched: episodesWatched ? episodesWatched : 0,
							next: nextEpisodeToWatch,
							available: 0,
							max: entry.series_episodes ? parseInt(entry.series_episodes) : -1,
							offset: episodesOffset
						}
					}
				})

				let tasks = []
				watching.forEach(anime => {
					tasks.push(arn.getAnimeIdBySimilarTitle(anime, 'MyAnimeList').then(match => {
						anime.id = match ? match.id : null
						anime.similarity = match ? match.similarity : 0
					}))
				})

				Promise.all(tasks).then(() => callback(undefined, watching))
			})
		}).catch(error => {
			callback(error, [])
		})
	}
}

module.exports = new MyAnimeList()