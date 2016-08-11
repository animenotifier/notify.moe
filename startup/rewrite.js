// Rewrite URLs
app.rewrite(function(request, response) {
	if(request.headers.host.indexOf('animereleasenotifier.com') !== -1) {
        response.redirect('https://notify.moe' + request.url)
        return true
    }

	if(request.url.startsWith('/+'))
		request.url = '/user/' + request.url.substring(2)
	else if(request.url.startsWith('/_/+'))
		request.url = '/_/user/' + request.url.substring(4)
})