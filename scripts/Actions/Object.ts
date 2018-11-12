import AnimeNotifier from "../AnimeNotifier"

// New
export function newObject(arn: AnimeNotifier, button: HTMLButtonElement) {
	let dataType = button.dataset.type

	arn.post(`/api/new/${dataType}`)
	.then(response => response.json())
	.then(obj => arn.app.load(`/${dataType}/${obj.id}/edit`))
	.catch(err => arn.statusMessage.showError(err))
}

// Delete
export function deleteObject(arn: AnimeNotifier, button: HTMLButtonElement) {
	let confirmType = button.dataset.confirmType
	let returnPath = button.dataset.returnPath

	if(!confirm(`Are you sure you want to delete this ${confirmType}?`)) {
		return
	}

	// Double confirmation on anime
	if(confirmType === "anime") {
		if(!confirm(`Just making sure because this is not reversible. Are you absolutely sure you want to delete this ${confirmType}?`)) {
			return
		}
	}

	let endpoint = arn.findAPIEndpoint(button)

	arn.post(endpoint + "/delete")
	.then(() => {
		if(returnPath) {
			arn.app.load(returnPath)
		} else {
			arn.reloadContent()
		}
	})
	.catch(err => arn.statusMessage.showError(err))
}