let arn = require('../lib')

arn.db.ready.then(() => {
	arn.db.forEach('Users', user => {
		arn.db.set('Users', user.id, {
			titleLanguage: 'romaji'
		})
	})
})