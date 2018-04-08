import { delay } from "./Utils"
import StatusMessage from "./StatusMessage"

let etags = new Map<string, string>()
let hasNewVersion = false
let newVersionCheckDelay = 60000

export async function checkNewVersion(url: string, statusMessage: StatusMessage) {
	if(hasNewVersion) {
		return
	}

	try {
		let response = await fetch(url)

		if(!response.ok) {
			console.warn("Error fetch", url)
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
	}

	delay(newVersionCheckDelay).then(() => checkNewVersion(url, statusMessage))
}

