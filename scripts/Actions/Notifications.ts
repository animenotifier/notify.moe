import { AnimeNotifier } from "../AnimeNotifier"

// Enable notifications
export async function enableNotifications(arn: AnimeNotifier, button: HTMLElement) {
	await arn.pushManager.subscribe(arn.user.dataset.id)
	arn.updatePushUI()
}

// Disable notifications
export async function disableNotifications(arn: AnimeNotifier, button: HTMLElement) {
	await arn.pushManager.unsubscribe(arn.user.dataset.id)
	arn.updatePushUI()
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
	arn.notificationManager.update()

	// Update notifications
	arn.reloadContent()
}