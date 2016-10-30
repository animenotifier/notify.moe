import { User } from 'arn/interfaces/User'

const userNameTakenMessage = 'Username is already taken.'

export function changeNick(user: User, newNick: string): Promise<any> {
	let oldNick = user.nick

	if(oldNick === newNick)
		return Promise.resolve()

	return this.db.get('NickToUser', newNick).then(record => {
		return Promise.reject(userNameTakenMessage)
	}).catch(error => {
		if(error === userNameTakenMessage)
			return

		user.nick = newNick

		return Promise.all([
			this.db.remove('NickToUser', oldNick),
			this.db.set('NickToUser', newNick, { userId: user.id }),
			this.db.set('Users', user.id, user)
		])
	})
}