import AnimeNotifier from "../AnimeNotifier"

// Publish
export async function publish(arn: AnimeNotifier, button: HTMLButtonElement) {
	const endpoint = arn.findAPIEndpoint(button)

	try {
		await arn.post(endpoint + "/publish")
		await arn.app.load(arn.app.currentPath.replace("/edit", ""))
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Unpublish
export async function unpublish(arn: AnimeNotifier, button: HTMLButtonElement) {
	const endpoint = arn.findAPIEndpoint(button)

	try {
		await arn.post(endpoint + "/unpublish")
		await arn.reloadContent()
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}
