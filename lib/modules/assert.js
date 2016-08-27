arn.assertParams = (request, requiredParams) => {
	for(let param of requiredParams) {
		if(!request.body[param]) {
			response.writeHead(HTTP.BAD_REQUEST)
			response.end(`Missing parameter: ${param}`)
			return false
		}
	}
	
	return true
}