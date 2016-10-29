const autoCorrectUserNames = [
	/anilist.co\/user\/(.*)/,
	/anilist.co\/animelist\/(.*)/,
	/hummingbird.me\/users\/(.*?)\/library/,
	/hummingbird.me\/users\/(.*)/,
	/anime-planet.com\/users\/(.*?)\/anime/,
	/anime-planet.com\/users\/(.*)/,
	/myanimelist.net\/profile\/(.*)/,
	/myanimelist.net\/animelist\/(.*?)\?/,
	/myanimelist.net\/animelist\/(.*)/,
	/myanimelist.net\/(.*)/,
	/myanimelist.com\/(.*)/,
	/twitter.com\/(.*)/,
	/osu.ppy.sh\/u\/(.*)/
]

export function fixListProviderUserName(userName: string): string {
	userName = userName.trim()

	for(let regex of autoCorrectUserNames) {
		let match = regex.exec(userName)

		if(match !== null) {
			userName = match[1]
			break
		}
	}

	return userName
}