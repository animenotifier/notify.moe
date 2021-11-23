import MutationQueue from "./MutationQueue"

// Diff provides diffing utilities to morph existing DOM elements
// into the target HTML string.
//
// Example:
// Diff.innerHTML(body, "<div>This is my new content</div>")
//
// Whatever contents will be in the body, they will be re-used and morphed
// into the new DOM defined by a simple HTML string. This is useful for
// Single Page Applications that use server rendered content. The server
// responds with the pre-rendered HTML and we can simply morph our current
// contents into the next page.
export default class Diff {
	static persistentClasses = new Set<string>()
	static persistentAttributes = new Set<string>()
	static mutations: MutationQueue = new MutationQueue()

	// innerHTML will diff the element with the given HTML string and apply DOM mutations.
	static innerHTML(aRoot: HTMLElement, html: string) {
		const container = document.createElement("main")
		container.innerHTML = html

		return new Promise((resolve, _) => {
			Diff.childNodes(aRoot, container)

			this.mutations.wait(() => {
				resolve(null)
			})
		})
	}

	// root will diff the document root element with the given HTML string and apply DOM mutations.
	static root(aRoot: HTMLElement, html: string) {
		return new Promise((resolve, _) => {
			const rootContainer = document.createElement("html")
			rootContainer.innerHTML = html.replace("<!DOCTYPE html>", "")

			Diff.childNodes(aRoot.getElementsByTagName("body")[0], rootContainer.getElementsByTagName("body")[0])

			this.mutations.wait(() => {
				resolve(null)
			})
		})
	}

	// childNodes diffs the child nodes of 2 given elements and applies DOM mutations.
	static childNodes(aRoot: Node, bRoot: Node) {
		const aChild = [...aRoot.childNodes]
		const bChild = [...bRoot.childNodes]
		const numNodes = Math.max(aChild.length, bChild.length)

		for(let i = 0; i < numNodes; i++) {
			const a = aChild[i]

			// Remove nodes at the end of a that do not exist in b
			if(i >= bChild.length) {
				this.mutations.queue(() => aRoot.removeChild(a))
				continue
			}

			const b = bChild[i]

			// If a doesn't have that many nodes, simply append at the end of a
			if(i >= aChild.length) {
				this.mutations.queue(() => aRoot.appendChild(b))
				continue
			}

			// If it's a completely different HTML tag or node type, replace it
			if(a.nodeName !== b.nodeName || a.nodeType !== b.nodeType) {
				this.mutations.queue(() => aRoot.replaceChild(b, a))
				continue
			}

			// Text node:
			// We don't need to check for b to be a text node as well because
			// we eliminated different node types in the previous condition.
			if(a.nodeType === Node.TEXT_NODE) {
				this.mutations.queue(() => a.textContent = b.textContent)
				continue
			}

			// HTML element:
			if(a.nodeType === Node.ELEMENT_NODE) {
				const elemA = a as HTMLElement
				const elemB = b as HTMLElement

				const removeAttributes: Attr[] = []

				for(const attrib of elemA.attributes) {
					if(attrib.specified) {
						if(!elemB.hasAttribute(attrib.name) && !Diff.persistentAttributes.has(attrib.name)) {
							removeAttributes.push(attrib)
						}
					}
				}

				this.mutations.queue(() => {
					for(const attr of removeAttributes) {
						elemA.removeAttributeNode(attr)
					}
				})

				for(const attrib of elemB.attributes) {
					if(!attrib.specified) {
						continue
					}

					// If the attribute value is exactly the same, skip this attribute.
					if(elemA.getAttribute(attrib.name) === attrib.value) {
						continue
					}

					if(attrib.name === "class") {
						const classesA = elemA.classList
						const classesB = elemB.classList
						const removeClasses: string[] = []

						for(const className of classesA) {
							if(!classesB.contains(className) && !Diff.persistentClasses.has(className)) {
								removeClasses.push(className)
							}
						}

						this.mutations.queue(() => {
							for(const className of removeClasses) {
								classesA.remove(className)
							}

							for(const className of classesB) {
								if(!classesA.contains(className)) {
									classesA.add(className)
								}
							}
						})

						continue
					}

					this.mutations.queue(() => elemA.setAttribute(attrib.name, attrib.value))
				}

				// Special case: Apply state of input elements
				if(elemA !== document.activeElement && elemA instanceof HTMLInputElement && elemB instanceof HTMLInputElement) {
					this.mutations.queue(() => {
						(elemA as HTMLInputElement).value = (elemB as HTMLInputElement).value
					})
				}
			}

			// Never diff the contents of web components
			if(a.nodeName.includes("-")) {
				continue
			}

			// Child nodes
			Diff.childNodes(a, b)
		}
	}
}
