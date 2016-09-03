self.addEventListener('install', function(event) {
	self.skipWaiting();
	console.log('Installed', event);
});

self.addEventListener('activate', function(event) {
	console.log('Activated', event);
});

self.addEventListener('push', function(event) {
	console.log('Received push event:', event);

	// Since there is no payload data with the first version
	// of push messages, we'll grab some data from
	// an API and use it to populate a notification
	event.waitUntil(
		fetch('https://notify.moe/api/notifications', {
			credentials: 'same-origin'
		}).then(function(response) {
			if(response.status !== 200) {
				// Either show a message to the user explaining the error
				// or enter a generic message and handle the
				// onnotificationclick event to direct the user to a web page
				console.log('Looks like there was a problem. Status Code: ' + response.status);
				throw new Error();
			}

			// Examine the text in the response
			return response.json().then(function(data) {
				if(data.error || !data.notifications) {
					console.error('The API returned an error.', data.error);
					throw new Error();
				}

				var tasks = data.notifications.map(function(notification) {
					return self.registration.showNotification(notification.title, {
						body: notification.body,
						icon: notification.icon,
						badge: '/images/elements/badge.png',
						tag: notification.tag
					});
				});

				if(tasks.length > 0)
					return tasks[0];
			});
		}).catch(function(err) {
			console.error('Unable to retrieve data', err);

			return self.registration.showNotification('An error occurred', {
				body: 'We were unable to get the information for this push message',
				icon: '/images/characters/arn-waifu.png',
				badge: '/images/elements/badge.png',
				tag: 'notification-error'
			});
		})
	);
});

self.addEventListener('notificationclick', function(event) {
	console.log('On notification click: ', event.notification);
	// Android doesn't close the notification when you click on it
	// See: http://crbug.com/463146
	if(event.notification.close)
		event.notification.close();

	var url = '/+';

	// This looks to see if the current is already open and
	// focuses if it is
	event.waitUntil(
		clients.matchAll({
			type: 'window'
		}).then(function(clientList) {
			for(var i = 0; i < clientList.length; i++) {
				var client = clientList[i];
				if(client.url === url && 'focus' in client)
					return client.focus();
			}

			if(clients.openWindow) {
				return clients.openWindow(url);
			}
		})
	);
});