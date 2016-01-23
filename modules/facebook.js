'use strict'

let shortid = require('shortid')
let passport = require('passport')
let Promise = require('bluebird')
let FacebookStrategy = require('passport-facebook').Strategy

module.exports = function(aero) {
	let facebookConfig = Object.assign({
	        callbackURL: 'https://notify.moe/auth/facebook/callback',
			profileFields: ['id', 'name', 'email', 'gender', 'age_range'],
	        enableProof: false,
			passReqToCallback: true
	    },
	    arn.apiKeys.facebook
	)

	passport.use(new FacebookStrategy(
        facebookConfig,
        function(request, accessToken, refreshToken, profile, done) {
			let fb = profile._json
			let email = fb.email

			if(email.endsWith('googlemail.com'))
				email = email.replace('googlemail.com', 'gmail.com')

			Promise.any([
				arn.get('FacebookToUser', fb.id),
				arn.get('EmailToUser', email)
			])
			.then(record => arn.get('Users', record.userId).then(user => {
				if(user && user.accounts)
					user.accounts.facebook = fb.id

				done(undefined, user)
			})).catch(error => {
				console.error(error)

				// New user
				let now = new Date()
				let user = {
					id: shortid.generate(),
					nick: 'fb' + fb.id,
					role: email === 'e.urbach@gmail.com' ? 'admin' : '',
					firstName: fb.first_name,
					lastName: fb.last_name,
					email: email ? email : '',
					gender: fb.gender,
					language: '',
					ageRange: fb.age_range,
					accounts: {
						facebook: fb.id
					},
					tagline: '',
					website: '',
					providers: {
						list: 'AniList',
						anime: 'Nyaa',
						airingDate: 'AniList'
					},
					listProviders: {},
					sortBy: 'airingDate',
					devices: {},
					registered: now.toISOString(),
					lastLogin: now.toISOString()
				}

				arn.registerNewUser(
					user,
					arn.set('FacebookToUser', fb.id, { userId: user.id })
				).then(function() {
					done(undefined, user)
				})
			})
        }
	))

	// Facebook login
	aero.get('/auth/facebook', passport.authenticate('facebook', {
	    scope: [
	        'email',
	        'public_profile'
	    ]
	}))

	// Facebook callback
	aero.get('/auth/facebook/callback',
	    passport.authenticate('facebook', {
	        successRedirect: '/',
	        failureRedirect: '/login'
	    })
	)
}