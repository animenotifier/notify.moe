'use strict'

let aero = require('aero')
let database = require('../modules/database')
let arn = require('../lib')

database(aero, function(error) {
    arn.scan('Users', function(user) {
        if(user.listProviders.AniList && user.listProviders.AniList.userName) {
            let userName = user.listProviders.AniList.userName
            if(userName.startsWith('http://anilist.co/animelist/')) {
                user.listProviders.AniList.userName = userName.substring('http://anilist.co/animelist/'.length)
                console.log(user.listProviders.AniList.userName)

                arn.setUser(user.id, user)
            }
        }
    }, function() {
        console.log('Finished updating all users')
    })
})