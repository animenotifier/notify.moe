import Diff from "scripts/Diff"

export default class SVGIcon extends HTMLElement {
	static cache = new Map<string, Promise<string>>()

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
			let text = await cache
			Diff.mutations.queue(() => this.innerHTML = text)
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
			Diff.mutations.queue(() => this.innerHTML = text)
			resolve(text)
		}))
	}

	get name() {
		return this.getAttribute("name") || ""
	}

	set name(value: string) {
		this.setAttribute("name", value)
	}
}