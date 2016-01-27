var arn = require('../lib')

arn.db.ready.then(() => {
	arn.forEach('Users', user => {
		if(!user.devices)
			return

		if(!user.pushEndpoints)
			user.pushEndpoints = {}

		Object.keys(user.devices).forEach(device => {
			user.pushEndpoints['https://android.googleapis.com/gcm/send/' + device] = {
				registered: user.devices[device]
			}
		})

		console.log(user.pushEndpoints)
		delete user.devices

		arn.set('Users', user.id, {
			pushEndpoints: user.pushEndpoints,
			devices: null
		})
	})
})