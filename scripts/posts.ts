import * as $ from './aqua'

function like(type, id) {
	$.post(`/api/${type}/like/` + id)
	$.find('like-' + id).style.display = 'none'
	$.find('unlike-' + id).style.display = 'inline-block'

	let likes = $.find('likes-' + id)
	likes.innerHTML = (parseInt(likes.textContent) + 1).toString()
}

function unlike(type, id) {
	$.post(`/api/${type}/unlike/` + id)
	$.find('like-' + id).style.display = 'inline-block'
	$.find('unlike-' + id).style.display = 'none'

	let likes = $.find('likes-' + id)
	likes.innerHTML = (parseInt(likes.textContent) - 1).toString()
}

function edit(id) {
	$.find('source-' + id).style.display = 'block'
	$.find('save-' + id).style.display = 'block'
	$.find('render-' + id).style.display = 'none'
	$.find('toolbar-' + id).style.display = 'none'
}

function cancelEdit(id) {
	$.find('source-' + id).style.display = 'none'
	$.find('save-' + id).style.display = 'none'
	$.find('render-' + id).style.display = ''
	$.find('toolbar-' + id).style.display = ''
}

function saveEdit(type, id) {
	let source = <HTMLInputElement> $.find('source-' + id)
	let text = source.value

	$.post(`/api/${type}/edit/` + id, {
		id,
		text
	}).then(response => {
		$.find('render-' + id).innerHTML = response
		cancelEdit(id)
	})
}