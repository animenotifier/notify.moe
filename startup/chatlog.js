if(!arn.production)
	return

// Server start
app.on('server started', _ => {
	arn.chatBot.sendMessage('log', 'Web server online.')
})

// New users
arn.on('new user', user => {
	let infos = [
		`New user: https://notify.moe/+${user.id}`,
		`Name: **${user.firstName} ${user.lastName}**`,
		`Email: **${user.email}**`,
		`Via: **${Object.keys(user.accounts)[0]}**`
	]

	arn.chatBot.sendMessage('log', infos.join('\n'))
})