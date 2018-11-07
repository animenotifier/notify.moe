import AnimeNotifier from "./AnimeNotifier"

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

	notificationCount(e: ServerEvent) {
		this.arn.notificationManager.setCounter(parseInt(e.data))
	}
}