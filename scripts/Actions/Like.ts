import { AnimeNotifier } from "../AnimeNotifier"

// like
export function like(arn: AnimeNotifier, element: HTMLElement) {
	let apiEndpoint = arn.findAPIEndpoint(element)

	arn.post(apiEndpoint + "/like", null)
	.then(() => arn.reloadContent())
	.catch(err => arn.statusMessage.showError(err))
}

// unlike
export function unlike(arn: AnimeNotifier, element: HTMLElement) {
	let apiEndpoint = arn.findAPIEndpoint(element)

	arn.post(apiEndpoint + "/unlike", null)
	.then(() => arn.reloadContent())
	.catch(err => arn.statusMessage.showError(err))
}