import * as arn from 'arn'
import * as gravatar from 'gravatar'
import * as chalk from 'chalk'
import * as bluebird from 'bluebird'

const striptags = require('striptags')

const sourceRegEx = /\(Source: (.*?)\)/i
const writtenByRegEx = /\[Written by (.*?)\]/i

export async function updateAnimePage(anime) {
	if(!isNaN(anime)) {
		anime = await arn.db.get('Anime', anime)
	}

	let providers: any = {}

	let createScanFunction = function(listProviderName) {
		providers[listProviderName] = []

		return match => {
			if(match.id === anime.id)
				providers[listProviderName].push(match)
		}
	}

	let userQueryTasks = new Array<Promise<any>>()

	let otherTasks = {
		HummingBird: arn.db.forEach('MatchHummingBird', createScanFunction('HummingBird')),
		MyAnimeList: arn.db.forEach('MatchMyAnimeList', createScanFunction('MyAnimeList')),
		AnimePlanet: arn.db.forEach('MatchAnimePlanet', createScanFunction('AnimePlanet')),
		Nyaa: arn.db.get('AnimeToNyaa', anime.id).catch(error => undefined),
		Watching: arn.db.forEach('AnimeLists', list => {
			if(!list.userId)
				return

			if(list.watching.find(entry => entry.id === anime.id)) {
				userQueryTasks.push(arn.db.get('Users', list.userId).then(user => {
					if(!arn.isActiveUser(user))
						return null

					return {
						id: user.id,
						nick: user.nick,
						avatar: user.avatar
					}
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
	} else {
		sourceMatch = anime.description.match(writtenByRegEx)

		if(sourceMatch !== null) {
			summarySource = sourceMatch[1]
			anime.description = anime.description.replace(writtenByRegEx, '').trim()
		}
	}

	let result: any = await bluebird.props(otherTasks)

	providers.MyAnimeList.sort(sortMatches)
	providers.HummingBird.sort(sortMatches)
	providers.AnimePlanet.sort(sortMatches)

	providers.MyAnimeList = providers.MyAnimeList[0]
	providers.HummingBird = providers.HummingBird[0]
	providers.AnimePlanet = providers.AnimePlanet[0]

	providers.Nyaa = result.Nyaa

	let usersWatching = (await bluebird.all(userQueryTasks)).filter(user => user !== null)
	let generated = (new Date()).toISOString()

	let animePage = {
		anime,
		providers,
		usersWatching,
		summarySource,
		generated
	}

	await bluebird.all([
		arn.db.set('Anime', anime.id, {
			watching: usersWatching.length,
			pageGenerated: generated
		}),
		arn.db.set('AnimePages', anime.id, animePage).then(() => {
			console.log(chalk.green('✔'), `Updated anime page ${chalk.cyan(anime.id)}`)
		})
	])
}