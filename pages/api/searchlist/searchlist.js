exports.get = (request, response) => {
	if(!arn.animeToIdJSONStringGzipped) {
		response.writeHead(HTTP.NOT_FOUND)
		response.end()
		return
	}

	response.writeHead(HTTP.OK, {
		'Content-Type': 'application/json; charset=UTF-8',
		'Content-Encoding': 'gzip',
		'Content-Length': arn.animeToIdJSONStringGzipped.length
	})
	response.end(arn.animeToIdJSONStringGzipped)
}