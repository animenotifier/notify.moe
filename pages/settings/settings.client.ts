var userName = $('nick').value;

makeSaveable('/api/users/me', function(key, value) {
	switch(key) {
		case 'nick':
			value = value.replace(/\s+/g, '');
			var oldPath = '/+' + userName;
			var newPath = '/+' + value;

			//window.history.pushState('', document.title, newPath);

			var links = $.queryAll('a');
			for(var l = 0; l < links.length; ++l) {
				var link = links[l];
				if(link.href.endsWith(oldPath))
					link.href = newPath;
			}

			break;
	}
});

// Push notifications
var pushButton = $.query('.push-button');

if(pushButton) {
	pushButton.addEventListener('click', function() {
		if(isPushEnabled) {
			unsubscribe();
		} else {
			subscribe();
		}
	});
} else {
	console.warn('Push button not found.');
}

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