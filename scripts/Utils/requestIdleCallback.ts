export function requestIdleCallback(func: Function) {
	if("requestIdleCallback" in window) {
		let requestIdleCallback = window["requestIdleCallback"] as Function
		requestIdleCallback(func)
	} else {
		func()
	}
}