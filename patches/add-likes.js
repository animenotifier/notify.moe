var arn = require('../lib')

arn.db.ready.then(() => {
	let tasks = []
	
    arn.forEach('Threads', post => {
		tasks.push(arn.set('Threads', post.id, {
			likes: []
		}))
    }).then(function() {
		Promise.all(tasks).then(() => console.log(`Finished updating ${tasks.length} records`))
    })
})