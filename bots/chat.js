const guildId = '134910939140063232'

const commands = [
	'**!addreply** [name] [message]',
	'**!g** [google search term]',
	'**!removereply** [name]',
	'**!replies**',
	'**!s** [number of sound file]',
	'**!say** [message in general chat]',
	'**!search** [search term for notify.moe only]',
	'**!sounds**',
	'**!user** [name]'
]

let Discord = require('discord.js')
let bot = new Discord.Client()
let dnode = require('dnode')

let sendMessage = function(channelName, message) {
	let channel = bot.channels.find('name', channelName)

	if(channel)
		channel.sendMessage(message)
	else
		console.error('Channel', channelName, 'not found.')
}

let nodeServer = dnode({
	sendMessage
})

nodeServer.listen(require('../config.json').ports.chatBot)

bot.on('message', Promise.coroutine(function*(message) {
	console.log(chalk.yellow(message.author.username), message.cleanContent)

	// Ignore own messages
	if(message.author.id === bot.user.id)
		return

	let mentioned = message.isMentioned(bot.user)

	if(message.content.startsWith('!')) {
		let command = message.content.substring(1).split(' ')[0]
		let parameters = message.content.substring(command.length + 2)

		if(command === 'say')
			return sendMessage('general', parameters)

		if(command === 'search')
			return message.reply('https://www.google.com/search?q=site:notify.moe+' + encodeURIComponent(parameters))

		if(command === 'g')
			return message.reply('https://www.google.com/search?q=' + encodeURIComponent(parameters))

		if(command === 'lmgtfy')
			return message.reply('http://lmgtfy.com/?q=' + encodeURIComponent(parameters))

		if(command === 'user') {
			let lowerCaseUserName = parameters.toLowerCase()
			let users = yield arn.db.filter('Users', user => user.nick.toLowerCase() === lowerCaseUserName)

			if(users.length === 0)
				return message.reply('That user d-doesn\'t exist, baaka!')
			else
				return message.reply(`https://notify.moe/+${users[0].nick}`)
		}

		if(command === 'sounds')
			return message.reply('\n' + (yield fs.readdirAsync('bots/sounds')).map(file => file.replace(/\.(mp3|ogg)/, '')).sort((a, b) => parseInt(a) - parseInt(b)).join('\n'))

		if(command === 's') {
			let soundNumber = parseInt(parameters)
			let sounds = yield fs.readdirAsync('bots/sounds')
			let filtered = sounds.filter(sound => parseInt(sound) === soundNumber)

			if(filtered.length > 0) {
				console.log(`bots/sounds/${filtered[0]}`)
				return bot.voiceConnections.get(guildId).playFile(`bots/sounds/${filtered[0]}`, {}, (error, streamIntent) => {
					if(error)
						console.error(error)
				})
			}
		}

		if(command === 'addreply') {
			let name = parameters.split(' ')[0]
			let reply = parameters.substring(name.length + 1)

			if(!name || !reply)
				return

			console.log('Adding reply:', name, reply)

			// Limit reply length
			if(reply.length > 512)
				return

			let botCommands = yield arn.db.get('Cache', 'botCommands').catch(error => {
				return {
					replies: {}
				}
			})

			botCommands.replies[name] = reply

			return arn.db.set('Cache', 'botCommands', botCommands).then(() => message.reply('Registered commands:\n' + Object.keys(botCommands.replies)))
		}

		if(command === 'removereply') {
			let name = parameters

			if(!name)
				return

			let botCommands = yield arn.db.get('Cache', 'botCommands').catch(error => {
				return {
					replies: {}
				}
			})

			delete botCommands.replies[name]

			return arn.db.set('Cache', 'botCommands', botCommands).then(() => message.reply('Registered commands:\n' + Object.keys(botCommands.replies).join(', ')))
		}

		if(command === 'replies') {
			let botCommands = yield arn.db.get('Cache', 'botCommands').catch(error => {
				return {
					replies: {}
				}
			})

			return message.reply('Registered commands:\n' + Object.keys(botCommands.replies).join(', '))
		}

		if(command === 'help') {
			let help = '\n' + commands.join('\n')
			return message.reply(help)
		}

		// Custom replies
		let botCommands = yield arn.db.get('Cache', 'botCommands').catch(error => {
			return {
				replies: {}
			}
		})

		if(botCommands.replies[command])
			return message.reply(botCommands.replies[command])

		return message.reply('I d-don\'t understand what business you have with me!')
	}

	if(mentioned) {
		return message.reply('B-Baka!')
	}
}))

bot.on('ready', () => {
	console.log(chalk.green('Bot is ready'))

	for(let channel of bot.channels.values()) {
		if(!channel.name)
			continue

		if(channel.type === 'voice' && channel.name.endsWith('Talk')) {
			channel.join()
			.then(() => console.log(chalk.green('Bot joined voice channel Talk')))
			.catch(console.error)
			break
		}
	}
})

bot.login(arn.api.discord.token).then(token => {
	console.log(chalk.green('Logged in'))
	sendMessage('log', 'I\'m online.')
})