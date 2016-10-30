import * as arn from 'arn'
import { User } from 'arn/interfaces/User'

export async function getUserByNick(nick: string): Promise<User> {
	// Very old Android app requests
	if(nick.indexOf('&animeProvider=') !== -1)
		return Promise.reject('Old Android app request')

	let record = await arn.db.get('NickToUser', nick)
	return arn.db.get('Users', record.userId)
}