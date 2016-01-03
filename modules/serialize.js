'use strict'

let arn = require('../lib')
let passport = require('passport')
let apiKeys = require('../security/api-keys.json')
let get = require('request-promise')

// Serialize
// This means we're reducing the user data to a single hash by which the user can be identified.
passport.serializeUser(function(request, user, done) {
	let now = new Date()
	user.lastLogin = now.toISOString()
	user.ip = request.headers['x-forwarded-for'] || request.connection.remoteAddress
	user.agent = request.headers['user-agent']

	let locationAPI = `http://api.ipinfodb.com/v3/ip-city/?key=${apiKeys.ipInfoDB.clientID}&ip=${user.ip}&format=json`

	get(locationAPI).then(body => {
		user.location = JSON.parse(body)
	}).catch(error => {
		user.location = null
	}).finally(() => {
		// Save in database
		arn.setUser(user.id, user)
	})

	done(null, user.id)
})

module.exports = function() {
	// ...
}