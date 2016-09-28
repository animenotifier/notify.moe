window.replyToThread = function(threadId) {
	let text = $('post-input').value

	if(!text)
		return

	$.post('/api/posts', {
		text,
		threadId
	}).then(response => {
		$.load(window.location.pathname)
	}).catch(e => {
		console.error(e)
	})
}