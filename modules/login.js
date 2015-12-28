'use strict'

let passport = require('passport')
let session = require('express-session')
let Promise = require('bluebird')

module.exports = function(aero) {
	// getUser
	aero.getUser = function(userId, callback) {
		aero.db.get({
			ns: 'arn',
			set: 'Accounts',
			key: userId
		}, function(error, user, metadata, key) {
			callback(error.code !== 0 ? error : null, user)
		})
	}

	aero.getUserAsync = Promise.promisify(aero.getUser)

	// Serialize
	// This means we're reducing the user data to a single hash by which the user can be identified.
	passport.serializeUser(function(user, done) {
		let now = new Date()
		user.lastLogin = now.toISOString()

		// Save in database
		aero.db.put({
			ns: 'arn',
			set: 'Accounts',
			key: user.id
		}, user, function(error) {
			if(error.code !== 0)
				console.log('error: %s', error.message)
		})

	    done(null, user.id)
	})

	// Deserialize
	// This means a web page is requesting full user data by some kind of hash.
	passport.deserializeUser(function(userId, done) {
		aero.getUser(userId, done)
	})

	// Session
	let sessionOptions = {
	    name: 'sid',
	    secret: require('crypto').randomBytes(64).toString('hex'),
	    saveUninitialized: true,
	    resave: false,
	    cookie: {
	        secure: false
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