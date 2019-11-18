import AnimeNotifier from "./AnimeNotifier"

export default class ServiceWorkerManager {
	private arn: AnimeNotifier
	private uri: string

	constructor(arn: AnimeNotifier, uri: string) {
		this.arn = arn
		this.uri = uri
	}

	public register() {
		if(!("serviceWorker" in navigator)) {
			console.warn("service worker not supported, skipping registration")
			return
		}

		navigator.serviceWorker.register(this.uri)
		navigator.serviceWorker.addEventListener("message", evt => this.onMessage(evt))
	}

	public postMessage(message: any) {
		const controller = navigator.serviceWorker.controller

		if(!controller) {
			return
		}

		controller.postMessage(JSON.stringify(message))
	}

	private onMessage(evt: MessageEvent) {
		const message = JSON.parse(evt.data)

		switch(message.type) {
			case "new notification":
				if(this.arn.notificationManager) {
					this.arn.notificationManager.update()
				}

				break
		}
	}
}
