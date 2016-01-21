'use strict'

exports.get = (request, response) => {
	response.writeHead(200, {
		'Content-Type': 'application/json; charset=UTF-8',
		'Content-Encoding': 'gzip',
		'Content-Length': arn.animeToIdJSONStringGzipped.length
	})
	response.end(arn.animeToIdJSONStringGzipped)
}