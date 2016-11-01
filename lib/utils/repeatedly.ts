export function repeatedly(updateTimeInSeconds: number, func: Function) {
	func()
	setInterval(func, updateTimeInSeconds * 1000)
}