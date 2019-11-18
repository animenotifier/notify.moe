import Diff from "./Diff"
import LoadOptions from "./LoadOptions"
import delay from "./Utils/delay"

export default class Application {
	public originalPath: string
	public currentPath: string
	public content: HTMLElement
	public loading: HTMLElement
	public fadeOutClass: string
	public onError: (err: Error) => void

	private activeLinkClass: string
	private lastRequestController: AbortController | null
	private contentInvisible: boolean

	constructor() {
		this.currentPath = window.location.pathname
		this.originalPath = window.location.pathname
		this.activeLinkClass = "active"
		this.fadeOutClass = "fade-out"
		this.onError = console.error
	}

	public init() {
		document.addEventListener("DOMContentLoaded", this.onContentLoaded.bind(this))
	}

	public async get(url: string): Promise<string> {
		if(this.lastRequestController) {
			this.lastRequestController.abort()
		}

		this.lastRequestController = new AbortController()

		try {
			const response = await fetch(url, {
				credentials: "same-origin",
				signal: this.lastRequestController.signal
			})

			if(!response.ok) {
				throw response.statusText
			}

			return await response.text()
		} catch(err) {
			throw "Seems like there was an error accessing this page...retrying after 3 seconds."
		}
	}

	public load(url: string, options?: LoadOptions) {
		// Remove protocol and hostname if it was specified
		if(url.startsWith(location.origin)) {
			url = url.substr(location.origin.length)
		}

		// Start sending a network request
		let request: Promise<string>

		const retry = () => {
			return this.get("/_" + url).catch(async error => {
				// Are we still on that page?
				if(this.currentPath !== url) {
					return
				}

				// Display connection error
				this.onError(error)

				// Retry after 3 seconds
				await delay(3000)

				// Are we still on that page?
				if(this.currentPath !== url) {
					return
				}

				return retry()
			})
		}

		request = retry()

		// Parse options
		if(!options) {
			options = new LoadOptions()
		}

		if(options.addToHistory === undefined) {
			options.addToHistory = true
		}

		// Set current path
		this.currentPath = url

		// Add to browser history
		if(options.addToHistory) {
			history.pushState(url, "", url)
		}

		// Mark active links
		this.markActiveLinks()

		const consume = async () => {
			const html = await request

			if(this.currentPath !== url) {
				return
			}

			// Set content
			this.setContent(html)
			this.scrollToTop()

			// Fade in listener
			const onFadedIn: EventListener = (e: Event) => {
				// Ignore transitions of child elements.
				// We only care about the transition event on the content element.
				if(e.target !== this.content) {
					return
				}

				// Reset the transition ended flag
				this.contentInvisible = false

				// Remove listener after we finally got the correct event.
				this.content.removeEventListener("transitionend", onFadedIn)
			}

			this.content.addEventListener("transitionend", onFadedIn)

			// Fade animations
			this.content.classList.remove(this.fadeOutClass)
			this.loading.classList.add(this.fadeOutClass)

			// Send DOMContentLoaded Event
			this.emit("DOMContentLoaded")
		}

		if(this.contentInvisible) {
			consume()
		} else {
			// Fade out listener
			const onFadedOut: EventListener = (e: Event) => {
				// Ignore transitions of child elements.
				// We only care about the transition event on the content element.
				if(e.target !== this.content) {
					return
				}

				this.contentInvisible = true

				// Remove listener after we finally got the correct event.
				this.content.removeEventListener("transitionend", onFadedOut)

				// Outdated response.
				if(this.currentPath !== url) {
					return
				}

				// Wait for the network request to end.
				consume()
			}

			this.content.addEventListener("transitionend", onFadedOut)

			// Add fade out class
			this.content.classList.add(this.fadeOutClass)
			this.loading.classList.remove(this.fadeOutClass)
		}

		return request
	}

	public ajaxify(links?: HTMLCollectionOf<HTMLAnchorElement>) {
		if(!links) {
			links = document.getElementsByTagName("a")
		}

		for(const link of links) {
			// Don't ajaxify links to a different host
			if(link.hostname !== window.location.hostname) {
				if(!link.target) {
					link.target = "_blank"
				}

				continue
			}

			// Don't ajaxify invalid links, links with a target or links that disable ajax specifically
			if(link.href === "" || link.href.includes("#") || link.target.length > 0 || link.dataset.ajax === "false") {
				continue
			}

			const self = this

			link.onclick = function(e) {
				// Middle mouse button and Ctrl clicks should have standard behaviour
				if(e.which === 2 || e.ctrlKey) {
					return
				}

				e.preventDefault()

				// Prevent loading the same page
				const url = (this as HTMLAnchorElement).getAttribute("href")

				if(!url || url === self.currentPath) {
					return
				}

				// Load requested page
				self.load(url)
				self.emit("linkclicked")
			}
		}
	}

	public markActiveLinks(links?: HTMLCollectionOf<HTMLAnchorElement>) {
		if(!links) {
			links = document.getElementsByTagName("a")
		}

		for(const link of links) {
			Diff.mutations.queue(() => {
				if(link.getAttribute("href") === this.currentPath) {
					link.classList.add(this.activeLinkClass)
				} else {
					link.classList.remove(this.activeLinkClass)
				}
			})
		}
	}

	public emit(eventName: string) {
		document.dispatchEvent(new Event(eventName))
	}

	private onContentLoaded() {
		const links = document.getElementsByTagName("a")

		this.markActiveLinks(links)
		this.ajaxify(links)
	}

	private setContent(html: string) {
		this.content.innerHTML = html
	}

	private scrollToTop() {
		let parent: HTMLElement | null = this.content

		Diff.mutations.queue(() => {
			while(parent) {
				parent.scrollTop = 0
				parent = parent.parentElement
			}
		})
	}
}
