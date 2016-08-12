let arn = require('../lib')

arn.db.ready.then(() => {
	arn.forEach('Users', user => {
		arn.set('Users', user.id, {
			titleLanguage: 'romaji'
		})
	})
})