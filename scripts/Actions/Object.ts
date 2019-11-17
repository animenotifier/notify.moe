import AnimeNotifier from "../AnimeNotifier"

// New
export async function newObject(arn: AnimeNotifier, button: HTMLButtonElement) {
	const dataType = button.dataset.type

	if(!dataType) {
		console.error("Missing data type:", button)
		return
	}

	try {
		const response = await arn.post(`/api/new/${dataType}`)

		if(!response) {
			throw `Failed creating ${dataType}`
		}

		const json = await response.json()
		await arn.app.load(`/${dataType}/${json.id}/edit`)
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Delete
export async function deleteObject(arn: AnimeNotifier, button: HTMLButtonElement) {
	const confirmType = button.dataset.confirmType
	const returnPath = button.dataset.returnPath

	if(!confirm(`Are you sure you want to delete this ${confirmType}?`)) {
		return
	}

	// Double confirmation on anime
	if(confirmType === "anime") {
		if(!confirm(`Just making sure because this is not reversible. Are you absolutely sure you want to delete this ${confirmType}?`)) {
			return
		}
	}

	const endpoint = arn.findAPIEndpoint(button)

	try {
		await arn.post(endpoint + "/delete")

		if(returnPath) {
			await arn.app.load(returnPath)
		} else {
			await arn.reloadContent()
		}
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}
