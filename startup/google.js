let chalk = require('chalk')

app.auth.google = {
	login: function*(google) {
		console.log(chalk.cyan('Google data:\n'), google)

		let email = google.emails.length > 0 ? google.emails[0].value : ''

		if(email.endsWith('googlemail.com'))
			email = email.replace('googlemail.com', 'gmail.com')

		try {
			let record = yield Promise.any([
				db.get('GoogleToUser', google.id),
				db.get('EmailToUser', email)
			])

			let user = yield db.get('Users', record.userId)

			// Existing user
			if(user && user.accounts)
				user.accounts.google = google.id

			db.set('GoogleToUser', google.id, {
				id: google.id,
				userId: user.id
			})

			console.log(`Existing user ${chalk.yellow(user.nick)} logged in`)

			return user
		} catch(_) {
			// New user
			let user = yield arn.registerNewUser({
				nick: 'g' + google.id,
				firstName: google.name.givenName ? google.name.givenName : '',
				lastName: google.name.familyName ? google.name.familyName : '',
				email,
				gender: google.gender ? google.gender : '',
				language: google.language ? google.language : '',
				ageRange: google.ageRange ? google.ageRange : null,
				accounts: {
					google: google.id
				}
			})

			db.set('GoogleToUser', google.id, {
				id: google.id,
				userId: user.id
			})

			console.log(`New user ${chalk.yellow(user.nick)} logged in`)

			return user
		}
	}
}