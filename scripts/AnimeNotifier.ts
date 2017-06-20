import { Application } from "./Application"
import { findAll } from "./utils"
import * as actions from "./actions"

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

	loading(isLoading: boolean) {
		if(isLoading) {
			this.app.loading.classList.remove(this.app.fadeOutClass)
		} else {
			this.app.loading.classList.add(this.app.fadeOutClass)
		}
	}

	onContentLoaded() {
		this.updateAvatars()

		for(let element of findAll(".action")) {
			let actionName = element.dataset.action

			element.onclick = () => {
				actions[actionName](this, element)
			}
		}
	}

	updateAvatars() {
		for(let element of findAll(".user-image")) {
			let img = element as HTMLImageElement

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