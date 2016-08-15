let arn = require('../lib')

arn.db.ready.then(() => {
	let tasks = []

    arn.forEach('Users', function(user) {
		if(user.ip !== '::ffff:127.0.0.1' && user.ip !== '127.0.0.1')
			return
		
		tasks.push(arn.set('Users', user.id, {
			ip: null
		}))
    }).then(function() {
		Promise.all(tasks).then(() => console.log(`Finished updating ${tasks.length} users`))
    })
})