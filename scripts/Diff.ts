export class Diff {
	static childNodes(aRoot: Node, bRoot: Node) {
		let aChild = [...aRoot.childNodes]
		let bChild = [...bRoot.childNodes]
		let numNodes = Math.max(aChild.length, bChild.length)
		
		for(let i = 0; i < numNodes; i++) {
			let a = aChild[i]

			if(i >= bChild.length) {
				aRoot.removeChild(a)
				continue
			}

			let b = bChild[i]

			if(i >= aChild.length) {
				aRoot.appendChild(b)
				continue
			}

			if(a.nodeName !== b.nodeName || a.nodeType !== b.nodeType) {
				aRoot.replaceChild(b, a)
				continue
			}

			if(a.nodeType === Node.TEXT_NODE) {
				a.textContent = b.textContent
				continue
			}

			if(a.nodeType === Node.ELEMENT_NODE) {
				let elemA = a as HTMLElement
				let elemB = b as HTMLElement

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
						if(attrib.name == "class" && elemA.classList.contains("mounted")) {
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

	static innerHTML(aRoot: HTMLElement, html: string) {
		let bRoot = document.createElement("main")
		bRoot.innerHTML = html
		
		Diff.childNodes(aRoot, bRoot)
	}
}