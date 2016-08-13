window.newThread = function() {
	let text = $('post-input').value
	let tag = $('tag').value
	let title = $('new-thread-title').value
	
	if(!title || !text || !tag)
		return
	
	$.post('/api/threads', {
		title,
		text,
		tag
	}).then(threadId => {
		$.loadURL('/threads/' + threadId, true)
	}).catch(e => {
		console.error(e)
	})
}

$('new-thread-title').focus()