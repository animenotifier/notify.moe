const oneDay = 24 * 60 * 60 * 1000

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

export function displayLocalDate(element: HTMLElement, now: Date) {
	let startDate = new Date(element.dataset.startDate)
	let endDate = new Date(element.dataset.endDate)

	let h = startDate.getHours()
	let m = startDate.getMinutes()
	let startTime = (h <= 9 ? "0" + h : h) + ":" + (m <= 9 ? "0" + m : m)

	h = endDate.getHours()
	m = endDate.getMinutes()
	let endTime = (h <= 9 ? "0" + h : h) + ":" + (m <= 9 ? "0" + m : m)
	
	let dayDifference = Math.round((startDate.getTime() - now.getTime()) / oneDay)

	if(isNaN(dayDifference)) {
		element.style.opacity = "0"
		return
	}

	let dayInfo = dayNames[startDate.getDay()] + ", " + monthNames[startDate.getMonth()] + " " + startDate.getDate()

	let airingVerb = "will be airing"

	switch(dayDifference) {
		case 0:
			element.innerText = "Today"
			break
		case 1:
			element.innerText = "Tomorrow"
			break
		case -1:
			element.innerText = "Yesterday"
			break
		default:
			let text = Math.abs(dayDifference) + " days"

			if(dayDifference < 0) {
				text += " ago"
				airingVerb = "aired"
			} else {
				element.innerText = text
			}
	}

	element.title = "Episode " + element.dataset.episodeNumber + " " + airingVerb + " " + startTime + " - " + endTime + " your time"
}