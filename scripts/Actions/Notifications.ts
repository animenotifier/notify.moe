import AnimeNotifier from "../AnimeNotifier"

// Enable notifications
export async function enableNotifications(arn: AnimeNotifier, button: HTMLElement) {
	arn.statusMessage.showInfo("Enabling instant notifications...")
	await arn.pushManager.subscribe(arn.user.dataset.id)
	arn.updatePushUI()
	arn.statusMessage.showInfo("Enabled instant notifications for this device.")
}

// Disable notifications
export async function disableNotifications(arn: AnimeNotifier, button: HTMLElement) {
	arn.statusMessage.showInfo("Disabling instant notifications...")
	await arn.pushManager.unsubscribe(arn.user.dataset.id)
	arn.updatePushUI()
	arn.statusMessage.showInfo("Disabled instant notifications for this device.")
}

// Test notification
export async function testNotification(arn: AnimeNotifier) {
	arn.statusMessage.showInfo("Sending test notification...this might take a few seconds...")

	await fetch("/api/test/notification", {
		credentials: "same-origin"
	})
}

// Mark notifications as seen
export async function markNotificationsAsSeen(arn: AnimeNotifier) {
	await fetch("/api/mark/notifications/seen", {
		credentials: "same-origin"
	})

	// Update notifications
	arn.reloadContent()
}