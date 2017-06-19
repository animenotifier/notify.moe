import { Application } from "./Application"

export class AnimeNotifier {
	app: Application

	constructor(app: Application) {
		this.app = app
	}

	onReadyStateChange() {
		if(document.readyState !== "interactive") {
			return
		}

		this.run()
	}

	run() {
		this.app.content = this.app.find("content")
		this.app.loading = this.app.find("loading")
		this.app.run()
	}

	onContentLoaded() {
		this.updateAvatars()
	}

	updateAvatars() {
		let images = document.querySelectorAll('.user-image')

		for(let i = 0; i < images.length; ++i) {
			let img = images[i] as HTMLImageElement

			if(img.naturalWidth === 0) {
				img.onload = function() {
					this.classList.add("user-image-found")
				}

				img.onerror = function() {
					this.classList.add("user-image-not-found")
				}
			} else {
				img.classList.add("user-image-found")
			}
		}
	}
}