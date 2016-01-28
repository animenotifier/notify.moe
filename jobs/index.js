'use strict'

global.arn = require('../lib')

let fs = require('fs')
let chalk = require('chalk')
let Promise = require('bluebird')

arn.db.ready.then(Promise.coroutine(function*() {
	arn.animeList = yield arn.filter('Anime', anime => true)

	console.log(arn.animeList.length + ' anime')

	fs.readdirSync('jobs').forEach(file => {
		if(file === 'index.js')
			return

		console.log(chalk.green('[Starting job]'), chalk.blue(file.replace('.js', '')))
		require('./' + file)
	})
}))