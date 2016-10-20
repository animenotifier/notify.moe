let titleToId = {}

export function getIdByTitle(title: string): number {
	return titleToId[title]
}

export { titleToId }