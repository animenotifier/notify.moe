let nicks = {}

arn.db.ready.then(() => {
	arn.db.forEach('Users', user => {
		// arn.db.get('NickToUser', user.nick).then(record => {
		// 	if(record.userId !== user.id)
		// 		console.log(user.nick)
		// })
		if(nicks[user.nick]) {
			console.log('Double nick: ' + user.nick)

			if(user.accounts.google) {
				arn.changeNick(user, 'g' + user.accounts.google)
			} else if(user.accounts.facebook) {
				arn.changeNick(user, 'fb' + user.accounts.facebook)
			} else if(user.accounts.twitter) {
				arn.changeNick(user, 't' + user.accounts.twitter)
			}
		} else {
			arn.db.set('NickToUser', user.nick, {
				userId: user.id
			})
		}

		nicks[user.nick] = user.id
	})
})