global.arn = require('../lib')

let patch = process.argv[2]

if(!patch) {
	console.error('Patch name not specified')
	process.exit()
}

require('./' + patch)