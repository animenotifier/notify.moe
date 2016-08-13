window.like = (type, id) => {
	$.post(`/api/${type}/like/` + id)
	$('like-' + id).style.display = 'none'
	$('unlike-' + id).style.display = 'inline-block'
	
	let likes = $('likes-' + id)
	likes.innerHTML = parseInt(likes.textContent) + 1
}

window.unlike = (type, id) => {
	$.post(`/api/${type}/unlike/` + id)
	$('like-' + id).style.display = 'inline-block'
	$('unlike-' + id).style.display = 'none'
	
	let likes = $('likes-' + id)
	likes.innerHTML = parseInt(likes.textContent) - 1
}

window.edit = id => {
	$('source-' + id).style.display = 'block'
	$('save-' + id).style.display = 'block'
	$('render-' + id).style.display = 'none'
	$('toolbar-' + id).style.display = 'none'
}

window.cancelEdit = id => {
	$('source-' + id).style.display = 'none'
	$('save-' + id).style.display = 'none'
	$('render-' + id).style.display = ''
	$('toolbar-' + id).style.display = ''
}

window.saveEdit = (type, id) => {
	let text = $('source-' + id).value
	
	$.post(`/api/${type}/edit/` + id, {
		id,
		text
	}).then(response => {
		$('render-' + id).innerHTML = response
		cancelEdit(id)
	})
}