let aerospike = require('aero-aerospike')

arn.db = aerospike.client(require('../../config.json').database)

arn.db.connect().then(() => console.log('Successfully connected to database!'))

arn.get = arn.db.get
arn.set = arn.db.set
arn.remove = arn.db.remove
arn.forEach = arn.db.forEach
arn.filter = arn.db.filter
arn.all = arn.db.all
arn.batchGet = arn.db.getMany

arn.addProperties = (set, properties) => {
	let tasks = []
	
	arn.forEach(set, entry => {
		tasks.push(arn.set(set, entry.id, properties))
	})
	.then(() => Promise.all(tasks))
	.then(() => console.log(`Added properties to ${tasks.length} records`))
}