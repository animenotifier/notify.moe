let dnode = require('dnode')

arn.chatBot = {
	sendMessage: function(channel, message) {
		dnode.connect(app.config.ports.chatBot)
		.on('remote', remote => remote.sendMessage(channel, message))
		.on('fail', console.error)
		.on('error', console.error)
	}
}