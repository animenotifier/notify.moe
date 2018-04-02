import Diff from "./Diff"

export default class NotificationManager {
	unseen: number
	icon: HTMLElement
	counter: HTMLElement

	constructor() {
		this.icon = document.getElementById("notification-icon")
		this.counter = document.getElementById("notification-count")
	}

	async update() {
		let response = await fetch("/api/count/notifications/unseen", {
			credentials: "same-origin"
		})

		let body = await response.text()
		this.unseen = parseInt(body)

		if(this.unseen > 99) {
			this.unseen = 99
		}

		this.render()
	}

	render() {
		Diff.mutations.queue(() => {
			this.counter.innerText = this.unseen.toString()

			if(this.unseen === 0) {
				this.counter.classList.add("hidden")
				this.icon.classList.remove("hidden")
			} else {
				this.icon.classList.add("hidden")
				this.counter.classList.remove("hidden")
			}
		})
	}
}