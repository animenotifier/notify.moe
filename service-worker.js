self.addEventListener('install', function(event) {
	console.log('Service worker installation...');
	// Perform install steps
});

self.addEventListener('push', function(event) {
	console.log('Received a push message', event);

	var title = 'Yay a message.';
	var body = 'We have received a push message.';
	var icon = '/images/characters/arn-waifu.png';
	var tag = 'simple-push-demo-notification-tag';

	event.waitUntil(
		self.registration.showNotification(title, {
			body: body,
			icon: icon,
			tag: tag
		})
	);
});

self.addEventListener('push', function(event) {
	// Since there is no payload data with the first version
	// of push messages, we'll grab some data from
	// an API and use it to populate a notification
	event.waitUntil(
		fetch('/api/notifications').then(function(response) {
			if(response.status !== 200) {
				// Either show a message to the user explaining the error
				// or enter a generic message and handle the
				// onnotificationclick event to direct the user to a web page
				console.log('Looks like there was a problem. Status Code: ' + response.status);
				throw new Error();
			}

			// Examine the text in the response
			return response.json().then(function(data) {
				if(data.error || !data.notification) {
					console.error('The API returned an error.', data.error);
					throw new Error();
				}

				var title = data.notification.title;
				var message = data.notification.message;
				var icon = data.notification.icon;
				var notificationTag = data.notification.tag;

				return self.registration.showNotification(title, {
					body: message,
					icon: icon,
					tag: notificationTag
				});
			});
		}).catch(function(err) {
			console.error('Unable to retrieve data', err);

			var title = 'An error occurred';
			var message = 'We were unable to get the information for this push message';
			var icon = '/images/characters/arn-waifu.png';
			var notificationTag = 'notification-error';

			return self.registration.showNotification(title, {
				body: message,
				icon: icon,
				tag: notificationTag
			});
		})
	);
});

self.addEventListener('notificationclick', function(event) {
	console.log('On notification click: ', event.notification.tag);
	// Android doesn't close the notification when you click on it
	// See: http://crbug.com/463146
	event.notification.close();

	// This looks to see if the current is already open and
	// focuses if it is
	event.waitUntil(
		clients.matchAll({
			type: 'window'
		}).then(function(clientList) {
			for(var i = 0; i < clientList.length; i++) {
				var client = clientList[i];
				if(client.url === '/' && 'focus' in client)
					return client.focus();
			}
			if(clients.openWindow) {
				return clients.openWindow('/');
			}
		})
	);
});