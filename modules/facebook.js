'use strict'

let fs = require('fs')
let arn = require('../lib')
let shortid = require('shortid')
let passport = require('passport')
let Promise = require('bluebird')
let FacebookStrategy = require('passport-facebook').Strategy

module.exports = function(aero) {
	let apiKeys = JSON.parse(fs.readFileSync('security/api-keys.json'))

	let facebookConfig = Object.assign({
	        callbackURL: '/auth/facebook/callback',
			profileFields: ['id', 'name', 'email', 'gender', 'age_range'],
	        enableProof: false,
			passReqToCallback: true
	    },
	    apiKeys.facebook
	)

	passport.use(new FacebookStrategy(
        facebookConfig,
        function(request, accessToken, refreshToken, profile, done) {
			let fb = profile._json
			let email = fb.email

			if(email.endsWith('googlemail.com'))
				email = email.replace('googlemail.com', 'gmail.com')

			Promise.any([
				arn.getAsync('FacebookToUser', fb.id),
				arn.getAsync('EmailToUser', email)
			])
			.then(record => arn.getUser(record.userId, function(error, user) {
				if(user && user.accounts)
					user.accounts.facebook = fb.id

				done(error, user)
			}))
			.catch(error => {
				// New user
				let now = new Date()
				let user = {
					id: shortid.generate(),
					nick: 'fb' + fb.id,
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
						anime: '',
						airingDate: 'AniList'
					},
					listProviders: {},
					ip: request.connection.remoteAddress,
					registered: now.toISOString(),
					lastLogin: now.toISOString()
				}

				arn.registerNewUser(
					user,
					arn.setAsync('FacebookToUser', fb.id, { userId: user.id })
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