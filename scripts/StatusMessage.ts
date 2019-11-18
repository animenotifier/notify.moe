import delay from "./Utils/delay"

export default class StatusMessage {
	private container: HTMLElement
	private text: HTMLElement

	constructor(container: HTMLElement, text: HTMLElement) {
		this.container = container
		this.text = text
	}

	public showError(message: string | Error, duration?: number) {
		this.clearStyle()
		this.show(message.toString(), duration || 4000)
		this.container.classList.add("error-message")
	}

	public showInfo(message: string, duration?: number) {
		this.clearStyle()
		this.show(message, duration || 2000)
		this.container.classList.add("info-message")
	}

	public close() {
		this.container.classList.add("fade-out")
	}

	private show(message: string, duration: number) {
		const messageId = String(Date.now())

		this.text.textContent = message

		this.container.classList.remove("fade-out")
		this.container.dataset.messageId = messageId

		// Negative duration means we're displaying it forever until the user manually closes it
		if(duration === -1) {
			return
		}

		delay(duration || 4000).then(() => {
			if(this.container.dataset.messageId !== messageId) {
				return
			}

			this.close()
		})
	}

	private clearStyle() {
		this.container.classList.remove("info-message")
		this.container.classList.remove("error-message")
	}
}
