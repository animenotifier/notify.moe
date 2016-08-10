global.arn = require('./lib')

let app = require('aero')()
let fs = require('fs')
let path = require('path')
let zlib = require('zlib')
let bodyParser = require('body-parser')
let request = require('request-promise')
let paypal = require('paypal-rest-sdk')

paypal.configure({
	mode: 'sandbox',
	client_id: arn.apiKeys.paypal.clientID,
	client_secret: arn.apiKeys.paypal.clientSecret
})

// Rewrite URLs
app.rewrite(function(request, response) {
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
/*if(!arn.production) {
	app.use(function(request, response, next) {
		let start = new Date()
		next()
		let end = new Date()

		if(request.user && request.user.nick)
			console.log(request.url, '|', end - start, 'ms', '|', request.user.nick)
		else
			console.log(request.url, '|', end - start, 'ms')
	})
}*/

// For POST requests
app.use(bodyParser.json())

// Service worker routes
app.sendFile('service-worker.js', 'worker/service-worker.js')
app.sendFile('cache-polyfill.js', 'worker/cache-polyfill.js')

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

// Create search index
arn.db.ready.then(() => {
	let processTitle = title => title.replace(/[^A-Za-z0-9.:!'"+ ]/g, ' ').replace(/  /g, ' ')
	arn.animeCount = 0
	arn.animeToId = {}

	arn.forEach('Anime', anime => {
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

// Authorize AniList
let anilist = arn.listProviders.AniList
anilist.authorize().then(accessToken => {
	console.log('AniList API token:', accessToken)
})

// Load all modules
fs.readdirSync('modules').forEach(mod => require('./modules/' + mod)(app))

// Start the server
app.run()