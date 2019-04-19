import AnimeNotifier from "../AnimeNotifier"
import Diff from "../Diff"

// Load more
export async function loadMore(arn: AnimeNotifier, button: HTMLButtonElement) {
	// Prevent firing this event multiple times
	if(arn.isLoading || button.disabled || button.classList.contains("hidden")) {
		return
	}

	const target = document.getElementById("load-more-target")

	if(!target) {
		return
	}

	arn.loading(true)
	button.disabled = true

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
		if(!newIndex || newIndex === "-1") {
			button.disabled = true
			button.classList.add("hidden")
		} else {
			button.dataset.index = newIndex
		}

		// Get the HTML response
		let html = await response.text()

		// Add the HTML to the existing target
		Diff.mutations.queue(() => {
			target.insertAdjacentHTML("beforeend", html)
			arn.app.emit("DOMContentLoaded")
		})
	} catch(err) {
		arn.statusMessage.showError(err)
	} finally {
		arn.loading(false)
		button.disabled = false
	}
}