var isPushEnabled = false;

function subscribeOnServer(subscription) {
	console.log('Send subscription to server...');
	console.log(subscription);

	kaze.postJSON('/api/notifications/subscribe', {
		endpoint: subscription.endpoint
	}).then(function(response) {
		console.log(response);
	});
}

function unsubscribeOnServer(subscription) {
	console.log('Send unsubscription to server...');
	console.log(subscription);

	kaze.postJSON('/api/notifications/unsubscribe', {
		endpoint: subscription.endpoint
	}).then(function(response) {
		console.log(response);
	});
}

function sendTestNotification() {
	console.log('Sending test notification...')

	kaze.get('/api/notifications/test').then(function(response) {
		// ...
	})
}

function subscribe() {
	// Disable the button so it can't be changed while
	// we process the permission request
	var pushButton = document.querySelector('.push-button');
	if(pushButton)
		pushButton.disabled = true;

	navigator.serviceWorker.ready.then(function(serviceWorkerRegistration) {
		serviceWorkerRegistration.pushManager.subscribe({
			userVisibleOnly: true
		}).then(function(subscription) {
			// The subscription was successful
			isPushEnabled = true;
			if(pushButton) {
				pushButton.textContent = 'Disable Notifications';
				pushButton.disabled = false;
			}
			document.querySelector('.test-notification').disabled = false;

			// TODO: Send the subscription.endpoint to your server
			// and save it to send a push message at a later date
			return subscribeOnServer(subscription);
		}).catch(function(e) {
			if(Notification.permission === 'denied') {
				// The user denied the notification permission which
				// means we failed to subscribe and the user will need
				// to manually change the notification permission to
				// subscribe to push messages
				console.warn('Permission for Notifications was denied');
				if(pushButton)
					pushButton.disabled = true;
			} else {
				// A problem occurred with the subscription; common reasons
				// include network errors, and lacking gcm_sender_id and/or
				// gcm_user_visible_only in the manifest.
				console.error('Unable to subscribe to push.', e);
				if(pushButton) {
					pushButton.disabled = false;
					pushButton.textContent = 'Enable Notifications';
				}
			}
		});
	});
}

function unsubscribe() {
	var pushButton = document.querySelector('.push-button');
	pushButton.disabled = true;
	document.querySelector('.test-notification').disabled = true;

	navigator.serviceWorker.ready.then(function(serviceWorkerRegistration) {
		// To unsubscribe from push messaging, you need get the
		// subscription object, which you can call unsubscribe() on.
		serviceWorkerRegistration.pushManager.getSubscription().then(function(pushSubscription) {
			// Check we have a subscription to unsubscribe
			if(!pushSubscription) {
				// No subscription object, so set the state
				// to allow the user to subscribe to push
				isPushEnabled = false;
				if(pushButton) {
					pushButton.disabled = false;
					pushButton.textContent = 'Enable Notifications';
				}
				return;
			}

			unsubscribeOnServer(pushSubscription);

			// We have a subscription, so call unsubscribe on it
			pushSubscription.unsubscribe().then(function(successful) {
				if(pushButton) {
					pushButton.disabled = false;
					pushButton.textContent = 'Enable Notifications';
				}
				isPushEnabled = false;
			}).catch(function(e) {
				// We failed to unsubscribe, this can lead to
				// an unusual state, so may be best to remove
				// the users data from your data store and
				// inform the user that you have done so

				console.log('Unsubscription error: ', e);
				if(pushButton) {
					pushButton.disabled = false;
					pushButton.textContent = 'Enable Notifications';
				}
			});
		}).catch(function(e) {
			console.error('Error thrown while unsubscribing from push messaging.', e);
		});
	});
}

// Once the service worker is registered set the initial state
function initialiseState(registration) {
	console.log('Initialise state...');
	console.log('Scope:', registration.scope);

	// Are Notifications supported in the service worker?
	if(!('showNotification' in ServiceWorkerRegistration.prototype)) {
		console.warn('Notifications aren\'t supported.');
		return;
	}

	// Check the current Notification permission.
	// If its denied, it's a permanent block until the
	// user changes the permission
	if(Notification.permission === 'denied') {
		console.warn('The user has blocked notifications.');
		return;
	}

	// Check if push messaging is supported
	if(!('PushManager' in window)) {
		console.warn('Push messaging isn\'t supported.');
		return;
	}

	// We need the service worker registration to check for a subscription
	navigator.serviceWorker.ready.then(function(serviceWorkerRegistration) {
		console.log('Get subscription...');

		// Do we already have a push message subscription?
		serviceWorkerRegistration.pushManager.getSubscription().then(function(subscription) {
			console.log('Enable Push UI...');

			// Enable any UI which subscribes / unsubscribes from
			// push messages.
			var pushButton = document.querySelector('.push-button');
			if(pushButton)
				pushButton.disabled = false;

			if(!subscription) {
				// We aren't subscribed to push, so set UI
				// to allow the user to enable push
				return;
			}

			// Keep your server in sync with the latest subscriptionId
			subscribeOnServer(subscription);

			// Set your UI to show they have subscribed for
			// push messages
			if(pushButton)
				pushButton.textContent = 'Disable Notifications';
			isPushEnabled = true;

			// Enable test notifications
			document.querySelector('.test-notification').disabled = false;
		}).catch(function(err) {
			console.warn('Error during getSubscription()', err);
		});
	});
}