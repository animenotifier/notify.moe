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
		const cache = SVGIcon.cache.get(this.name)

		if(cache) {
			const svg = await cache

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
			const url = `//media.notify.moe/images/icons/${this.name}.svg`
			const response = await fetch(url)

			if(!response.ok) {
				console.warn(`Failed loading SVG icon: ${url}`)
				reject(response.statusText)
				return
			}

			const text = await response.text()

			Diff.mutations.queue(() => {
				this.innerHTML = text
				const svg = this.firstChild

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
