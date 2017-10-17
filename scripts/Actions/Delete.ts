import { AnimeNotifier } from "../AnimeNotifier"

// Delete
export function deleteObject(arn: AnimeNotifier, button: HTMLButtonElement) {
	let confirmType = button.dataset.confirmType
	let returnPath = button.dataset.returnPath

	if(!confirm(`Are you sure you want to delete this ${confirmType}?`)) {
		return
	}
	
	let endpoint = arn.findAPIEndpoint(button)

	arn.post(endpoint + "/delete", "")
	.then(() => arn.app.load(returnPath))
	.catch(err => arn.statusMessage.showError(err))
}