class LoadOptions {
	addToHistory?: boolean
	forceReload?: boolean
}

export class Application {
	ajaxClass: string
	fadeOutClass: string
	activeLinkClass: string
	content: HTMLElement
	loading: HTMLElement
	currentPath: string
	originalPath: string
	lastRequest: XMLHttpRequest

	constructor() {
		this.currentPath = window.location.pathname
		this.originalPath = window.location.pathname
		this.ajaxClass = "ajax"
		this.activeLinkClass = "active"
		this.fadeOutClass = "fade-out"
	}

	run() {
		this.ajaxify()
		this.markActiveLinks()
		this.loading.classList.add(this.fadeOutClass)
	}

	find(id: string): HTMLElement {
		return document.getElementById(id)
	}

	get(url: string): Promise<string> {
		if(this.lastRequest) {
			this.lastRequest.abort()
			this.lastRequest = null
		}
		
		return new Promise((resolve, reject) => {
			let request = new XMLHttpRequest()

			request.onerror = () => reject(new Error("You are either offline or the requested page doesn't exist."))
			request.ontimeout = () => reject(new Error("The page took too much time to respond."))
			request.onload = () => {
				if(request.status < 200 || request.status >= 400)
					reject(request.responseText)
				else
					resolve(request.responseText)
			}

			request.open("GET", url, true)
			request.send()

			this.lastRequest = request
		})
	}

	load(url: string, options?: LoadOptions) {
		if(!options) {
			options = new LoadOptions()
		}

		if(options.addToHistory === undefined) {
			options.addToHistory = true
		}
	
		this.currentPath = url

		// Start sending a network request
		let request = this.get("/_" + url).catch(error => error)

		let onTransitionEnd = e => {
			// Ignore transitions of child elements.
			// We only care about the transition event on the content element.
			if(e.target !== this.content) {
				return
			}
			
			// Remove listener after we finally got the correct event.
			this.content.removeEventListener("transitionend", onTransitionEnd)

			// Wait for the network request to end.
			request.then(html => {
				// Add to browser history
				if(options.addToHistory)
					history.pushState(url, null, url)
				
				// Set content
				this.setContent(html)
				this.scrollToTop()

				// Fade animations
				this.content.classList.remove(this.fadeOutClass)
				this.loading.classList.add(this.fadeOutClass)

				// Send DOMContentLoaded Event
				this.emit("DOMContentLoaded")
			})
		}

		this.content.addEventListener("transitionend", onTransitionEnd)

		this.content.classList.add(this.fadeOutClass)
		this.loading.classList.remove(this.fadeOutClass)
		this.markActiveLinks()

		return request
	}

	setContent(html: string) {
		this.content.innerHTML = html
		this.ajaxify(this.content)
		this.markActiveLinks(this.content)
	}

	markActiveLinks(element?: HTMLElement) {
		if(element === undefined)
			element = document.body

		let links = element.querySelectorAll("a")

		for(let i = 0; i < links.length; i++) {
			let link = links[i]
			let href = link.getAttribute("href")

			if(href === this.currentPath)
				link.classList.add(this.activeLinkClass)
			else
				link.classList.remove(this.activeLinkClass)
		}
	}

	ajaxify(element?: HTMLElement) {
		if(!element)
			element = document.body

		let links = element.querySelectorAll("." + this.ajaxClass)

		for(let i = 0; i < links.length; i++) {
			let link = links[i] as HTMLElement

			link.classList.remove(this.ajaxClass)

			let self = this
			link.onclick = function(e) {
				// Middle mouse button should have standard behaviour
				if(e.which === 2)
					return

				let url = this.getAttribute("href")

				e.preventDefault()
				e.stopPropagation()

				if(!url || url === self.currentPath)
					return
				
				// Load requested page
				self.load(url)
			}
		}
	}

	scrollToTop() {
		let parent = this.content

		while(parent = parent.parentElement) {
			parent.scrollTop = 0
		}
	}

	emit(eventName: string) {
		document.dispatchEvent(new Event(eventName, {
			"bubbles": true,
			"cancelable": true
		}))
	}
}