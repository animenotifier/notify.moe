function updateAvatars() {
	let images = document.querySelectorAll('.user-image')

	for(var i = 0; i < images.length; ++i) {
		let img = images[i]
		
		if(img.naturalWidth === 0) {
			img.onload = function() {
				this.style.opacity = 1.0
			}

			img.onerror = function() {
				this.src = '/images/elements/no-gravatar.png'
				this.style.opacity = 1.0
			}
		} else {
			img.style.opacity = 1.0
		}
	}
}