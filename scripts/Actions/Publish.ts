import AnimeNotifier from "../AnimeNotifier"

// Publish
export function publish(arn: AnimeNotifier, button: HTMLButtonElement) {
	let endpoint = arn.findAPIEndpoint(button)

	arn.post(endpoint + "/publish", "")
	.then(() => arn.app.load(arn.app.currentPath.replace("/edit", "")))
	.catch(err => arn.statusMessage.showError(err))
}

// Unpublish
export function unpublish(arn: AnimeNotifier, button: HTMLButtonElement) {
	let endpoint = arn.findAPIEndpoint(button)

	arn.post(endpoint + "/unpublish", "")
	.then(() => arn.reloadContent())
	.catch(err => arn.statusMessage.showError(err))
}