export function* findAll(className: string): IterableIterator<HTMLElement> {
	let elements = document.getElementsByClassName(className)

	for(let i = 0; i < elements.length; ++i) {
		yield elements[i] as HTMLElement
	}
}