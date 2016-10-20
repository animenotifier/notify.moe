import { Database } from './interfaces/Database'
import { User } from './interfaces/User'
import { EventEmitter } from 'events'
import { getLikeController } from './getLikeController'
import { getUnlikeController } from './getUnlikeController'
import { fixListProviderUserName } from './fixListProviderUserName'
import { isActiveUser } from './isActiveUser'

const aerospike = require('aero-aerospike')

class AnimeReleaseNotifier {
	fixListProviderUserName = fixListProviderUserName
	getLikeController = getLikeController
	getUnlikeController = getUnlikeController
	isActiveUser = isActiveUser
	events: EventEmitter
	api: any
	db: Database
	listProviders: any
	animeProviders: any

	constructor() {
		this.events = new EventEmitter()
		this.api = require('../security/api-keys.json')
		this.db = aerospike.client(require('../config.json').database)
		this.db.connect().then(() => console.log('Successfully connected to database!'))
		this.listProviders = {
			AniList: require('./services/AniList.js')
		}

		this.animeProviders = {
			Nyaa: require('./services/Nyaa.js')
		}
	}

	auth(req, res, role): boolean {
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

	getLocation(user: User) {
		let locationAPI = `http://api.ipinfodb.com/v3/ip-city/?key=${this.api.ipInfoDB.clientID}&ip=${user.ip}&format=json`
		return request(locationAPI).then(JSON.parse)
	}

	changeNick(user: User, newNick) {
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

	fixGenre(genre: string) {
		return genre.replace(/ /g, '').replace(/-/g, '').toLowerCase()
	}
}

module.exports = new AnimeReleaseNotifier()
export { AnimeReleaseNotifier }