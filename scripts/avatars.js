function updateAvatars() {
	let images = document.querySelectorAll('.user-image')

	for(var i = 0; i < images.length; ++i) {
		let element = images[i]
		
		element.onload = function() {
			this.style.opacity = 1.0
		}

		element.onerror = function() {
			this.src = '/images/elements/no-gravatar.png'
			this.style.opacity = 1.0
		}
	}
}