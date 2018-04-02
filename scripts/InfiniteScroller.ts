export default class InfiniteScroller {
	container: HTMLElement
	threshold: number

	constructor(container, threshold) {
		this.container = container
		this.threshold = threshold

		this.container.addEventListener("scroll", e => {
			if(this.container.scrollTop + this.container.clientHeight >= this.container.scrollHeight - threshold) {
				this.loadMore()
			}
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