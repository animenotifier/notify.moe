let passport = require('passport')

// Deserialize
// This means a web page is requesting full user data by some kind of hash.
passport.deserializeUser(function(userId, done) {
	return arn.get('Users', userId).then(user => done(undefined, user)).catch(error => console.error('Deserialize error:', error))
})