'use strict'

global.arn = require('../lib')
global.chalk = require('chalk')
global.Promise = require('bluebird')
global.fetch = require('request-promise')
global.fs = Promise.promisifyAll(require('fs'))
global.coroutine = Promise.coroutine
global.gravatar = require('gravatar')

// Time units
global.seconds = 1
global.minutes = 60 * seconds
global.hours = 60 * minutes

arn.runningBackgroundJobs = true

arn.db.ready.then(Promise.coroutine(function*() {
	arn.animeList = yield arn.filter('Anime', anime => true)

	console.log(arn.animeList.length + ' anime')

	let files = yield fs.readdirAsync('jobs')

	files.forEach(file => {
		if(file === 'index.js')
			return

		console.log(chalk.green('[Starting job]'), chalk.blue(file.replace('.js', '')))
		require('./' + file)
	})
}))