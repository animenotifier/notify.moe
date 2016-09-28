// Authorize AniList
arn.listProviders.AniList.authorize().then(accessToken => {
	app.ready.then(() => console.log('AniList API token:', accessToken))
})