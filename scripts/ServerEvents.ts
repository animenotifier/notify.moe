class ServerEvent {
	data: string
}

export default class ServerEvents {
	supported: boolean
	eventSource: EventSource

	constructor() {
		this.supported = ("EventSource" in window)

		if(!this.supported) {
			return
		}

		this.eventSource = new EventSource("/api/sse/events", {
			withCredentials: true
		})

		this.eventSource.addEventListener("ping", (e: any) => this.ping(e))
	}

	ping(e: ServerEvent) {
		console.log("sse: ping")
	}
}