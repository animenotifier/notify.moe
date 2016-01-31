'use strict'

let passport = require('passport')
let session = require('express-session')
let FileStore = require('session-file-store')(session)
let AerospikeStore = require('connect-aerospike')(session)

const cookieDurationInSeconds = 6 * 30 * 24 * 60 * 60

module.exports = function(aero) {
	// Session
	let sessionOptions = {
		store: new AerospikeStore({
			ttl: 24 * 60 * 60,
			hosts: ['127.0.0.1:3000'],
			prefix: 'arn:',
			ns: 'arn',
			st: 'Sessions'
		}),
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