export function* findAll(className: string): IterableIterator<HTMLElement> {
	const elements = document.getElementsByClassName(className)

	for(let i = 0; i < elements.length; ++i) {
		yield elements[i] as HTMLElement
	}
}

export function* findAllInside(className: string, root: HTMLElement): IterableIterator<HTMLElement> {
	const elements = root.getElementsByClassName(className)

	for(let i = 0; i < elements.length; ++i) {
		yield elements[i] as HTMLElement
	}
}
