app.on('database ready', db => {
	db.addProperties = function(set, properties) {
		let tasks = []
		
		this.forEach(set, entry => {
			tasks.push(db.set(set, entry.id, properties))
	    })
		.then(() => Promise.all(tasks))
		.then(() => console.log(`Added properties to ${tasks.length} records`))
	}
})