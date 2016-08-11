// Authorize AniList
arn.listProviders.AniList.authorize().then(accessToken => {
	console.log('AniList API token:', accessToken)
})