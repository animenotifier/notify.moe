'use strict'

let Promise = require('bluebird')
let gravatar = require('gravatar')
let striptags = require('striptags')
let chalk = require('chalk')

const sourceRegEx = /\(Source: (.*?)\)/i

arn.updateAnimePage = Promise.coroutine(function*(anime) {
	if(!isNaN(anime)) {
		anime = yield arn.get('Anime', anime)
	}

	let providers = {}

	let createScanFunction = function(listProviderName) {
		providers[listProviderName] = []

		return match => {
			if(match.id === anime.id)
				providers[listProviderName].push(match)
		}
	}

	let userQueryTasks = []

	let otherTasks = {
		HummingBird: arn.forEach('MatchHummingBird', createScanFunction('HummingBird')),
		MyAnimeList: arn.forEach('MatchMyAnimeList', createScanFunction('MyAnimeList')),
		AnimePlanet: arn.forEach('MatchAnimePlanet', createScanFunction('AnimePlanet')),
		Nyaa: arn.get('AnimeToNyaa', anime.id).catch(error => undefined),
		Watching: arn.forEach('AnimeLists', list => {
			if(!list.userId)
				return

			if(list.watching.find(entry => entry.id === anime.id)) {
				userQueryTasks.push(arn.get('Users', list.userId).then(user => {
					if(user.avatar)
						user.gravatarURL = user.avatar + '?s=50&r=x&d=mm'
					else
						user.gravatarURL = '/images/elements/no-gravatar.png'
					
					return user
				}))
			}
		})
	}

	let sortMatches = (a, b) => {
		if((a.edited && b.edited) || (!a.edited && !b.edited))
			return a.similarity < b.similarity ? 1 : -1

		if(a.edited)
			return -1

		// b edited
		return 1
	}

	anime.description = striptags(anime.description)

	let summarySource = ''
	let sourceMatch = anime.description.match(sourceRegEx)

	if(sourceMatch !== null) {
		summarySource = sourceMatch[1]
		anime.description = anime.description.replace(sourceRegEx, '').trim()
	}

	let result = yield otherTasks

	providers.MyAnimeList.sort(sortMatches)
	providers.HummingBird.sort(sortMatches)
	providers.AnimePlanet.sort(sortMatches)

	providers.MyAnimeList = providers.MyAnimeList[0]
	providers.HummingBird = providers.HummingBird[0]
	providers.AnimePlanet = providers.AnimePlanet[0]

	providers.Nyaa = result.Nyaa

	let usersWatching = yield userQueryTasks
	let generated = (new Date()).toISOString()

	let animePage = {
		anime,
		providers,
		usersWatching,
		summarySource,
		generated
	}

	yield [
		arn.set('Anime', anime.id, {
			watching: usersWatching.length,
			pageGenerated: generated
		}),
		arn.set('AnimePages', anime.id, animePage).then(() => {
			console.log(chalk.green('✔'), `Updated anime page ${chalk.cyan(anime.id)}`)
		})
	]
})