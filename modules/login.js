'use strict'

let passport = require('passport')
let session = require('express-session')
let apiKeys = require('../security/api-keys.json')

module.exports = function(aero) {
	// Session
	let sessionOptions = {
	    name: 'sid',
	    secret: require('crypto').randomBytes(64).toString('hex'),
	    saveUninitialized: true,
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