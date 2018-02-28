export class NotificationManager {
	unseen: number

	async update() {
		let response = await fetch("/api/count/notifications/unseen", {
			credentials: "same-origin"
		})

		let body = await response.text()
		this.unseen = parseInt(body)
		this.unseen = 2
		this.render()
	}

	render() {
		console.log("notification count", this.unseen)

		let notificationIcon = document.getElementById("notification-icon")
		let notificationCount = document.getElementById("notification-count")

		notificationCount.innerText = this.unseen.toString()

		if(this.unseen === 0) {
			notificationCount.classList.add("hidden")
			notificationIcon.classList.remove("hidden")
		} else {
			notificationIcon.classList.add("hidden")
			notificationCount.classList.remove("hidden")
		}
	}
}