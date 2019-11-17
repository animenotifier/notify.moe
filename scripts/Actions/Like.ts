import AnimeNotifier from "../AnimeNotifier"

// like
export async function like(arn: AnimeNotifier, element: HTMLElement) {
	arn.statusMessage.showInfo("Liked!", 1000)
	const apiEndpoint = arn.findAPIEndpoint(element)

	try {
		await arn.post(apiEndpoint + "/like")
		arn.reloadContent()
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// unlike
export async function unlike(arn: AnimeNotifier, element: HTMLElement) {
	arn.statusMessage.showInfo("Disliked!", 1000)
	const apiEndpoint = arn.findAPIEndpoint(element)

	try {
		await arn.post(apiEndpoint + "/unlike")
		arn.reloadContent()
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}
