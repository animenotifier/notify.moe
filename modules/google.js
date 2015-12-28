'use strict'

let fs = require('fs')
let arn = require('../lib')
let shortid = require('shortid')
let passport = require('passport')
let Promise = require('bluebird')
let GoogleStrategy = require('passport-google-oauth').OAuth2Strategy

module.exports = function(aero) {
	let apiKeys = JSON.parse(fs.readFileSync('security/api-keys.json'))

	let googleConfig = Object.assign({
	        callbackURL: '/auth/google/callback',
			passReqToCallback: true
	    },
	    apiKeys.google
	)

	passport.use(new GoogleStrategy(
        googleConfig,
        function(request, accessToken, refreshToken, profile, done) {
			let google = profile._json
			let email = google.emails.length > 0 ? google.emails[0].value : ''

			if(email.endsWith('googlemail.com'))
				email = email.replace('googlemail.com', 'gmail.com')

			Promise.any([
				arn.getAsync('GoogleToUser', google.id),
				arn.getAsync('EmailToUser', email)
			])
			.then(record => arn.getUser(record.userId, function(error, user) {
				if(user && user.accounts)
					user.accounts.google = google.id

				done(error, user)
			}))
			.catch(error => {
				// New user
				let now = new Date()
				let user = {
					id: shortid.generate(),
					nick: 'g' + google.id,
					firstName: google.name.givenName,
					lastName: google.name.familyName,
					email: email,
					gender: google.gender,
					language: google.language,
					ageRange: google.ageRange,
					registered: now.toISOString(),
					lastLogin: now.toISOString(),
					ip: request.connection.remoteAddress,
					accounts: {
						google: google.id
					}
				}

				Promise.all([
					arn.setAsync('GoogleToUser', google.id, { userId: user.id }),
					arn.setAsync('NickToUser', user.nick, { userId: user.id }),
					arn.setAsync('EmailToUser', user.email, { userId: user.id })
					// arn.setAsync('AnimeList', user.id, {})
				])

				done(undefined, user)
			})
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