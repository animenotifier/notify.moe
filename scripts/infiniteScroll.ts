import Diff from "./Diff"

let container: HTMLElement
let threshold: number

export default function infiniteScroll(scrollContainer: HTMLElement, scrollThreshold: number) {
	container = scrollContainer
	threshold = scrollThreshold

	container.addEventListener("scroll", _ => {
		// Wait for mutations to finish before checking if we need infinite scroll to trigger.
		if(Diff.mutations.length() > 0) {
			Diff.mutations.wait(check)
			return
		}

		// Otherwise, queue up the check immediately.
		// Don't call check() directly to make scrolling as smooth as possible.
		Diff.mutations.queue(() => check())
	})
}

function check() {
	if(container.scrollTop + container.clientHeight >= container.scrollHeight - threshold) {
		loadMore()
	}
}

function loadMore() {
	const button = document.getElementById("load-more-button")

	if(!button) {
		return
	}

	button.click()
}
