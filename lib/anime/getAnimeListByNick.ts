import * as arn from 'arn'

export function getAnimeListByNick(nick: string, clearCache: boolean): Promise<any> {
	return arn.getUserByNick(nick)
	.then(user => arn.getAnimeList(user, clearCache))
	.catch(error => {
		if(error.message === 'AEROSPIKE_ERR_RECORD_NOT_FOUND')
			return Promise.reject(`User '${nick}' not found`)

		if(error.message)
			return Promise.reject(error.message)

		return Promise.reject(error.toString())
	})
}