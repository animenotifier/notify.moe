import AnimeNotifier from "../AnimeNotifier"
import plural from "../Utils/plural"
import ServerEvent from "./ServerEvent"

const reconnectDelay = 3000

let supported: boolean
let eventSource: EventSource
let arn: AnimeNotifier
let etags: Map<string, string>

export default function receiveServerEvents(animeNotifier: AnimeNotifier) {
	supported = ("EventSource" in window)

	if(!supported) {
		return
	}

	arn = animeNotifier
	etags = new Map<string, string>()
	connect()
}

function connect() {
	if(eventSource) {
		eventSource.close()
	}

	eventSource = new EventSource("/api/sse/events", {
		withCredentials: true
	})

	eventSource.addEventListener("ping", (e: any) => ping(e))
	eventSource.addEventListener("etag", (e: any) => etag(e))
	eventSource.addEventListener("post activity", (e: any) => activity(e, "post"))
	eventSource.addEventListener("watch activity", (e: any) => activity(e, "watch"))
	eventSource.addEventListener("notificationCount", (e: any) => notificationCount(e))

	eventSource.onerror = () => {
		setTimeout(() => connect(), reconnectDelay)
	}
}

function ping(_: ServerEvent) {
	console.log("ping")
}

function etag(e: ServerEvent) {
	const data = JSON.parse(e.data)
	const oldETag = etags.get(data.url)
	const newETag = data.etag

	if(oldETag && newETag && oldETag !== newETag) {
		arn.statusMessage.showInfo("A new version of the website is available. Please refresh the page.", -1)
	}

	etags.set(data.url, newETag)
}

function activity(e: ServerEvent, activityType: string) {
	if(activityType === "post" && !location.pathname.endsWith("/activity")) {
		return
	}

	if(activityType === "watch" && !location.pathname.endsWith("/activity/watch")) {
		return
	}

	const showFollowedOnlyButton = document.getElementById("Activity.ShowFollowedOnly")

	if(!showFollowedOnlyButton) {
		return
	}

	const showFollowedOnly = (showFollowedOnlyButton.dataset.action === "disable")
	const isFollowingUser = JSON.parse(e.data)

	// If we're on the followed only feed and we receive an activity
	// about a user we don't follow, ignore the message.
	if(showFollowedOnly && !isFollowingUser) {
		return
	}

	const button = document.getElementById("load-new-activities")

	if(!button || !button.dataset.count) {
		return
	}

	const buttonText = document.getElementById("load-new-activities-text")

	if(!buttonText) {
		return
	}

	const newCount = parseInt(button.dataset.count, 10) + 1
	button.dataset.count = newCount.toString()
	buttonText.textContent = plural(newCount, "new activity")
}

function notificationCount(e: ServerEvent) {
	if(!arn.notificationManager) {
		return
	}

	arn.notificationManager.setCounter(parseInt(e.data, 10))
}
