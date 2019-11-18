export default function* findAll(className: string): IterableIterator<HTMLElement> {
	const elements = document.getElementsByClassName(className)

	for(const element of elements) {
		yield element as HTMLElement
	}
}
