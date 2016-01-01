'use strict'

let request = require('request-promise')
let plural = require('../../plural')
let datediff = require('../../datediff')
let apiKeys = require('../../../security/api-keys.json')

class HummingBird {
	constructor() {
		this.headers = {
			'User-Agent': 'Anime Release Notifier'
		}
	}

	getAnimeList(userName, callback) {
		let apiURL = `https://hummingbirdv1.p.mashape.com/users/${userName}/library?status=currently-watching`
	}
}