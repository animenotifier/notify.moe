// Log requests
/*if(!arn.production) {
	app.use(function(request, response, next) {
		let start = new Date()
		next()
		let end = new Date()

		if(request.user && request.user.nick)
			console.log(request.url, '|', end - start, 'ms', '|', request.user.nick)
		else
			console.log(request.url, '|', end - start, 'ms')
	})
}*/