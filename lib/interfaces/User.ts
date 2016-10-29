import { Location } from './Location'

export interface User {
	id: string
	nick: string
	providers: {
		anime: string
		list: string
		airingDate: string
	}
	listProviders: {
		AniList?: {
			userName: string
		}
	},
	sortBy: string
	titleLanguage: string
	location: Location
	pushEndpoints: Map<string, any>
	ip: string
	lastView?: {
		date: string
	}
}