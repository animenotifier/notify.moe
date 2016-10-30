let bluebird = require('bluebird')

let addYieldHandler = bluebird.coroutine(function*() {
	try {
		yield []
		yield {}
	} catch(error) {
		bluebird.coroutine.addYieldHandler(function(value) {
			if(Array.isArray(value))
				return bluebird.all(value)

			if(typeof value === 'object')
				return bluebird.props(value)
		})
	}
})

addYieldHandler()