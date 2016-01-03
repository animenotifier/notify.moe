'use strict'

let aero = require('aero')
let database = require('../modules/database')
let arn = require('../lib')

database(aero, function(error) {
        //arn.removeAsync('NickToUser', 'Sebastian ').then(() => {console.log('OK')})
        //arn.set('NickToUser', 'Sebastian', {userId: 'Ny9pwwZvg'})
        /*arn.getUserByNickAsync('Aky').then(user => {
                user.role = 'admin'

                arn.setUserAsync(user.id, user).then(() => {
                        console.log('Finished updating ' + user.nick)
                })
        })*/

        arn.scan('Users', function(user) {
				let old = user.nick
				let newNick = arn.fixNick(user.nick)

				console.log(old, '=', newNick)

                arn.changeNickAsync(user, newNick).then(() => {
                    console.log('Finished updating ' + user.nick)
                }).catch(error => {
					console.error('Error updating ' + user.nick)
				})
        }, function() {
                console.log('Finished updating all users')
        })
})
