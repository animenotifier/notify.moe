'use strict'

let arn = require('../lib')
let useragent = require('useragent')

arn.db.ready.then(() => {
	arn.forEach('Users', user => {
		if(typeof user.agent !== 'string')
			return

		let parsed = useragent.parse(user.agent)
		console.log(parsed)

		arn.set('Users', user.id, {
			agent: parsed
		})
	})
})