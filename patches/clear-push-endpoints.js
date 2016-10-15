var arn = require('../lib')

arn.db.ready.then(() => {
	arn.db.forEach('Users', user => {
		if(!user.pushEndpoints)
			user.pushEndpoints = {}

		Object.keys(user.pushEndpoints).forEach(endpoint => {
			let subscription = user.pushEndpoints[endpoint]

			if(!subscription.keys)
				delete user.pushEndpoints[endpoint]
		})

		console.log(user.pushEndpoints)

		arn.set('Users', user.id, {
			pushEndpoints: user.pushEndpoints
		})
	})
})