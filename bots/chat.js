let Discord = require('discord.js')
let bot = new Discord.Client()
let generalChannel = null

bot.on('message', Promise.coroutine(function*(message) {
	console.log(message.cleanContent)
	
	let mentioned = message.isMentioned(bot.user)
	
	if(message.content.startsWith('!')) {
		let command = message.content.substring(1).split(' ')[0]
		let parameters = message.content.substring(command.length + 2)
		
		if(command === 'user') {
			let lowerCaseUserName = parameters.toLowerCase()
			let users = yield arn.filter('Users', user => user.nick.toLowerCase() === lowerCaseUserName)
			
			if(users.length === 0)
				return bot.reply(message, 'That user d-doesn\'t exist, baaka!')
			else
				return bot.reply(message, `https://notify.moe/+${users[0].nick}`)
		}
		
		if(command === 'say')
			return bot.sendMessage(generalChannel, parameters)
		
		if(command === 'rin')
			return bot.reply(message, 'http://pa1.narvii.com/5930/db735965b205ff5fa6783ae8aa3be0ff16766b2d_hq.gif')
		
		return bot.reply(message, 'I d-don\'t understand what business you have with me!')
	}
	
	if(mentioned) {
		return bot.reply(message, 'B-Baka!')
	}
}))

bot.on('ready', () => {
	let server = bot.servers.get('id', '134910939140063232')
	generalChannel = server.channels.get('id', '134910939140063232')
	
	console.log(chalk.green('Bot is ready'))
})

bot.loginWithToken(arn.apiKeys.discord.token, (error, token) => {
	console.log(chalk.green('Logged in'))
})