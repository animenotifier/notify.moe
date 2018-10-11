export function requestIdleCallback(func: Function) {
	if("requestIdleCallback" in window) {
		(window["requestIdleCallback"] as Function)(func)
	} else {
		func()
	}
}