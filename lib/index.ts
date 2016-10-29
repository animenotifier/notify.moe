import { EventEmitter } from 'events'
import { Database } from './interfaces/Database'
import { User } from './interfaces/User'

const events = new EventEmitter()
const api = require('../security/api-keys.json')
const aerospike = require('aero-aerospike')
const db: Database = aerospike.client(require('../config.json').database)

const listProviders = {
	AniList: require('./services/AniList'),
	MyAnimeList: require('./services/MyAnimeList')
}

const animeProviders = {
	Nyaa: require('./services/Nyaa')
}

const airingDateProviders = {
	AniList: require('./services/AniList')
}

// Connect to DB
db.connect().then(() => console.log('Successfully connected to database!'))

// Anime title to ID
let titleToId = {}

const production = process.env.NODE_ENV === 'production'

// Export
export {
	events,
	api,
	db,
	listProviders,
	animeProviders,
	airingDateProviders,
	titleToId,
	production
}

export * from './auth'
export * from './changeNick'
export * from './getAnimeList'
export * from './getAnimeListByNick'
export * from './getLocation'
export * from './getLikeController'
export * from './getUnlikeController'
export * from './getAnimeIdBySimilarTitle'
export * from './getIdByTitle'
export * from './getUserByNick'
export * from './isActiveUser'
export * from './refreshAnimeList'
export * from './sendNotification'

export * from './autocorrect/fixGenre'
export * from './autocorrect/fixListProviderUserName'
export * from './autocorrect/fixNick'