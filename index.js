'use strict';

let aero = require('aero')
let passport = require('passport')
let session = require('express-session')
let fs = require('fs')
let apiKeys = JSON.parse(fs.readFileSync('security/api-keys.json'))

// Google
let GoogleStrategy = require('passport-google-oauth').OAuth2Strategy

let googleConfig = Object.assign(
	{
		callbackURL: '/auth/google/callback'
	},
	apiKeys.google
)

passport.use(
	new GoogleStrategy(
		googleConfig,
		function(accessToken, refreshToken, profile, done) {
			done(null, profile)
		}
	)
)

// Facebook
let FacebookStrategy = require('passport-facebook').Strategy

let facebookConfig = Object.assign(
	{
		callbackURL: '/auth/facebook/callback',
		enableProof: false
	},
	apiKeys.facebook
)

passport.use(
	new FacebookStrategy(
		facebookConfig,
		function(accessToken, refreshToken, profile, done) {
			done(null, profile)
		}
	)
)

passport.serializeUser(function(user, done) {
	console.log('serialize', user)
    done(null, user.id)
})

passport.deserializeUser(function(id, done) {
	console.log('deserialize', id)
    done(null, {
        id
    })
})

let sessionOptions = {
	name: 'sid',
	secret: require('crypto').randomBytes(64).toString('hex'),
	saveUninitialized: true,
	resave: false,
	cookie: {
		secure: false
	}
}

// Middleware
aero.use(
    session(sessionOptions),
    passport.initialize(),
    passport.session()
)

// Start the server
aero.run()

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

// Google login
aero.get('/auth/google', passport.authenticate('google', {
	scope: 'https://www.googleapis.com/auth/plus.login'
}))

// Google callback
aero.get('/auth/google/callback',
	passport.authenticate('google', {
		successRedirect: '/',
		failureRedirect: '/login'
	})
)

// Logout
aero.get('/logout', function(req, res) {
    req.logout()
    res.redirect('/')
})