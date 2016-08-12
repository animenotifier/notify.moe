let aero = require('aero')
let database = require('../modules/database')
let arn = require('../lib')

database(aero, function(error) {
	let tasks = []

    arn.forEach('Anime', function(anime) {
		tasks.push(arn.remove('Anime', anime.id))
    }).then(function() {
		console.log('Waiting...')
		Promise.all(tasks).then(() => console.log(`Finished deleting ${tasks.length} anime`))
    })
})