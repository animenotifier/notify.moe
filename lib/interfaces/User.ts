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
	ip: string,
	lastView?: {
		date: string
	}
}