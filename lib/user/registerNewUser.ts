import * as arn from 'arn'
import * as Promise from 'bluebird'
import * as chalk from 'chalk'
import * as shortid from 'shortid'
import * as gravatar from 'gravatar'
import { User } from 'arn/interfaces/User'

export function registerNewUser(userData: User) {
	let now = new Date()
	let user: User = {
		id: shortid.generate(),
		nick: '',
		role: '',
		firstName: '',
		lastName: '',
		email: '',
		gender: '',
		language: '',
		ageRange: null,
		accounts: {},
		tagline: '',
		website: '',
		providers: {
			list: 'AniList',
			anime: 'CrunchyRoll',
			airingDate: 'AniList'
		},
		listProviders: {},
		sortBy: 'airingDate',
		titleLanguage: 'romaji',
		pushEndpoints: <Map<string, any>> {},
		following: [],
		registered: now.toISOString(),
		lastLogin: now.toISOString(),
		avatar: ''
	}

	// Assign provider specific data from Google, Facebook, Twitter...
	Object.assign(user, userData)

	if(user.email)
		user.avatar = gravatar.url(user.email)
	else
		user.avatar = '/images/elements/no-gravatar.svg'

	if(user.email === 'e.urbach@gmail.com')
		user.role = 'admin'

	console.log(chalk.green('New user:'), user)

	let tasks = [
		arn.db.set('NickToUser', user.nick, { userId: user.id })
	]

	if(user.email)
		tasks.push(arn.db.set('EmailToUser', user.email, { userId: user.id }))

	return Promise.all(tasks).then(function() {
		arn.events.emit('new user', user)
		return user
	})
}
