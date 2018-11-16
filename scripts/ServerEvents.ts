import AnimeNotifier from "./AnimeNotifier"
import { plural } from "./Utils"

const reconnectDelay = 3000

class ServerEvent {
	data: string
}

export default class ServerEvents {
	supported: boolean
	eventSource: EventSource
	arn: AnimeNotifier
	etags: Map<string, string>

	constructor(arn: AnimeNotifier) {
		this.supported = ("EventSource" in window)

		if(!this.supported) {
			return
		}

		this.arn = arn
		this.etags = new Map<string, string>()
		this.connect()
	}

	connect() {
		if(this.eventSource) {
			this.eventSource.close()
		}

		this.eventSource = new EventSource("/api/sse/events", {
			withCredentials: true
		})

		this.eventSource.addEventListener("ping", (e: any) => this.ping(e))
		this.eventSource.addEventListener("etag", (e: any) => this.etag(e))
		this.eventSource.addEventListener("activity", (e: any) => this.activity(e))
		this.eventSource.addEventListener("notificationCount", (e: any) => this.notificationCount(e))

		this.eventSource.onerror = () => {
			setTimeout(() => this.connect(), reconnectDelay)
		}
	}

	ping(e: ServerEvent) {
		console.log("ping")
	}

	etag(e: ServerEvent) {
		let data = JSON.parse(e.data)
		let oldETag = this.etags.get(data.url)
		let newETag = data.etag

		if(oldETag && newETag && oldETag != newETag) {
			this.arn.statusMessage.showInfo("A new version of the website is available. Please refresh the page.", -1)
		}

		this.etags.set(data.url, newETag)
	}

	activity(e: ServerEvent) {
		if(!location.pathname.startsWith("/activity")) {
			return
		}

		let isFollowingUser = JSON.parse(e.data)

		// If we're on the followed only feed and we receive an activity
		// about a user we don't follow, ignore the message.
		if(location.pathname.startsWith("/activity/followed") && !isFollowingUser) {
			return
		}

		let button = document.getElementById("load-new-activities")

		if(!button) {
			return
		}

		let buttonText = document.getElementById("load-new-activities-text")
		let newCount = parseInt(button.dataset.count) + 1
		button.dataset.count = newCount.toString()
		buttonText.textContent = plural(newCount, "new activity")
	}

	notificationCount(e: ServerEvent) {
		this.arn.notificationManager.setCounter(parseInt(e.data))
	}
}