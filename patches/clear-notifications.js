var arn = require('../lib')

arn.db.ready.then(() => {
	let tasks = []

    arn.forEach('Users', function(user) {
		tasks.push(arn.remove('Notifications', user.id).catch(error => null))
    }).then(function() {
		console.log('Waiting...')
		Promise.all(tasks).then(() => console.log(`Finished deleting ${tasks.length} notification lists`))
    })
})