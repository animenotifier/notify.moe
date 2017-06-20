export function* findAll(query: string) {
	let elements = document.querySelectorAll(query)

	for(let i = 0; i < elements.length; ++i) {
		yield elements[i] as HTMLElement
	}
}