'use strict'

let request = require('request-promise')
let Promise = require('bluebird')
let arn = require('../lib')
let RateLimiter = require('limiter').RateLimiter
let limiter = new RateLimiter(1, 1100)
let database = require('../modules/database')
let aero = require('aero')

database(aero, function(error) {
	arn.listProviders.AniList.authorize().then(() => {
		let maxPage = 237
		for(let page = 1; page <= maxPage; page++) {
			limiter.removeTokens(1, function() {
				arn.listProviders.AniList.getAnimeFromPage(page).then(animeList => {
					let tasks = animeList.map(anime => arn.set('Anime', anime.id, anime))
					Promise.all(tasks).then(() => console.log('Finished importing page', page))
				})
			})
		}
	})
})