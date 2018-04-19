import AnimeNotifier from "../AnimeNotifier"

// like
export async function like(arn: AnimeNotifier, element: HTMLElement) {
	arn.statusMessage.showInfo("Liked!", 1000)

	let apiEndpoint = arn.findAPIEndpoint(element)
	await arn.post(apiEndpoint + "/like", null).catch(err => arn.statusMessage.showError(err))
	arn.reloadContent()
}

// unlike
export async function unlike(arn: AnimeNotifier, element: HTMLElement) {
	arn.statusMessage.showInfo("Disliked!", 1000)

	let apiEndpoint = arn.findAPIEndpoint(element)
	await arn.post(apiEndpoint + "/unlike", null).catch(err => arn.statusMessage.showError(err))
	arn.reloadContent()
}