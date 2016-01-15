'use strict'



exports.get = (request, response) => {
	response.json(arn.animeProviders.Nyaa.errors)
}