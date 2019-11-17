export default class TouchController {
	public leftSwipe: Function
	public rightSwipe: Function
	public upSwipe: Function
	public downSwipe: Function

	private x: number
	private y: number
	private threshold: number

	constructor() {
		document.addEventListener("touchstart", evt => this.handleTouchStart(evt), false)
		document.addEventListener("touchmove", evt => this.handleTouchMove(evt), false)

		this.downSwipe = this.upSwipe = this.rightSwipe = this.leftSwipe = () => null
		this.threshold = 3
		this.x = -1
		this.y = -1
	}

	private handleTouchStart(evt) {
		this.x = evt.touches[0].clientX
		this.y = evt.touches[0].clientY
	}

	private handleTouchMove(evt) {
		if(this.x === -1 || this.y === -1) {
			return
		}

		const xUp = evt.touches[0].clientX
		const yUp = evt.touches[0].clientY

		const xDiff = this.x - xUp
		const yDiff = this.y - yUp

		if(Math.abs(xDiff) > Math.abs(yDiff)) {
			if(xDiff > this.threshold) {
				this.leftSwipe()
			} else if(xDiff < -this.threshold) {
				this.rightSwipe()
			}
		} else {
			if(yDiff > this.threshold) {
				this.upSwipe()
			} else if(yDiff < -this.threshold) {
				this.downSwipe()
			}
		}

		this.x = -1
		this.y = -1
	}
}
