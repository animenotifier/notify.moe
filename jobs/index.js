'use strict'

global.arn = require('../lib')

let fs = require('fs')
let chalk = require('chalk')

arn.db.ready.then(() => {
	fs.readdirSync('jobs').forEach(file => {
		if(file === 'index.js')
			return

		console.log(chalk.green('[Starting job]'), chalk.blue(file.replace('.js', '')))
		require('./' + file)
	})
})