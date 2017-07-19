export class Diff {
	static persistentClasses = new Set<string>()

	// Reuse container for diffs to avoid memory allocation
	static container: HTMLElement

	// innerHTML will diff the element with the given HTML string and apply DOM mutations.
	static innerHTML(aRoot: HTMLElement, html: string) {
		if(!Diff.container) {
			Diff.container = document.createElement("main")
		}
		
		Diff.container.innerHTML = html
		Diff.childNodes(aRoot, Diff.container)
	}

	// childNodes diffs the child nodes of 2 given elements and applies DOM mutations.
	static childNodes(aRoot: Node, bRoot: Node) {
		let aChild = [...aRoot.childNodes]
		let bChild = [...bRoot.childNodes]
		let numNodes = Math.max(aChild.length, bChild.length)
		
		for(let i = 0; i < numNodes; i++) {
			let a = aChild[i]

			// Remove nodes at the end of a that do not exist in b
			if(i >= bChild.length) {
				aRoot.removeChild(a)
				continue
			}

			let b = bChild[i]

			// If a doesn't have that many nodes, simply append at the end of a
			if(i >= aChild.length) {
				aRoot.appendChild(b)
				continue
			}

			// If it's a completely different HTML tag or node type, replace it
			if(a.nodeName !== b.nodeName || a.nodeType !== b.nodeType) {
				aRoot.replaceChild(b, a)
				continue
			}

			// Text node:
			// We don't need to check for b to be a text node as well because
			// we eliminated different node types in the previous condition.
			if(a.nodeType === Node.TEXT_NODE) {
				a.textContent = b.textContent
				continue
			}

			// HTML element:
			if(a.nodeType === Node.ELEMENT_NODE) {
				let elemA = a as HTMLElement
				let elemB = b as HTMLElement

				// Skip iframes
				// This part needs to be executed AFTER lazy images check
				// to allow lazily loaded iframes to update their data src.
				if(elemA.tagName === "IFRAME") {
					continue
				}

				let removeAttributes: Attr[] = []
				
				for(let x = 0; x < elemA.attributes.length; x++) {
					let attrib = elemA.attributes[x]

					if(attrib.specified) {
						if(!elemB.hasAttribute(attrib.name)) {
							removeAttributes.push(attrib)
						}
					}
				}

				for(let attr of removeAttributes) {
					elemA.removeAttributeNode(attr)
				}

				for(let x = 0; x < elemB.attributes.length; x++) {
					let attrib = elemB.attributes[x]

					if(attrib.specified) {
						// Skip mountables
						if(attrib.name == "class") {
							let classesA = elemA.classList
							let classesB = elemB.classList

							for(let className of classesA) {
								if(!classesB.contains(className) && !Diff.persistentClasses.has(className)) {
									classesA.remove(className)
								}
							}

							for(let className of classesB) {
								if(!classesA.contains(className)) {
									classesA.add(className)
								}
							}

							continue
						}

						elemA.setAttribute(attrib.name, elemB.getAttribute(attrib.name))
					}
				}

				// Special case: Apply state of input elements
				if(elemA !== document.activeElement && elemA instanceof HTMLInputElement && elemB instanceof HTMLInputElement) {
					elemA.value = elemB.value
				}
			}

			Diff.childNodes(a, b)
		}
	}
}