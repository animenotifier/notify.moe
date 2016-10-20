const serverId = '134910939140063232'

let Discord = require('discord.js')
let bot = new Discord.Client()
let dnode = require('dnode')

let server = null
let generalChannel = null

let nodeServer = dnode({
	sendMessage: function(channelName, message) {
		let channel = server.channels.get('name', channelName)
		bot.sendMessage(channel, message)
	}
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
			return bot.sendMessage(generalChannel, parameters)
			
		if(command === 'search')
			return bot.reply(message, 'https://www.google.com/search?q=site:notify.moe+' + encodeURIComponent(parameters))
			
		if(command === 'google')
			return bot.reply(message, 'https://www.google.com/search?q=' + encodeURIComponent(parameters))
			
		if(command === 'lmgtfy')
			return bot.reply(message, 'http://lmgtfy.com/?q=' + encodeURIComponent(parameters))
			
		if(command === 'user') {
			let lowerCaseUserName = parameters.toLowerCase()
			let users = yield arn.db.filter('Users', user => user.nick.toLowerCase() === lowerCaseUserName)
			
			if(users.length === 0)
				return bot.reply(message, 'That user d-doesn\'t exist, baaka!')
			else
				return bot.reply(message, `https://notify.moe/+${users[0].nick}`)
		}
			
		if(command === 'sounds')
			return bot.reply(message, '\n' + (yield fs.readdirAsync('bots/sounds')).map(file => file.replace(/\.(mp3|ogg)/, '')).sort((a, b) => parseInt(a) - parseInt(b)).join('\n'))
			
		if(command === 's') {
			let soundNumber = parseInt(parameters)
			let sounds = yield fs.readdirAsync('bots/sounds')
			let filtered = sounds.filter(sound => parseInt(sound) === soundNumber)
			
			if(filtered.length > 0) {
				return bot.voiceConnection.playFile(`bots/sounds/${filtered[0]}`, {}, (error, streamIntent) => {
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
			
			return arn.db.set('Cache', 'botCommands', botCommands).then(() => bot.reply(message, 'Registered commands:\n' + Object.keys(botCommands.replies)))
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
			
			return arn.db.set('Cache', 'botCommands', botCommands).then(() => bot.reply(message, 'Registered commands:\n' + Object.keys(botCommands.replies).join(', ')))
		}
		
		if(command === 'replies') {
			let botCommands = yield arn.db.get('Cache', 'botCommands').catch(error => {
				return {
					replies: {}
				}
			})
			
			return bot.reply(message, 'Registered commands:\n' + Object.keys(botCommands.replies).join(', '))
		}
		
		if(command === 'help') {
			let commands = [
				'**!addreply** [name] [message]',
				'**!google** [search term]',
				'**!removereply** [name]',
				'**!replies**',
				'**!s** [number of sound file]',
				'**!say** [message in general chat]',
				'**!search** [search term for notify.moe only]',
				'**!sounds**',
				'**!user** [name]'
			]
			return bot.reply(message, '\n' + commands.join('\n'))
		}
		
		// Custom replies
		let botCommands = yield arn.db.get('Cache', 'botCommands').catch(error => {
			return {
				replies: {}
			}
		})
		
		if(botCommands.replies[command])
			return bot.reply(message, botCommands.replies[command])
		
		return bot.reply(message, 'I d-don\'t understand what business you have with me!')
	}
	
	if(mentioned) {
		return bot.reply(message, 'B-Baka!')
	}
}))

bot.on('ready', () => {
	server = bot.servers.get('id', '134910939140063232')
	generalChannel = server.channels.get('id', '134910939140063232')
	
	console.log(chalk.green('Bot is ready'))
	
	for(let channel of server.channels) {
		if(channel.type === 'voice' && channel.name.endsWith('Talk')) {
			bot.joinVoiceChannel(channel, (error, voiceConnection) => {
				if(error)
					return console.error(error)
				
				console.log(chalk.green('Bot joined voice channel Talk'))
			})
			break
		}
	}
})

bot.loginWithToken(arn.api.discord.token, (error, token) => {
	console.log(chalk.green('Logged in'))
})