import AnimeNotifier from "scripts/AnimeNotifier"

// join
export async function join(arn: AnimeNotifier, element: HTMLElement) {
	if(!confirm(`Are you sure you want to join this group?`)) {
		return
	}

	const apiEndpoint = arn.findAPIEndpoint(element)

	try {
		await arn.post(`${apiEndpoint}/join`)
		arn.reloadContent()
		arn.statusMessage.showInfo("Joined group!", 1000)
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// leave
export async function leave(arn: AnimeNotifier, element: HTMLElement) {
	if(!confirm(`Are you sure you want to leave this group?`)) {
		return
	}

	const apiEndpoint = arn.findAPIEndpoint(element)

	try {
		await arn.post(`${apiEndpoint}/leave`)
		arn.reloadContent()
		arn.statusMessage.showInfo("Left group!", 1000)
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}
