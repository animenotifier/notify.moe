'use strict'

global.arn = require('./lib')

let aero = require('aero')
let fs = require('fs')
let path = require('path')
let zlib = require('zlib')
let bodyParser = require('body-parser')
let request = require('request-promise')
let RateLimiter = require('limiter').RateLimiter

// Start the server
aero.run()

// Rewrite URLs
aero.preRoute(function(request, response) {
	if(request.headers.host.indexOf('animereleasenotifier.com') !== -1) {
        response.redirect('https://notify.moe' + request.url)
        return true
    }

	if(request.url.startsWith('/+'))
		request.url = '/user/' + request.url.substring(2)
	else if(request.url.startsWith('/_/+'))
		request.url = '/_/user/' + request.url.substring(4)
})

// Log requests
aero.use(function(request, response, next) {
	let start = new Date()
	next()
	let end = new Date()

	if(request.user && request.user.nick)
		console.log(request.url, '|', end - start, 'ms', '|', request.user.nick)
	else
		console.log(request.url, '|', end - start, 'ms')
})

// For POST requests
aero.use(bodyParser.json())

// Web app manifest
aero.get('manifest.json', (request, response) => {
	let manifest = {
		name: 'Anime Notifier',
		short_name: 'Anime Notifier',
		icons: [{
			src: 'images/characters/arn-waifu.png',
			sizes: '300x300',
			type: 'image/png'
		}],
		start_url: '/',
		display: 'standalone',
		lang: 'en',
		gcm_sender_id: arn.apiKeys.gcm.senderID
	}

	let manifestString = JSON.stringify(manifest)

	response.writeHead(200, {
		'Content-Type': 'application/json',
		'Content-Length': Buffer.byteLength(manifestString, 'utf8'),
		'Cache-Control': 'max-age=3600'
	})

	response.end(manifestString)
})

let serveFile = fileName => {
	return (request, response) => {
		let filePath = path.join(__dirname, fileName)
		let stat = fs.statSync(filePath)

		response.writeHead(200, {
			'Content-Type': 'application/javascript',
			'Content-Length': stat.size
		})

		fs.createReadStream(filePath).pipe(response)
	}
}

aero.get('service-worker.js', serveFile('worker/service-worker.js'))
aero.get('cache-polyfill.js', serveFile('worker/cache-polyfill.js'))

// Send slack messages
arn.on('new user', function(user) {
	// Ignore my own attempts on empty databases
	if(user.email === 'e.urbach@gmail.com')
		return

	let host = 'https://notify.moe'
	let webhook = 'https://hooks.slack.com/services/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'

	request.post({
		url: webhook,
		body: JSON.stringify({
			text: `<${host}/users|${user.firstName} ${user.lastName} (${user.email})>`
		})
	}).then(body => {
		console.log(`Sent slack message about the new user registration: ${user.email}`)
	}).catch(error => {
		console.error('Error sending slack message:', error, error.stack)
	})
})

arn.on('new forum reply', function(link, userName) {
	let webhook = 'https://hooks.slack.com/services/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'

	request.post({
		url: webhook,
		body: JSON.stringify({
			text: `<${link}|${userName}>`
		})
	}).then(body => {
		console.log(`Sent slack message about a new forum reply from ${userName}`)
	}).catch(error => {
		console.error('Error sending slack message:', error, error.stack)
	})
})

// Create search index
arn.on('database ready', function() {
	arn.cacheLimiter.removeTokens(1, () => {
		let processTitle = title => title.replace(/[^A-Za-z0-9.:!'"+ ]/g, ' ').replace(/  /g, ' ')
		arn.animeToId = {}
		arn.forEach('Anime', anime => {
			if(anime.type === 'Music')
				return

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
})

// Authorize AniList
let anilist = arn.listProviders.AniList
anilist.authorize().then(accessToken => {
	console.log('AniList API token:', accessToken)

	// Check forum thread replies
	setInterval(anilist.checkForumReplies.bind(anilist), 5 * 60 * 1000)
	anilist.checkForumReplies()
})

// Do a full import every 12 hours
let anilistImport = function() {
	if(!arn.db)
		return Promise.reject('No database connection')

	console.log('Doing full anilist import')

	let limiter = new RateLimiter(1, 1100)

	return arn.listProviders.AniList.authorize().then(() => {
		let maxPage = 238
		for(let page = 1; page <= maxPage; page++) {
			limiter.removeTokens(1, function() {
				arn.listProviders.AniList.getAnimeFromPage(page).then(animeList => {
					let tasks = animeList.map(anime => arn.set('Anime', anime.id, anime))
					Promise.all(tasks).then(() => console.log('Finished importing page', page))
				})
			})
		}
	})
}
setInterval(anilistImport, 12 * 60 * 60 * 1000)

// Load all modules
fs.readdirSync('modules').forEach(mod => require('./modules/' + mod)(aero))