updateAvatars()

window.sendMessage = function() {
	let userName = $('nick').textContent
	let postInput = $('post-input')
	
	if(!postInput.value)
		return
	
	$.post('/api/messages/' + userName, {
		text: postInput.value
	}).then(response => {
		console.log(response)
		postInput.value = ''
		
		if(window.loadMessages)
			window.loadMessages()
	})
}