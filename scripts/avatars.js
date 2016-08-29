function updateAvatars() {
	let images = document.querySelectorAll('.user-image')

	for(let i = 0; i < images.length; ++i) {
		let img = images[i]
		
		if(img.naturalWidth === 0) {
			img.onload = function() {
				this.style.opacity = 1.0
			}

			img.onerror = function() {
				this.src = '/images/elements/no-gravatar.svg'
				this.style.opacity = 1.0
			}
		} else {
			img.style.opacity = 1.0
		}
	}
	
	// Tooltips
	// let links = document.querySelectorAll('.user')
	// 
	// for(let i = 0; i < links.length; ++i) {
	// 	let link = links[i]
	// 	
	// 	link.classList.add('tooltip')
	// 	link.setAttribute('data-tooltip', link.title)
	// 	link.title = ''
	// }
}