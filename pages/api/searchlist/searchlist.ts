import * as arn from 'arn'
import * as HTTP from 'http-status-codes'

exports.get = async (request, response) => {
	try {
		let searchList = await arn.db.get('Cache', 'animeTitleToId')

		response.writeHead(HTTP.OK, {
			'Content-Type': 'application/json; charset=utf-8',
			'Content-Encoding': 'gzip',
			'Content-Length': searchList.compressed.length
		})

		response.end(searchList.compressed)
	} catch(error) {
		response.writeHead(HTTP.NOT_FOUND)
		response.end(error.toString())
	}
}