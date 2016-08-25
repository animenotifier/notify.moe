let dnode = require('dnode')

const reconnectTime = 5000

let reconnect = () => {
	let d = dnode.connect(app.config.ports.chatBot)
	
	d.on('remote', function(remote) {
		arn.chatBot = remote
		
		console.log(chalk.green('Connected to chat bot'))
	})
	
	let onError = error => {
		arn.chatBot = null
		Promise.delay(reconnectTime).then(reconnect)
	}
	
	d.on('fail', onError)
	d.on('error', onError)
}

app.on('config loaded', reconnect)