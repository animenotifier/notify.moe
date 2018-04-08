import { delay, requestIdleCallback } from "./Utils"
import StatusMessage from "./StatusMessage"

let etags = new Map<string, string>()
let hasNewVersion = false
let newVersionCheckDelay = 60000

async function checkNewVersion(url: string, statusMessage: StatusMessage) {
	if(hasNewVersion) {
		return
	}

	try {
		let headers = {}

		if(etags.has(url)) {
			headers["If-None-Match"] = etags.get(url)
		}

		let response = await fetch(url, {
			headers
		})

		// Not modified response
		if(response.status === 304) {
			return
		}

		if(!response.ok) {
			console.warn("Error fetching", url, response.status)
			return
		}

		let newETag = response.headers.get("ETag")
		let oldETag = etags.get(url)

		if(newETag) {
			etags.set(url, newETag)
		}

		if(oldETag && newETag && oldETag !== newETag) {
			statusMessage.showInfo("A new version of the website is available. Please refresh the page.", -1)

			// Do not check for new versions again.
			hasNewVersion = true
			return
		}
	} catch(err) {
		console.warn("Error fetching", url + "\n", err)
	} finally {
		checkNewVersionDelayed(url, statusMessage)
	}
}

export function checkNewVersionDelayed(url: string, statusMessage: StatusMessage) {
	return delay(newVersionCheckDelay).then(() => {
		requestIdleCallback(() => checkNewVersion(url, statusMessage))
	})
}
