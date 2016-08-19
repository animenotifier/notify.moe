let passport = require('passport')
let TwitterStrategy = require('passport-twitter').Strategy

let twitterConfig = {
    callbackURL: arn.production ? 'https://notify.moe/auth/twitter/callback' : 'https://127.0.0.1:5001/auth/twitter/callback',
	passReqToCallback: true,
	consumerKey: arn.apiKeys.twitter.clientID,
	consumerSecret: arn.apiKeys.twitter.clientSecret
}

passport.use(new TwitterStrategy(
    twitterConfig,
    function(request, accessToken, refreshToken, profile, done) {
		console.log(chalk.cyan('Twitter data:'), (profile && profile._json) ? profile._json : profile)
		
		let twitter = profile._json
		let email = twitter.email || ''

		if(email.endsWith('googlemail.com'))
			email = email.replace('googlemail.com', 'gmail.com')

		Promise.any([
			arn.get('TwitterToUser', twitter.id),
			arn.get('EmailToUser', email)
		])
		.then(record => arn.get('Users', record.userId).then(user => {
			// Existing user
			if(user && user.accounts)
				user.accounts.twitter = twitter.id

			done(undefined, user)
		})).catch(error => {
			let nameParts = twitter.name.split(' ')
			
			// New user
			arn.registerNewUser({
				nick: 't' + twitter.id,
				email,
				firstName: nameParts[0],
				lastName: nameParts[1] ? nameParts[1] : '',
				tagline: twitter.description,
				language: twitter.lang,
				accounts: {
					twitter: twitter.id
				}
			}).then(user => {
				arn.set('TwitterToUser', twitter.id, {
					userId: user.id
				})
				
				done(undefined, user)
			}).catch(error => done(error, false))
		})
    }
))

// Twitter login
app.get('/auth/twitter', passport.authenticate('twitter'))

// Twitter callback
app.get('/auth/twitter/callback',
    passport.authenticate('twitter', {
        successRedirect: '/',
        failureRedirect: '/login'
    })
)