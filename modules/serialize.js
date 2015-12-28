'use strict'

let arn = require('../lib')
let passport = require('passport')

// Serialize
// This means we're reducing the user data to a single hash by which the user can be identified.
passport.serializeUser(function(user, done) {
	let now = new Date()
	user.lastLogin = now.toISOString()

	// Save in database
	arn.setUser(user.id, user)

	done(null, user.id)
})

module.exports = function() {
	// ...
}