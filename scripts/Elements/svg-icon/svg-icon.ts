import Diff from "scripts/Diff"

export default class SVGIcon extends HTMLElement {
	static cache = new Map<string, Promise<Node>>()

	static get observedAttributes() {
		return ["name"]
	}

	attributeChangedCallback(attrName) {
		if(attrName === "name") {
			this.render()
		}
	}

	async render() {
		let cache = SVGIcon.cache.get(this.name)

		if(cache) {
			let svg = await cache

			Diff.mutations.queue(() => {
				// Remove all existing child nodes
				while(this.firstChild) {
					this.removeChild(this.firstChild)
				}

				// Append a clone of the cached SVG node
				this.appendChild(svg.cloneNode(true))
			})

			return
		}

		SVGIcon.cache.set(this.name, new Promise(async (resolve, reject) => {
			let url = `//media.notify.moe/images/icons/${this.name}.svg`
			let response = await fetch(url)

			if(!response.ok) {
				console.warn(`Failed loading SVG icon: ${url}`)
				reject(response.statusText)
				return
			}

			let text = await response.text()

			Diff.mutations.queue(() => {
				this.innerHTML = text
				let svg = this.firstChild

				if(!svg) {
					console.warn("Invalid SVG icon:", svg)
					return
				}

				resolve(svg)
			})
		}))
	}

	get name() {
		return this.getAttribute("name") || ""
	}

	set name(value: string) {
		this.setAttribute("name", value)
	}
}