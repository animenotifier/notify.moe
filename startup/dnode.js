let dnode = require('dnode')

const reconnectTime = 5000

let reconnect = () => {
	let d = dnode.connect(app.config.ports.chatBot)
	
	d.on('remote', function(remote) {
		arn.chatBot = remote
		
		console.log(chalk.green('Connected to chat bot'))
	})
	
	d.on('error', function(error) {
		arn.chatBot = null
		Promise.delay(reconnectTime).then(reconnect)
	})
}

app.on('config loaded', reconnect)