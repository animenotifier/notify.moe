let dnode = require('dnode')

arn.chatBot = {
	sendMessage: function(channel, message) {
		if(!arn.production)
			return

		dnode.connect(app.config.ports.chatBot)
		.on('remote', remote => remote.sendMessage(channel, message))
		.on('fail', console.error)
		.on('error', console.error)
	}
}