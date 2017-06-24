export class Diff {
	static childNodes(aRoot: HTMLElement, bRoot: HTMLElement) {
		let aChild = [...aRoot.childNodes]
		let bChild = [...bRoot.childNodes]
		let numNodes = Math.max(aChild.length, bChild.length)
		
		for(let i = 0; i < numNodes; i++) {
			let a = aChild[i] as HTMLElement

			if(i >= bChild.length) {
				aRoot.removeChild(a)
				continue
			}

			let b = bChild[i] as HTMLElement

			if(i >= aChild.length) {
				aRoot.appendChild(b)
				continue
			}

			if(a.nodeName !== b.nodeName || a.nodeType !== b.nodeType) {
				aRoot.replaceChild(b, a)
				continue
			}

			if(a.nodeType === Node.ELEMENT_NODE) {
				if(a.tagName === "IFRAME") {
					continue
				}

				let removeAttributes: Attr[] = []
				
				for(let x = 0; x < a.attributes.length; x++) {
					let attrib = a.attributes[x]

					if(attrib.specified) {
						if(!b.hasAttribute(attrib.name)) {
							removeAttributes.push(attrib)
						}
					}
				}

				for(let attr of removeAttributes) {
					a.removeAttributeNode(attr)
				}

				for(let x = 0; x < b.attributes.length; x++) {
					let attrib = b.attributes[x]

					if(attrib.specified) {
						a.setAttribute(attrib.name, b.getAttribute(attrib.name))
					}
				}

				// Special case: Apply state of input elements
				if(a !== document.activeElement && a instanceof HTMLInputElement && b instanceof HTMLInputElement) {
					a.value = b.value
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