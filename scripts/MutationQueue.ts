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

export class CustomMutationQueue {
	mutations: Array<() => void>
	onClearCallBack: () => void
	timeCapacity = 6.5

	constructor() {
		this.mutations = []
	}

	queue(mutation: () => void) {
		this.mutations.push(mutation)

		if(this.mutations.length === 1) {
			window.requestAnimationFrame(() => this.mutateAll())
		}
	}

	mutateAll() {
		let start = performance.now()

		for(let i = 0; i < this.mutations.length; i++) {
			if(performance.now() - start > this.timeCapacity) {
				let end = performance.now()
				// console.log(i, "mutations in", performance.now() - start, "ms")
				this.mutations = this.mutations.slice(i)
				window.requestAnimationFrame(() => this.mutateAll())
				return
			}

			this.mutations[i]()
		}

		this.clear()
	}

	clear() {
		this.mutations.length = 0

		if(this.onClearCallBack) {
			this.onClearCallBack()
			this.onClearCallBack = null
		}
	}

	wait(callBack: () => void) {
		this.onClearCallBack = callBack
	}
}