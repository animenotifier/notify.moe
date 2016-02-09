/*$("#container").animate({
	'opacity' : '1'
}, 290, 'linear');
$("#title").animate({
	'opacity' : '1',
	'letter-spacing' : '6px'
}, 290, 'linear');*/

// Fix Facebook login hash in URL
if(window.location.hash && window.location.hash === '#_=_') {
	window.history.pushState('', document.title, window.location.pathname);
}

// Fade out loading animation
document.addEventListener('DOMContentLoaded', function(event) {
	aero.loadingAnimation.classList.add('fade-out');
});