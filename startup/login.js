let passport = require('passport')
let session = require('express-session')
let FileStore = require('session-file-store')(session)

const cookieDurationInSeconds = 6 * 30 * 24 * 60 * 60

// Session
let sessionOptions = {
	store: new FileStore(),
    name: 'sid',
    secret: arn.apiKeys.session.secret,
    saveUninitialized: false,
    resave: false,
    cookie: {
        secure: true,
		maxAge: cookieDurationInSeconds * 1000
    }
}

// Middleware
app.use(
    session(sessionOptions),
    passport.initialize(),
    passport.session(),
	(request, response, next) => {
		if(!request.isAuthenticated())
			request.user = null
		
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
	}
)

// Logout
app.get('/logout', function(req, res) {
    req.logout()
	
	if(req.session && req.session.destroy) {
		req.session.destroy(function(err) {
			res.redirect('/')
		})
	}
})