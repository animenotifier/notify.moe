import { AnimeNotifier } from "../AnimeNotifier"
import { Diff } from "../Diff"

// Load more
export async function loadMore(arn: AnimeNotifier, button: HTMLButtonElement) {
	// Prevent firing this event multiple times
	if(arn.isLoading || button.disabled || button.classList.contains("hidden")) {
		return
	}

	arn.loading(true)
	button.disabled = true

	let target = arn.app.find("load-more-target")
	let index = button.dataset.index

	try {
		let response = await fetch("/_" + arn.app.currentPath + "/from/" + index, {
			credentials: "same-origin"
		})

		if(!response.ok) {
			throw response.statusText
		}

		let newIndex = response.headers.get("X-LoadMore-Index")

		// End of data?
		if(newIndex === "-1") {
			button.disabled = true
			button.classList.add("hidden")
		} else {
			button.dataset.index = newIndex
		}

		let body = await response.text()

		let tmp = document.createElement(target.tagName)
		tmp.innerHTML = body

		let children = [...tmp.childNodes]

		Diff.mutations.queue(() => {
			for(let child of children) {
				target.appendChild(child)
			}

			arn.app.emit("DOMContentLoaded")
		})
	} catch(err) {
		arn.statusMessage.showError(err)
	} finally {
		arn.loading(false)
		button.disabled = false
	}
}