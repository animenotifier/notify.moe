'use strict'

module.exports = function(aero) {
	let fs = require('fs')
	let passport = require('passport')
	let GoogleStrategy = require('passport-google-oauth').OAuth2Strategy

	let apiKeys = JSON.parse(fs.readFileSync('security/api-keys.json'))

	let googleConfig = Object.assign({
	        callbackURL: '/auth/google/callback'
	    },
	    apiKeys.google
	)

	passport.use(new GoogleStrategy(
        googleConfig,
        function(accessToken, refreshToken, profile, done) {
            done(null, profile)
        }
	))

	// Google login
	aero.get('/auth/google', passport.authenticate('google', {
	    scope: [
	        'https://www.googleapis.com/auth/plus.login',
	        'email'
	    ]
	}))

	// Google callback
	aero.get('/auth/google/callback',
	    passport.authenticate('google', {
	        successRedirect: '/',
	        failureRedirect: '/login'
	    })
	)
}