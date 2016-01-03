'use strict'

let arn = require('../lib')
let passport = require('passport')

// Serialize
// This means we're reducing the user data to a single hash by which the user can be identified.
passport.serializeUser(function(request, user, done) {
	let now = new Date()
	user.lastLogin = now.toISOString()
	user.ip = request.headers['x-forwarded-for'] || request.connection.remoteAddress || ''
	user.agent = request.headers['user-agent']

	arn.setUserAsync(user.id, user).then(() => {
		done(null, user.id)

		if(!user.ip)
			return

		arn.getLocation(user).then(location => {
			user.location = location
		}).catch(error => {
			user.location = null
		}).finally(() => {
			// Save in database
			arn.setUser(user.id, user)
		})
	})
})

module.exports = function() {
	// ...
}