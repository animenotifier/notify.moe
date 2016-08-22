let dnode = require('dnode')

app.on('config loaded', () => {
	let d = dnode.connect(app.config.ports.chatBot)
	
	d.on('error', function(error) {
		console.error(chalk.yellow('Failed connecting to chat bot!'))
	})

	d.on('remote', function(remote) {
		arn.chatBot = remote
	})
})