'use strict'

let passport = require('passport')
let session = require('express-session')
let apiKeys = require('../security/api-keys.json')
let FileStore = require('session-file-store')(session)

module.exports = function(aero) {
	// Session
	let sessionOptions = {
		store: new FileStore(),
	    name: 'sid',
	    secret: apiKeys.session.secret,
	    saveUninitialized: false,
	    resave: false,
	    cookie: {
	        secure: true
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