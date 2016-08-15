let arn = require('../lib')
let chalk = require('chalk')

arn.db.ready.then(() => {
	let tasks = []

    arn.forEach('Users', user => {
		if(!user.ip)
			return
		
		tasks.push(arn.getLocation(user).then(location => {
			user.location = location
			
			if(user.location && user.location.countryName)
				console.log(`${chalk.blue(user.nick)} from ${chalk.yellow(user.location.countryName)}`)
		}).catch(error => {
			user.location = null
		}).finally(() => {
			// Save in database
			return arn.set('Users', user.id, {
				location: user.location
			})
		}))
    }).then(() => {
		console.log(`Waiting for ${tasks.length} tasks to finish...`)
		Promise.all(tasks).then(() => console.log(`Finished updating ${tasks.length} users`))
    })
})