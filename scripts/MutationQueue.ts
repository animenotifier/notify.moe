// Computation time allowed per frame, in milliseconds.
// On a 100 Hz monitor this would ideally be 10 ms.
// On a 200 Hz monitor it should be 5 ms.
// However, the renderer also needs a bit of time,
// so setting the value a little lower guarantees smooth transitions.
const timeCapacity = 6.5

// MutationQueue queues up DOM mutations to batch execute them before a frame is rendered.
// It checks the time used to process these mutations and if the time is over the
// defined time capacity, it will pause and continue the mutations in the next frame.
export default class MutationQueue {
	private mutations: Array<() => void>
	private onClearCallBacks: Array<() => void>

	constructor() {
		this.mutations = []
		this.onClearCallBacks = []
	}

	public queue(mutation: () => void) {
		this.mutations.push(mutation)

		if(this.mutations.length === 1) {
			window.requestAnimationFrame(() => this.mutateAll())
		}
	}

	public wait(callBack: () => void) {
		if(this.mutations.length === 0) {
			callBack()
			return
		}

		this.onClearCallBacks.push(callBack)
	}

	public length() {
		return this.mutations.length
	}

	private mutateAll() {
		const start = Date.now()

		for(let i = 0; i < this.mutations.length; i++) {
			if(Date.now() - start > timeCapacity) {
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

	private clear() {
		this.mutations.length = 0

		if(this.onClearCallBacks.length > 0) {
			for(const callback of this.onClearCallBacks) {
				callback()
			}

			this.onClearCallBacks.length = 0
		}
	}
}
