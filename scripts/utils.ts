export function* findAll(className: string) {
	// getElementsByClassName failed for some reason.
	// TODO: Test getElementsByClassName again.
	// let elements = document.querySelectorAll("." + className)
	let elements = document.getElementsByClassName(className)
	
	for(let i = 0; i < elements.length; ++i) {
		yield elements[i] as HTMLElement
	}
}

export function delay<T>(millis: number, value?: T): Promise<T> {
	return new Promise(resolve => setTimeout(() => resolve(value), millis))
}