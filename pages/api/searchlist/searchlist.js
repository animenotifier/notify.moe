exports.get = (request, response) => {
	if(!arn.animeToIdJSONStringGzipped) {
		response.writeHead(404)
		response.end()
		return
	}

	response.writeHead(200, {
		'Content-Type': 'application/json; charset=UTF-8',
		'Content-Encoding': 'gzip',
		'Content-Length': arn.animeToIdJSONStringGzipped.length
	})
	response.end(arn.animeToIdJSONStringGzipped)
}