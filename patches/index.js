'use strict'

let aero = require('aero')
let database = require('../modules/database')
let arn = require('../lib')

database(aero, function(error) {
	//arn.remove('NickToUser', 'Sebastian ').then(() => {console.log('OK')})
	//arn.set('NickToUser', 'Sebastian', {userId: 'Ny9pwwZvg'})
	/*arn.getUserByNickAsync('Aky').then(user => {
		user.role = 'admin'

		arn.setUserAsync(user.id, user).then(() => {
				console.log('Finished updating ' + user.nick)
		})
	})*/

	arn.scan('Users', function(user) {
		user.sortBy = 'airingDate'
		arn.setUser(user.id, user)
		/*arn.getLocation(user).then(location => {
			user.location = location
			console.log(user.location)

			arn.setUser(user.id, user)
		})*/
	}, function() {
		console.log('Finished updating all users')
	})
})