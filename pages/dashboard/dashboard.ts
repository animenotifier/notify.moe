import * as arn from 'arn'
import * as gravatar from 'gravatar'

exports.get = function(request, response) {
	let user = request.user

	if(user) {
		let listProviderSettings = user.listProviders[user.providers.list]
		if(!listProviderSettings || !listProviderSettings.userName)
			user.hasListProviderUserName = false
		else
			user.hasListProviderUserName = true

		if(user.email)
			user.gravatarURL = gravatar.url(user.email, {s: '1', r: 'x', d: '404'}, true)

		user.hasNick = !(user.nick.startsWith('fb') || user.nick.startsWith('g') || user.nick.startsWith('t'))
		user.isActive = arn.isActiveUser(user)

		if(user.hasNick)
			user.welcomeLine = 'Hi ' + user.nick + '!'
		else if(user.firstName)
			user.welcomeLine = 'Hi ' + user.firstName + '!'
		else
			user.welcomeLine = 'Hi!'
	}

	response.render({
		user
	})
}