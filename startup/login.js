let passport = require('passport')
let session = require('express-session')
let FileStore = require('session-file-store')(session)

const cookieDurationInSeconds = 6 * 30 * 24 * 60 * 60

// Session
let sessionOptions = {
	store: new FileStore(),
    name: 'sid',
    secret: arn.apiKeys.session.secret,
    saveUninitialized: false,
    resave: false,
    cookie: {
        secure: true,
		maxAge: cookieDurationInSeconds * 1000
    }
}

// Middleware
app.use(
    session(sessionOptions),
    passport.initialize(),
    passport.session()
)

// Logout
app.get('/logout', function(req, res) {
    req.logout()
	req.session.destroy(function(err) {
		res.redirect('/')
	})
})