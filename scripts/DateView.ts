import { plural } from "./Utils"

const oneSecond = 1000
const oneMinute = 60 * oneSecond
const oneHour = 60 * oneMinute
const oneDay = 24 * oneHour
const oneWeek = 7 * oneDay
const oneMonth = 30 * oneDay
const oneYear = 365.25 * oneDay

export var monthNames = [
	"January", "February", "March",
	"April", "May", "June", "July",
	"August", "September", "October",
	"November", "December"
]

export var dayNames = [
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday"
]

function getRemainingTime(remaining: number): string {
	let remainingAbs = Math.abs(remaining)

	if(remainingAbs >= oneYear) {
		return plural(Math.round(remaining / oneYear), "year")
	}

	if(remainingAbs >= oneMonth) {
		return plural(Math.round(remaining / oneMonth), "month")
	}

	if(remainingAbs >= oneWeek) {
		return plural(Math.round(remaining / oneWeek), "week")
	}

	if(remainingAbs >= oneDay) {
		return plural(Math.round(remaining / oneDay), "day")
	}

	if(remainingAbs >= oneHour) {
		return plural(Math.round(remaining / oneHour), " hour")
	}

	if(remainingAbs >= oneMinute) {
		return plural(Math.round(remaining / oneMinute), " minute")
	}

	if(remainingAbs >= oneSecond) {
		return plural(Math.round(remaining / oneSecond), " second")
	}
	
	return "Just now"
}

export function displayLocalDate(element: HTMLElement, now: Date) {
	let startDate = new Date(element.dataset.startDate)
	let endDate = new Date(element.dataset.endDate)

	let h = startDate.getHours()
	let m = startDate.getMinutes()
	let startTime = (h <= 9 ? "0" + h : h) + ":" + (m <= 9 ? "0" + m : m)

	h = endDate.getHours()
	m = endDate.getMinutes()
	let endTime = (h <= 9 ? "0" + h : h) + ":" + (m <= 9 ? "0" + m : m)
	
	let airingVerb = "will be airing"


	// let dayInfo = dayNames[startDate.getDay()] + ", " + monthNames[startDate.getMonth()] + " " + startDate.getDate()

	let remaining = startDate.getTime() - now.getTime()
	let remainingString = getRemainingTime(remaining)

	// Add "ago" if the date is in the past
	if(remainingString.startsWith("-")) {
		remainingString = remainingString.substring(1) + " ago"
	}

	element.innerText = remainingString

	// let remainingString = seconds + plural(seconds, 'second')

	// let days = seconds / (60 * )
	// if(Math.abs(days) >= 1) {
	// 	remainingString = plural(days, 'day')
	// } else {
	// 	let hours = arn.inHours(now, timeStamp)
	// 	if(Math.abs(hours) >= 1) {
	// 		remainingString = plural(hours, 'hour')
	// 	} else {
	// 		let minutes = arn.inMinutes(now, timeStamp)
	// 		if(Math.abs(minutes) >= 1) {
	// 			remainingString = plural(minutes, 'minute')
	// 		} else {
	// 			let seconds = arn.inSeconds(now, timeStamp)
	// 			remainingString = plural(seconds, 'second')
	// 		}
	// 	}
	// }

	// if(isNaN(oneHour)) {
	// 	element.style.opacity = "0"
	// 	return
	// }

	// switch(Math.floor(dayDifference)) {
	// 	case 0:
	// 		element.innerText = "Today"
	// 		break
	// 	case 1:
	// 		element.innerText = "Tomorrow"
	// 		break
	// 	case -1:
	// 		element.innerText = "Yesterday"
	// 		break
	// 	default:
	// 		let text = Math.abs(dayDifference) + " days"

	// 		if(dayDifference < 0) {
	// 			text += " ago"
	// 			airingVerb = "aired"
	// 		} else {
	// 			element.innerText = text
	// 		}
	// }

	if(remaining < 0) {
		airingVerb = "aired"
	}

	element.title = "Episode " + element.dataset.episodeNumber + " " + airingVerb + " " + dayNames[startDate.getDay()] + " from " + startTime + " - " + endTime
}