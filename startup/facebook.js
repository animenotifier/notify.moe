app.auth.facebook = {
	login: function*(fb) {
		console.log(chalk.cyan('Facebook data:\n'), fb)

		let email = fb.email || ''

		if(email.endsWith('googlemail.com'))
			email = email.replace('googlemail.com', 'gmail.com')

		try {
			let record = yield Promise.any([
				db.get('FacebookToUser', fb.id),
				db.get('EmailToUser', email)
			])

			let user = yield db.get('Users', record.userId)

			// Existing user
			if(user && user.accounts)
				user.accounts.facebook = fb.id

			db.set('FacebookToUser', fb.id, {
				id: fb.id,
				userId: user.id
			})

			console.log(`Existing user ${chalk.yellow(user.nick)} logged in`)

			return user
		} catch(_) {
			// New user
			let user = yield arn.registerNewUser({
				nick: 'fb' + fb.id,
				firstName: fb.first_name,
				lastName: fb.last_name,
				email: email ? email : '',
				gender: fb.gender,
				ageRange: fb.age_range,
				accounts: {
					facebook: fb.id
				}
			})

			db.set('FacebookToUser', fb.id, {
				id: fb.id,
				userId: user.id
			})

			console.log(`New user ${chalk.yellow(user.nick)} logged in`)

			return user
		}
	}
}