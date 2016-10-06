let chalk = require('chalk')

app.auth.twitter = {
	login: function*(twitter) {
		console.log(chalk.cyan('Twitter data:\n'), twitter)

		let email = twitter.email || ''

		if(email.endsWith('googlemail.com'))
			email = email.replace('googlemail.com', 'gmail.com')

		try {
			let record = yield Promise.any([
				db.get('TwitterToUser', twitter.id),
				db.get('EmailToUser', email)
			])

			let user = yield db.get('Users', record.userId)

			// Existing user
			if(user && user.accounts)
				user.accounts.twitter = twitter.id

			db.set('TwitterToUser', twitter.id, {
				id: twitter.id,
				userId: user.id
			})

			console.log(`Existing user ${chalk.yellow(user.nick)} logged in`)

			return user
		} catch(_) {
			let nameParts = twitter.name.split(' ')

			// New user
			let user = yield arn.registerNewUser({
				nick: 't' + twitter.id,
				email,
				firstName: nameParts[0],
				lastName: nameParts[1] ? nameParts[1] : '',
				tagline: twitter.description,
				language: twitter.lang,
				twitter: twitter.screen_name,
				accounts: {
					twitter: twitter.id
				}
			})

			db.set('TwitterToUser', twitter.id, {
				id: twitter.id,
				userId: user.id
			})

			console.log(`New user ${chalk.yellow(user.nick)} logged in`)

			return user
		}
	}
}