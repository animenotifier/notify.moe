import { delay } from "./Utils"

export default class StatusMessage {
	container: HTMLElement
	text: HTMLElement

	constructor(container: HTMLElement, text: HTMLElement) {
		this.container = container
		this.text = text
	}

	show(message: string, duration: number) {
		let messageId = String(Date.now())

		this.text.innerText = message

		this.container.classList.remove("fade-out")
		this.container.dataset.messageId = messageId

		delay(duration || 4000).then(() => {
			if(this.container.dataset.messageId !== messageId) {
				return
			}

			this.close()
		})
	}

	clearStyle() {
		this.container.classList.remove("info-message")
		this.container.classList.remove("error-message")
	}

	showError(message: string, duration?: number) {
		this.clearStyle()
		this.show(message, duration || 4000)
		this.container.classList.add("error-message")
	}

	showInfo(message: string, duration?: number) {
		this.clearStyle()
		this.show(message, duration || 2000)
		this.container.classList.add("info-message")
	}

	close() {
		this.container.classList.add("fade-out")
	}
}