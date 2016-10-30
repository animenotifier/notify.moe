// Alphabetically
function alphabetically(a, b): number {
	const aLower: string = a.preferredTitle.toLowerCase()
	const bLower: string = b.preferredTitle.toLowerCase()

	if(aLower < bLower)
		return -1

	if(aLower > bLower)
		return 1

	return 0
}

// By airing date
function airingDate(a, b): number {
	if(a.airingDate.timeStamp === null && b.airingDate.timeStamp === null)
		return alphabetically(a, b)

	if(a.airingDate.timeStamp !== null && b.airingDate.timeStamp === null)
		return -1

	if(a.airingDate.timeStamp === null && b.airingDate.timeStamp !== null)
		return 1

	return a.airingDate.timeStamp - b.airingDate.timeStamp
}

// Sort algorithms
export const sortAlgorithms = {
	airingDate,
	alphabetically
}