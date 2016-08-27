global.arn = require('../lib')
global.chalk = require('chalk')
global.Promise = require('bluebird')
global.fs = Promise.promisifyAll(require('fs'))

arn.db.scanPriority = require('aerospike').scanPriority.LOW

arn.db.ready.then(Promise.coroutine(function*() {
	let files = yield fs.readdirAsync('bots')
	let filterJob = process.argv[2]

	files.forEach(file => {
		if(file === 'index.js' || file === 'sounds')
			return
		
		if(filterJob && file !== filterJob + '.js')
			return

		console.log(chalk.green('[Starting bot]'), chalk.blue(file.replace('.js', '')))
		require('./' + file)
	})
}))