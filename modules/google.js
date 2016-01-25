'use strict'

let shortid = require('shortid')
let passport = require('passport')
let Promise = require('bluebird')
let GoogleStrategy = require('passport-google-oauth').OAuth2Strategy

module.exports = function(aero) {
	let googleConfig = Object.assign({
	        callbackURL: arn.production ? 'https://notify.moe/auth/google/callback' : '/auth/google/callback',
			passReqToCallback: true
	    },
	    arn.apiKeys.google
	)

	passport.use(new GoogleStrategy(
        googleConfig,
        function(request, accessToken, refreshToken, profile, done) {
			let google = profile._json
			let email = google.emails.length > 0 ? google.emails[0].value : ''

			if(email.endsWith('googlemail.com'))
				email = email.replace('googlemail.com', 'gmail.com')

			Promise.any([
				arn.get('GoogleToUser', google.id),
				arn.get('EmailToUser', email)
			])
			.then(record => arn.get('Users', record.userId).then(user => {
				if(user && user.accounts)
					user.accounts.google = google.id

				done(undefined, user)
			})).catch(error => {
				console.error(error)

				// New user
				let now = new Date()
				let user = {
					id: shortid.generate(),
					nick: 'g' + google.id,
					role: email === 'e.urbach@gmail.com' ? 'admin' : '',
					firstName: google.name.givenName ? google.name.givenName : '',
					lastName: google.name.familyName ? google.name.familyName : '',
					email: email,
					gender: google.gender ? google.gender : '',
					language: google.language,
					ageRange: google.ageRange ? google.ageRange : null,
					accounts: {
						google: google.id
					},
					tagline: '',
					website: '',
					providers: {
						list: 'AniList',
						anime: 'CrunchyRoll',
						airingDate: 'AniList'
					},
					listProviders: {},
					sortBy: 'airingDate',
					devices: {},
					registered: now.toISOString(),
					lastLogin: now.toISOString(),
				}

				arn.registerNewUser(
					user,
					arn.set('GoogleToUser', google.id, { userId: user.id })
				).then(function() {
					done(undefined, user)
				})
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