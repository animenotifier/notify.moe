let passport = require('passport')
let useragent = require('useragent')

// Serialize
// This means we're reducing the user data to a single hash by which the user can be identified.
passport.serializeUser(function(request, user, done) {
	let now = new Date()
	user.lastLogin = now.toISOString()
	user.agent = useragent.parse(request.headers['user-agent'])
	
	// Save in database
	arn.set('Users', user.id, user).then(() => done(null, user.id))
})