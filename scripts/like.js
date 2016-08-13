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