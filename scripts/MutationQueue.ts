const timeCapacity = 6.5

export class MutationQueue {
	mutations: Array<() => void>
	onClearCallBack: () => void

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
			if(performance.now() - start > timeCapacity) {
				let end = performance.now()
				this.mutations = this.mutations.slice(i)
				window.requestAnimationFrame(() => this.mutateAll())
				return
			}

			try {
				this.mutations[i]()
			} catch(err) {
				console.error(err)
			}
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