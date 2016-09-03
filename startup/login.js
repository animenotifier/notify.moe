let useragent = require('useragent')

app.use((request, response, next) => {
	let user = request.user

	if(!user) {
		next()
		return
	}

	// Save last view date
	if(!request.url.startsWith('/api/') && request.url !== '/favicon.ico' && request.url !== '/service-worker.js') {
		arn.set('Users', user.id, {
			lastView: {
				url: request.url,
				date: (new Date()).toISOString()
			}
		})
	}

	// Save user agent
	user.agent = useragent.parse(request.headers['user-agent'])

	arn.set('Users', user.id, {
		agent: user.agent
	})

	// IP
	let newIP = request.headers['x-forwarded-for'] || request.connection.remoteAddress || ''

	if(!newIP) {
		next()
		return
	}

	if(user.ip === newIP) {
		next()
		return
	}

	user.ip = newIP

	// IP changed: Update location
	arn.set('Users', user.id, {
		ip: user.ip
	}).then(() => {
		arn.getLocation(user).then(location => {
			user.location = location
		}).catch(error => {
			user.location = null
		}).finally(() => {
			// Save in database
			arn.set('Users', user.id, {
				location: user.location
			})
		})
	})

	next()
})