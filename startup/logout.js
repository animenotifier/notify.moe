// Logout
app.get('/logout', function(request, response) {
	request.logout()
	Promise.delay(500).then(() => {
		response.redirect('/')
	})
})