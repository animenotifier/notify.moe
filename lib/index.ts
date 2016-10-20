import { Database } from './interfaces/Database'
import { User } from './interfaces/User'
import { EventEmitter } from 'events'
import { getLikeController } from './getLikeController'
import { getUnlikeController } from './getUnlikeController'
import { fixListProviderUserName } from './fixListProviderUserName'
import { isActiveUser } from './isActiveUser'
import { getAnimeIdBySimilarTitle } from './getAnimeIdBySimilarTitle'
import { getIdByTitle } from './getIdByTitle'

const aerospike = require('aero-aerospike')

const events = new EventEmitter()
const api = require('../security/api-keys.json')
const db: Database = aerospike.client(require('../config.json').database)
const listProviders = {
	AniList: require('./services/AniList.js')
}

const animeProviders = {
	Nyaa: require('./services/Nyaa.js')
}

function auth(req, res, role): boolean {
	if(!req.user) {
		res.end('Not logged in!')
		return false
	}

	if(req.user.role !== 'admin' && req.user.role !== role) {
		res.end('Not authorized to view this page!')
		return false
	}

	return true
}

function getLocation(user: User) {
	let locationAPI = `http://api.ipinfodb.com/v3/ip-city/?key=${this.api.ipInfoDB.clientID}&ip=${user.ip}&format=json`
	return request(locationAPI).then(JSON.parse)
}

function changeNick(user: User, newNick) {
	const userNameTakenMessage = 'Username is already taken.'
	let oldNick = user.nick

	if(oldNick === newNick)
		return Promise.resolve()

	return this.db.get('NickToUser', newNick).then(record => {
		return Promise.reject(userNameTakenMessage)
	}).catch(error => {
		if(error === userNameTakenMessage)
			return

		user.nick = newNick

		return Promise.all([
			this.db.remove('NickToUser', oldNick),
			this.db.set('NickToUser', newNick, { userId: user.id }),
			this.db.set('Users', user.id, user)
		])
	})
}

function fixGenre(genre: string) {
	return genre.replace(/ /g, '').replace(/-/g, '').toLowerCase()
}

// Connect to DB
db.connect().then(() => console.log('Successfully connected to database!'))

// Export
export {
	events,
	api,
	db,
	listProviders,
	animeProviders,
	fixListProviderUserName,
	getLikeController,
	getUnlikeController,
	isActiveUser,
	getAnimeIdBySimilarTitle,
	getIdByTitle,
	auth,
	getLocation,
	changeNick,
	fixGenre
}