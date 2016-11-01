import * as arn from 'arn'

export function addProperties(set: string, properties: Object) {
	let tasks = new Array<Promise<any>>()

	arn.db.forEach(set, entry => {
		tasks.push(arn.db.set(set, entry.id, properties))
	})
	.then(() => Promise.all(tasks))
	.then(() => console.log(`Added properties to ${tasks.length} records`))
}