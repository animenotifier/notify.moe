window.gravatarAvailable = function(available) {
	var gravatarText = $('gravatar-text')
	gravatarText.innerHTML = 'Add a gravatar image'
	gravatarText.className = available ? 'finished' : 'not-finished'
}