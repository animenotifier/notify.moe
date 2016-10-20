let arn = require('../lib')
let chalk = require('chalk')
let Promise = require('bluebird')

arn.db.ready.then(Promise.coroutine(function*() {
	let users = yield arn.db.all('Users')
	
	for(let user of users) {
		if(!user.ip || (user.location && user.location.countryName))
			continue
		
		console.log(`Querying ${chalk.blue(user.nick)} (${user.id} | ${user.ip})`)
		
		yield arn.getLocation(user).then(location => {
			user.location = location
			
			if(user.location && user.location.countryName)
				console.log(`${chalk.blue(user.nick)} from ${chalk.yellow(user.location.countryName)}`)
			
			return arn.db.set('Users', user.id, {
				location: user.location
			})
		}).catch(error => {
			user.location = null
		})
		
		yield Promise.delay(800)
	}
	
	console.log('Finished updating locations')
}))