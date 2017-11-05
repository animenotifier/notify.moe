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
export function testNotification(arn: AnimeNotifier) {
	fetch("/api/test/notification", {
		credentials: "same-origin"
	})
}