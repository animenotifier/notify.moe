'use strict'

let fs = require('fs')
let shortid = require('shortid')
let passport = require('passport')
let FacebookStrategy = require('passport-facebook').Strategy

module.exports = function(aero) {
	let apiKeys = JSON.parse(fs.readFileSync('security/api-keys.json'))

	let facebookConfig = Object.assign({
	        callbackURL: '/auth/facebook/callback',
			profileFields: ['id', 'name', 'email'],
	        enableProof: false
	    },
	    apiKeys.facebook
	)

	passport.use(new FacebookStrategy(
        facebookConfig,
        function(accessToken, refreshToken, profile, done) {
			let fb = profile._json

			aero.db.get({
				ns: 'arn',
				set: 'FacebookToAccount',
				key: fb.id
			}, function(error, record, metadata, key) {
				if(error.code !== 0) {
					// New user
					let user = {
						id: shortid.generate(),
						nick: 'fb' + fb.id,
						firstName: fb.first_name,
						lastName: fb.last_name,
						email: fb.email,
						accounts: {
							facebook: fb.id
						}
					}

					if(user.email.endsWith('googlemail.com'))
						user.email = user.email.replace('googlemail.com', 'gmail.com')

					aero.db.put({
						ns: 'arn',
						set: 'FacebookToAccount',
						key: fb.id
					}, {
						userId: user.id
					}, function(error) {
						if(error.code !== 0)
							console.log('error: %s', error.message)
					})

					aero.db.put({
						ns: 'arn',
						set: 'NickToAccount',
						key: user.nick
					}, {
						userId: user.id
					}, function(error) {
						if(error.code !== 0)
							console.log('error: %s', error.message)
					})

					done(null, user)
				} else {
					// Existing user
					aero.db.get({
						ns: 'arn',
						set: 'Accounts',
						key: record.userId
					}, function(error, user, metadata, key) {
						done(error.code !== 0 ? error : null, user)
					})
				}
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