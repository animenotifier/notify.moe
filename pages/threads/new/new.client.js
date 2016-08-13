window.newThread = function() {
	let text = $('post-input').value
	let tag = $('tag').value
	let title = $('new-thread-title').value
	let stickyBox = $('sticky')
	
	if(!title || !text || !tag)
		return
	
	$.post('/api/threads', {
		title,
		text,
		tag,
		sticky: (stickyBox && stickyBox.checked) ? 1 : 0
	}).then(threadId => {
		$.loadURL('/threads/' + threadId, true)
	}).catch(e => {
		console.error(e)
	})
}

$('new-thread-title').focus()