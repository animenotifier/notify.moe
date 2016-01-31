'use strict'

let passport = require('passport')
let session = require('express-session')
let FileStore = require('session-file-store')(session)

module.exports = function(aero) {
	// Session
	let sessionOptions = {
		store: new FileStore(),
	    name: 'sid',
	    secret: arn.apiKeys.session.secret,
	    saveUninitialized: false,
	    resave: false,
	    cookie: {
	        secure: true,
			maxAge: 6 * 30 * 24 * 60 * 60 * 1000
	    }
	}

	// Middleware
	aero.use(
	    session(sessionOptions),
	    passport.initialize(),
	    passport.session()
	)

	// Logout
	aero.get('/logout', function(req, res) {
	    req.logout()
	    res.redirect('/')
	})
}