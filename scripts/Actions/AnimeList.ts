import AnimeNotifier from "../AnimeNotifier"

// Add anime to collection
export async function addAnimeToCollection(arn: AnimeNotifier, button: HTMLButtonElement) {
	button.disabled = true

	let {animeId} = button.dataset

	if(!animeId) {
		console.error("Button without anime ID:", button)
		return
	}

	let apiEndpoint = arn.findAPIEndpoint(button)

	try {
		await arn.post(apiEndpoint + "/add/" + animeId)
		arn.reloadContent()

		// Show status message
		let response = await fetch("/api/anime/" + animeId)
		let anime = await response.json()
		arn.statusMessage.showInfo(`Added ${anime.title.canonical} to your collection.`)
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Remove anime from collection
export async function removeAnimeFromCollection(arn: AnimeNotifier, button: HTMLElement) {
	if(!confirm("Are you sure you want to remove it from your collection?")) {
		return
	}

	button.textContent = "Removing..."

	let {animeId, nick} = button.dataset

	if(!animeId || !nick) {
		console.error("Button without nick or anime ID:", button)
		return
	}

	let apiEndpoint = arn.findAPIEndpoint(button)
	let status = document.getElementById("Status") as HTMLSelectElement

	try {
		await arn.post(apiEndpoint + "/remove/" + animeId)
		await arn.app.load(`/+${nick}/animelist/` + status.value)
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}