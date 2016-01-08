'use strict'

exports.airingDate = (a, b) => {
	if(a.airingDate.timeStamp === null && b.airingDate.timeStamp === null)
		return sortAlgorithms.alphabetically(a, b)

	if(a.airingDate.timeStamp !== null && b.airingDate.timeStamp === null)
		return -1

	if(a.airingDate.timeStamp === null && b.airingDate.timeStamp !== null)
		return 1

	return a.airingDate.timeStamp - b.airingDate.timeStamp
}

exports.alphabetically = (a, b) => {
	let aLower = a.title.toLowerCase()
	let bLower = b.title.toLowerCase()

	if(aLower < bLower)
		return -1

	if(aLower > bLower)
		return 1

	return 0
}