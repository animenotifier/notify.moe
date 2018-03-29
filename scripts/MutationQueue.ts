// Computation time allowed per frame, in milliseconds.
// On a 100 Hz monitor this would ideally be 10 ms.
// On a 200 Hz monitor it should be 5 ms.
// However, the renderer also needs a bit of time,
// so setting the value a little lower guarantees smooth transitions.
const timeCapacity = 6.5

// MutationQueue queues up DOM mutations to batch execute them before a frame is rendered.
// It checks the time used to process these mutations and if the time is over the
// defined time capacity, it will pause and continue the mutations in the next frame.
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