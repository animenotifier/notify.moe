export class MutationQueue {
	elements: Array<HTMLElement>
	mutation: (elem: HTMLElement) => void

	constructor(mutation: (elem: HTMLElement) => void) {
		this.mutation = mutation
		this.elements = []
	}

	queue(elem: HTMLElement) {
		this.elements.push(elem)

		if(this.elements.length === 1) {
			window.requestAnimationFrame(() => this.mutateAll())
		}
	}

	mutateAll() {
		for(let i = 0; i < this.elements.length; i++) {
			this.mutation(this.elements[i])
		}

		this.clear()
	}

	clear() {
		this.elements.length = 0
	}
}