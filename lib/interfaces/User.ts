import { Location } from './Location'

export interface User {
	id: string
	nick: string
	role: string
	firstName: string
	lastName: string
	email: string
	gender: string
	language: string
	ageRange?: any,
	accounts: any
	tagline: string
	website: string
	providers: {
		list: string
		anime: string
		airingDate: string
	},
	listProviders: {
		AniList?: {
			userName: string
		}
		MyAnimeList?: {
			userName: string
		}
		HummingBird?: {
			userName: string
		}
		AnimePlanet?: {
			userName: string
		}
	}
	sortBy: string
	titleLanguage: string
	pushEndpoints: Map<string, any>
	following: Array<string>
	registered: string
	lastLogin: string
	ip?: string
	location?: Location
	avatar: string
	lastView?: {
		date: string
	}
}