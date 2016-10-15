self.addEventListener('install', function(event) {
	self.skipWaiting()
	console.log('Installed', event)
})

self.addEventListener('activate', function(event) {
	console.log('Activated', event)
})

self.addEventListener('push', function(event) {
	console.log('Received push event:', event)

	if(event.data) {
		let notification = event.data.json()

		event.waitUntil(
			self.registration.showNotification(notification.title, {
				body: notification.body,
				icon: notification.icon,
				badge: '/images/elements/badge.png',
				tag: notification.tag
			})
		)
	}
})

self.addEventListener('notificationclick', function(event) {
	console.log('On notification click: ', event.notification)
	// Android doesn't close the notification when you click on it
	// See: http://crbug.com/463146
	if(event.notification.close)
		event.notification.close()

	let url = '/+'

	// This looks to see if the current is already open and
	// focuses if it is
	event.waitUntil(
		clients.matchAll({
			type: 'window'
		}).then(function(clientList) {
			for(let i = 0; i < clientList.length; i++) {
				let client = clientList[i]
				if(client.url === url && 'focus' in client)
					return client.focus()
			}

			if(clients.openWindow) {
				return clients.openWindow(url)
			}
		})
	)
})