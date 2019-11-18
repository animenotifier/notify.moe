export default function* findAllInside(className: string, root: HTMLElement): IterableIterator<HTMLElement> {
	const elements = root.getElementsByClassName(className)

	for(const element of elements) {
		yield element as HTMLElement
	}
}
