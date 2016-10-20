let arn = require('../lib')
let useragent = require('useragent')

arn.db.ready.then(() => {
	arn.db.forEach('Users', user => {
		if(typeof user.agent !== 'string')
			return

		let parsed = useragent.parse(user.agent)
		console.log(parsed)

		arn.db.set('Users', user.id, {
			agent: parsed
		})
	})
})