var arn = require('../lib')

arn.db.ready.then(() => {
	let tasks = []
	
    arn.forEach('Messages', post => {
		tasks.push(arn.set('Messages', post.id, {
			likes: []
		}))
    }).then(function() {
		Promise.all(tasks).then(() => console.log(`Finished updating ${tasks.length} records`))
    })
})