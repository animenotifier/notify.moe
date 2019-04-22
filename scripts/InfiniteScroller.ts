import Diff from "./Diff"

export default class InfiniteScroller {
	container: HTMLElement
	threshold: number

	constructor(container, threshold) {
		this.container = container
		this.threshold = threshold

		let check = () => {
			if(this.container.scrollTop + this.container.clientHeight >= this.container.scrollHeight - threshold) {
				this.loadMore()
			}
		}

		this.container.addEventListener("scroll", _ => {
			// Wait for mutations to finish before checking if we need infinite scroll to trigger.
			if(Diff.mutations.mutations.length > 0) {
				Diff.mutations.wait(() => check())
				return
			}

			// Otherwise, queue up the check immediately.
			// Don't call check() directly to make scrolling as smooth as possible.
			Diff.mutations.queue(check)
		})
	}

	loadMore() {
		let button = document.getElementById("load-more-button")

		if(!button) {
			return
		}

		button.click()
	}
}