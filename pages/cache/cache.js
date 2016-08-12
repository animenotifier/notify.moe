exports.get = (request, response) => {
	let cache = arn.animeProviders.Nyaa.cache
	cache.keys((err, keys) => {
		response.render({
			keys,
			stats: cache.getStats()
		})
	})
}