// Fix Facebook login hash in URL
if(window.location.hash && window.location.hash === '#_=_') {
	window.history.pushState('', document.title, window.location.pathname)
}

// Fade out loading animation
document.addEventListener('DOMContentLoaded', function(event) {
	$.loadingAnimation.classList.add('fade-out')
})