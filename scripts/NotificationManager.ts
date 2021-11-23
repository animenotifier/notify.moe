import Diff from "./Diff"

export default class NotificationManager {
	unseen: number
	icon: HTMLElement
	counter: HTMLElement

	constructor(icon: HTMLElement, counter: HTMLElement) {
		this.icon = icon
		this.counter = counter
	}

	async update() {
		const response = await fetch("/api/count/notifications/unseen", {
			credentials: "same-origin"
		})

		const body = await response.text()
		this.setCounter(parseInt(body, 10))
	}

	setCounter(unseen: number) {
		this.unseen = unseen

		if(isNaN(this.unseen)) {
			this.unseen = 0
		}

		if(this.unseen > 99) {
			this.unseen = 99
		}

		this.render()
	}

	render() {
		Diff.mutations.queue(() => {
			this.counter.textContent = this.unseen.toString()

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
