import { AnimeNotifier } from "./AnimeNotifier"

export class ServiceWorkerManager {
	arn: AnimeNotifier
	uri: string

	constructor(arn: AnimeNotifier, uri: string) {
		this.arn = arn
		this.uri = uri
	}

	register() {
		if(!("serviceWorker" in navigator)) {
			console.log("service worker not supported, skipping registration")
			return
		}

		console.log("register service worker")

		navigator.serviceWorker.register(this.uri).then(registration => {
			// registration.update()
		})

		navigator.serviceWorker.addEventListener("message", evt => {
			this.onMessage(evt)
		})

		// This will send a message to the service worker that the DOM has been loaded
		let sendContentLoadedEvent = () => {
			if(!navigator.serviceWorker.controller) {
				return
			}

			// A reloadContent call should never trigger another reload
			if(this.arn.app.currentPath === this.arn.lastReloadContentPath) {
				console.log("reload finished.")
				this.arn.lastReloadContentPath = ""
				return
			}

			let message = {
				type: "loaded",
				url: ""
			}

			// If mainPageLoaded is set, it means every single request is now an AJAX request for the /_/ prefixed page
			if(this.arn.mainPageLoaded) {
				message.url = window.location.origin + "/_" + window.location.pathname
			} else {
				this.arn.mainPageLoaded = true
				message.url = window.location.href
			}

			// console.log("checking for updates:", message.url)

			this.postMessage(message)
		}

		// For future loaded events
		document.addEventListener("DOMContentLoaded", sendContentLoadedEvent)

		// If the page is loaded already, send the loaded event right now.
		if(document.readyState !== "loading") {
			sendContentLoadedEvent()
		}
	}

	postMessage(message: any) {
		navigator.serviceWorker.controller.postMessage(JSON.stringify(message))
	}

	onMessage(evt: ServiceWorkerMessageEvent) {
		let message = JSON.parse(evt.data)

		switch(message.type) {
			case "new content":
				if(message.url.includes("/_/")) {
					// Content reload
					this.arn.contentLoadedActions.then(() => {
						this.arn.reloadContent(true)
					})
				} else {
					// Full page reload
					this.arn.contentLoadedActions.then(() => {
						this.arn.reloadPage()
					})
				}

				break

			// case "offline":
			// 	this.arn.statusMessage.showError("You are viewing an offline version of the site now.")
			// 	break
		}
	}
}