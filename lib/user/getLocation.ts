import { User } from 'lib/interfaces/User'
import { Location } from 'lib/interfaces/Location'
import * as request from 'request-promise'

export async function getLocation(user: User): Promise<Location> {
	let locationAPI = `http://api.ipinfodb.com/v3/ip-city/?key=${this.api.ipInfoDB.clientID}&ip=${user.ip}&format=json`
	let location: Location = await request(locationAPI).then(JSON.parse)
	return location
}