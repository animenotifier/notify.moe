let aerospike = require('aero-aerospike')

arn.db = aerospike.client(require('../../config.json').database)

arn.db.connect().then(() => console.log('Successfully connected to database!'))

arn.addProperties = (set, properties) => {
	let tasks = []

	arn.db.forEach(set, entry => {
		tasks.push(arn.db.set(set, entry.id, properties))
	})
	.then(() => Promise.all(tasks))
	.then(() => console.log(`Added properties to ${tasks.length} records`))
}