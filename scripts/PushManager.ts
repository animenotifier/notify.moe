export default class PushManager {
	pushSupported: boolean

	constructor() {
		this.pushSupported = ("serviceWorker" in navigator) && ("PushManager" in window)
	}

	async subscription() {
		if(!this.pushSupported) {
			return Promise.resolve(null)
		}

		const registration = await navigator.serviceWorker.ready
		const subscription = await registration.pushManager.getSubscription()

		return Promise.resolve(subscription)
	}

	async subscribe(userId: string) {
		if(!this.pushSupported) {
			return
		}

		const registration = await navigator.serviceWorker.ready
		let subscription = await registration.pushManager.getSubscription()

		if(!subscription) {
			subscription = await registration.pushManager.subscribe({
				userVisibleOnly: true,
				applicationServerKey: urlBase64ToUint8Array("BAwPKVCWQ2_nc7SIGltYfWZhMpW54BSkbwelpa8eYMbqSitmCAGm3xRBdRiq1Wt-hUsE7x59GCcaJxqQtF2hZPw")
			})

			this.subscribeOnServer(subscription, userId)
		} else {
			console.log("Using existing subscription", subscription)
		}
	}

	async unsubscribe(userId: string) {
		if(!this.pushSupported) {
			return
		}

		const registration = await navigator.serviceWorker.ready
		const subscription = await registration.pushManager.getSubscription()

		if(!subscription) {
			console.error("Subscription does not exist")
			return
		}

		await subscription.unsubscribe()

		this.unsubscribeOnServer(subscription, userId)
	}

	subscribeOnServer(subscription: PushSubscription, userId: string) {
		console.log("Send subscription to server...")

		const rawKey = subscription.getKey("p256dh")
		const key = rawKey ? btoa(String.fromCharCode.apply(null, new Uint8Array(rawKey))) : ""

		const rawSecret = subscription.getKey("auth")
		const secret = rawSecret ? btoa(String.fromCharCode.apply(null, new Uint8Array(rawSecret))) : ""

		const endpoint = subscription.endpoint

		const pushSubscription = {
			endpoint,
			p256dh: key,
			auth: secret,
			platform: navigator.platform,
			userAgent: navigator.userAgent,
			screen: {
				width: window.screen.width,
				height: window.screen.height
			}
		}

		return fetch(`/api/pushsubscriptions/${userId}/add`, {
			method: "POST",
			credentials: "same-origin",
			body: JSON.stringify(pushSubscription)
		})
	}

	unsubscribeOnServer(subscription: PushSubscription, userId: string) {
		console.log("Send unsubscription to server...")
		console.log(subscription)

		const pushSubscription = {
			endpoint: subscription.endpoint
		}

		return fetch("/api/pushsubscriptions/" + userId + "/remove", {
			method: "POST",
			credentials: "same-origin",
			body: JSON.stringify(pushSubscription)
		})
	}
}

function urlBase64ToUint8Array(base64String: string) {
	const padding = "=".repeat((4 - base64String.length % 4) % 4)
	const base64 = (base64String + padding)
	.replace(/\-/g, "+")
	.replace(/_/g, "/")

	const rawData = window.atob(base64)
	return Uint8Array.from([...rawData].map(char => char.charCodeAt(0)))
}
