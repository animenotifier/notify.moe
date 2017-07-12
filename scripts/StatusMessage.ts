import { delay } from "./Utils"

export class StatusMessage {
	container: HTMLElement
	text: HTMLElement

	constructor(container: HTMLElement, text: HTMLElement) {
		this.container = container
		this.text = text
	}

	show(message: string, duration?: number) {
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

	showError(message: string, duration?: number) {
		this.show(message, duration)
		this.container.classList.add("error-message")
	}

	close() {
		this.container.classList.add("fade-out")
	}
}