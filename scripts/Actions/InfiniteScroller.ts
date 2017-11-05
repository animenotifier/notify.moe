import { AnimeNotifier } from "../AnimeNotifier"

// Load more
export function loadMore(arn: AnimeNotifier, button: HTMLButtonElement) {
	// Prevent firing this event multiple times
	if(arn.isLoading || button.disabled) {
		return
	}

	arn.loading(true)
	button.disabled = true

	let target = arn.app.find("load-more-target")
	let index = button.dataset.index
	
	fetch("/_" + arn.app.currentPath + "/from/" + index)
	.then(response => {
		let newIndex = response.headers.get("X-LoadMore-Index")

		// End of data?
		if(newIndex === "-1") {
			button.classList.add("hidden")
		} else {
			button.dataset.index = newIndex
		}
		
		return response
	})
	.then(response => response.text())
	.then(body => {
		let tmp = document.createElement(target.tagName)
		tmp.innerHTML = body

		let children = [...tmp.childNodes]

		window.requestAnimationFrame(() => {
			for(let child of children) {
				target.appendChild(child)
			}

			arn.app.emit("DOMContentLoaded")
		})
	})
	.catch(err => arn.statusMessage.showError(err))
	.then(() => {
		arn.loading(false)
		button.disabled = false
	})
}