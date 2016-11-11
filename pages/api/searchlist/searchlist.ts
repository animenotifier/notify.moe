import * as arn from 'arn'
import * as HTTP from 'http-status-codes'

exports.get = (request, response) => {
	if(!arn.titleToIdJSONStringGzipped) {
		response.writeHead(HTTP.NOT_FOUND)
		response.end()
		return
	}

	response.writeHead(HTTP.OK, {
		'Content-Type': 'application/json; charset=UTF-8',
		'Content-Encoding': 'gzip',
		'Content-Length': arn.titleToIdJSONStringGzipped.length
	})
	response.end(arn.titleToIdJSONStringGzipped)
}