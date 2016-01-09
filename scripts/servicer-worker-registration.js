// Check that service workers are supported, if so, progressively
// enhance and add push messaging support, otherwise continue without it.
if('serviceWorker' in navigator) {
	console.log('Registering service worker...');

	navigator.serviceWorker.register('/service-worker.js')
	.then(initialiseState)
	.catch(function(err) {
		console.error('ServiceWorker registration failed: ', err);
	});
} else {
	console.warn('Service workers aren\'t supported in this browser.');
}