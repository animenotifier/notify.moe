class Aero {
	ajaxClass: string
	fadeOutClass: string
	content: HTMLElement
	loading: HTMLElement
	currentURL: string
	originalURL: string
	lastRequest: XMLHttpRequest

	constructor() {
		this.currentURL = window.location.pathname
		this.originalURL = window.location.pathname
		this.ajaxClass = "ajax"
		this.fadeOutClass = "fade-out"
	}

	find(id: string): HTMLElement {
		return document.getElementById(id)
	}

	get(url: string): Promise<string> {
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

	load(url: string, addToHistory: boolean) {
		if(this.lastRequest) {
			this.lastRequest.abort()
			this.lastRequest = null
		}
	
		this.currentURL = url

		console.log(url)

		// function sleep(ms) {
		// 	return new Promise(resolve => setTimeout(resolve, ms))
		// }

		// sleep(500).then(() => {
			
		// })

		let request = this.get("/_" + url)

		this.content.addEventListener("transitionend", e => {
			request.then(html => {
				if(addToHistory)
					history.pushState(url, null, url)
				
				this.setContent(html)
				this.scrollToTop()

				this.content.classList.remove(this.fadeOutClass)
				this.loading.classList.add(this.fadeOutClass)
			})
		}, { once: true })

		this.content.classList.add(this.fadeOutClass)
		this.loading.classList.remove(this.fadeOutClass)
	}

	setContent(html: string) {
		this.content.innerHTML = html
		this.ajaxify(this.content)
	}

	ajaxify(element?: HTMLElement) {
		if(!element)
			element = document.body

		let links = element.querySelectorAll("." + this.ajaxClass)

		for(let i = 0; i < links.length; i++) {
			let link = links[i] as HTMLElement

			link.classList.remove(this.ajaxClass)
			link.onclick = function(e) {
				// Middle mouse button should have standard behaviour
				if(e.which === 2)
					return

				let url = this.getAttribute("href")

				e.preventDefault()
				e.stopPropagation()

				if(!url || url === window.location.pathname)
					return

				// Load requested page
				aero.load(url, true)
			}
		}
	}

	run() {
		this.ajaxify()
		this.loading.classList.add(this.fadeOutClass)
	}

	scrollToTop() {
		let parent = this.content

		while(parent = parent.parentElement) {
			parent.scrollTop = 0
		}
	}
}

export var aero = new Aero()