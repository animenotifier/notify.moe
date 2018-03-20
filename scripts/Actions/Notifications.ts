import { AnimeNotifier } from "../AnimeNotifier"

// Enable notifications
export async function enableNotifications(arn: AnimeNotifier, button: HTMLElement) {
	arn.statusMessage.showInfo("Enabling push notifications...")
	await arn.pushManager.subscribe(arn.user.dataset.id)
	arn.updatePushUI()
	arn.statusMessage.showInfo("Enabled push notifications for this device.")
}

// Disable notifications
export async function disableNotifications(arn: AnimeNotifier, button: HTMLElement) {
	arn.statusMessage.showInfo("Disabling push notifications...")
	await arn.pushManager.unsubscribe(arn.user.dataset.id)
	arn.updatePushUI()
	arn.statusMessage.showInfo("Disabled push notifications for this device.")
}

// Test notification
export async function testNotification(arn: AnimeNotifier) {
	await fetch("/api/test/notification", {
		credentials: "same-origin"
	})
}

// Mark notifications as seen
export async function markNotificationsAsSeen(arn: AnimeNotifier) {
	await fetch("/api/mark/notifications/seen", {
		credentials: "same-origin"
	})

	// Update notification counter
	if("serviceWorker" in navigator) {
		// If we have service worker support, broadcast the "notifications marked as seen" message to all open tabs
		arn.serviceWorkerManager.postMessage({
			type: "broadcast",
			realType: "notifications marked as seen"
		})
	} else {
		// If there is no service worker, update the counter on this tab
		arn.notificationManager.update()
	}

	// Update notifications
	arn.reloadContent()
}