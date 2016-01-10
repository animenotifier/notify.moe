var images = document.querySelectorAll('.user-image');

for(var i = 0; i < images.length; ++i) {
	var element = images[i];
	var oldSource = element.src;
	
	element.onerror = function() {
		this.src = oldSource;
	};

	element.src = element.dataset.image;
}