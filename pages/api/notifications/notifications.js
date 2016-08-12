exports.get = (request, response) => {
	let user = request.user

	if(!user) {
		response.json({
			notifications: [{
				title: 'You have new notifications',
				message: 'Log in on notify.moe to view them',
				icon: '/images/characters/arn-waifu.png',
				tag: 'not-logged-in'
			}]
		})
		return
	}

	arn.get('Notifications', user.id).then(record => {
		console.log(`Service worker retrieved notifications of ${user.nick}:`, record.notifications)

		arn.remove('Notifications', user.id).then(() => {
			response.json({
				notifications: record.notifications
			})
		})
	}).catch(error => {
		if(error.message === 'AEROSPIKE_ERR_RECORD_NOT_FOUND') {
			response.json({
				notifications: []
			})
			return error
		}

		console.error(error, error.stack)

		response.json({
			notifications: [{
				title: 'Error fetching notifications',
				message: 'Open notify.moe to view them',
				icon: '/images/characters/arn-waifu.png',
				tag: 'error'
			}]
		})
	})
}