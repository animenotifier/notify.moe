let Promise = require('bluebird')
let request = require('request-promise')

let autoCorrectUserNames = [
	/anilist.co\/user\/(.*)/,
	/anilist.co\/animelist\/(.*)/,
	/hummingbird.me\/users\/(.*?)\/library/,
	/hummingbird.me\/users\/(.*)/,
	/anime-planet.com\/users\/(.*?)\/anime/,
	/anime-planet.com\/users\/(.*)/,
	/myanimelist.net\/profile\/(.*)/,
	/myanimelist.net\/animelist\/(.*?)\?/,
	/myanimelist.net\/animelist\/(.*)/,
	/myanimelist.net\/(.*)/,
	/myanimelist.com\/(.*)/,
	/twitter.com\/(.*)/,
	/osu.ppy.sh\/u\/(.*)/
]

arn.fixListProviderUserName = function(userName) {
	userName = userName.trim()

	for(let regex of autoCorrectUserNames) {
		let match = regex.exec(userName)

		if(match !== null) {
			userName = match[1]
			break
		}
	}

	return userName
}

arn.fixNick = function(nick) {
	nick = nick.replace(/[\W\s\d]/g, '')

	if(nick)
		nick = nick[0].toUpperCase() + nick.substring(1)

	return nick
}

arn.registerNewUser = function(user, customTask) {
	let tasks = [
		arn.set('NickToUser', user.nick, { userId: user.id })
	]
	
	if(user.email)
		tasks.push(arn.set('EmailToUser', user.email, { userId: user.id }))
	
	tasks.push(customTask)
	
	return Promise.all(tasks).then(function() {
		arn.events.emit('new user', user)
	}).catch(error => {
		console.error(`New user <${user.email}> registration error:`, error)
	})
}

arn.getUserByNick = Promise.coroutine(function*(nick) {
	// Very old Android app requests
	if(nick.indexOf('&animeProvider=') !== -1)
		return Promise.reject('Old Android app request')
	
	let record = yield arn.get('NickToUser', nick)
	return arn.get('Users', record.userId)
})

arn.changeNick = function(user, newNick) {
	let oldNick = user.nick

	if(oldNick === newNick)
		return Promise.resolve()

	return arn.get('NickToUser', newNick).then(record => {
		return Promise.reject('Username is already taken.')
	}).catch(error => {
		user.nick = newNick

		return Promise.all([
			arn.remove('NickToUser', oldNick),
			arn.set('NickToUser', newNick, { userId: user.id }),
			arn.set('Users', user.id, user)
		])
	})
}

arn.auth = (req, res, role) => {
	if(!req.user) {
		res.end('Not logged in!')
		return false
	}

	if(req.user.role !== 'admin' && req.user.role !== role) {
		res.end('Not authorized to view this page!')
		return false
	}

	return true
}

arn.isActiveUser = function(user) {
	if(user.nick.startsWith('g'))
		return false

	if(user.nick.startsWith('fb'))
		return false

	let listProviderName = user.providers.list

	if(!listProviderName)
		return false

	let listProvider = user.listProviders[listProviderName]

	if(!listProvider || !listProvider.userName)
		return false

	return true
}

arn.getLocation = function(user) {
	let locationAPI = `http://api.ipinfodb.com/v3/ip-city/?key=${arn.apiKeys.ipInfoDB.clientID}&ip=${user.ip}&format=json`
	return request(locationAPI).then(JSON.parse)
}