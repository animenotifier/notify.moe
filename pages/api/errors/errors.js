'use strict'

let arn = require('../../../lib')

exports.get = (request, response) => {
	response.json(arn.animeProviders.Nyaa.errors)
}