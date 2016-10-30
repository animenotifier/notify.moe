import { User } from 'arn/interfaces/User'

export function isActiveUser(user: User): boolean {
	if(user.nick.startsWith('g'))
		return false

	if(user.nick.startsWith('fb'))
		return false

	if(user.nick.startsWith('t'))
		return false

	if(!user.lastView)
		return false

	let now = new Date().valueOf()
	let lastView = new Date(user.lastView.date).valueOf()

	if((now - lastView) > 14 * 24 * 60 * 60 * 1000)
		return false

	let listProviderName = user.providers.list

	if(!listProviderName)
		return false

	let listProvider = user.listProviders[listProviderName]

	if(!listProvider || !listProvider.userName)
		return false

	return true
}