'use strict'

let arn = require('../lib')
let passport = require('passport')

// Deserialize
// This means a web page is requesting full user data by some kind of hash.
passport.deserializeUser(function(userId, done) {
	arn.getUser(userId, done)
})

module.exports = function() {
	// ...
}