import Diff from "scripts/Diff"

export default class ToolTip extends HTMLElement {
	anchor: HTMLElement
	box: HTMLElement
	arrow: HTMLElement

	connectedCallback() {
		this.box = document.createElement("div")
		this.box.classList.add("box")
		this.appendChild(this.box)

		this.arrow = document.createElement("div")
		this.arrow.classList.add("arrow")
		this.appendChild(this.arrow)
	}

	hide() {
		this.setAttribute("active", "false")
	}

	show(anchor: HTMLElement) {
		const distanceToBorder = 5

		this.anchor = anchor
		this.box.textContent = this.anchor.getAttribute("aria-label")

		const anchorRect = this.anchor.getBoundingClientRect()
		const boxRect = this.box.getBoundingClientRect()

		let finalX = anchorRect.left + anchorRect.width / 2 - boxRect.width / 2
		let finalY = anchorRect.top - boxRect.height

		const contentRect = {
			left: distanceToBorder,
			top: distanceToBorder,
			right: document.body.clientWidth - distanceToBorder,
			bottom: document.body.clientHeight - distanceToBorder
		}

		let offsetX = 0
		let offsetY = 0

		if(finalX < contentRect.left) {
			offsetX = contentRect.left - finalX
			finalX += offsetX
		} else if(finalX + boxRect.width > contentRect.right) {
			offsetX = contentRect.right - (finalX + boxRect.width)
			finalX += offsetX
		}

		if(finalY < contentRect.top) {
			offsetY = contentRect.top - finalY
			finalY += offsetY
		}

		const arrowX = boxRect.width / 2 - offsetX

		Diff.mutations.queue(() => {
			this.style.left = finalX + "px"
			this.style.top = finalY + "px"
			this.arrow.style.left = arrowX + "px"
			this.setAttribute("active", "true")
		})
	}
}
