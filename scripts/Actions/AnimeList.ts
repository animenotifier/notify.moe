import AnimeNotifier from "../AnimeNotifier"

// Add anime to collection
export async function addAnimeToCollection(arn: AnimeNotifier, button: HTMLButtonElement) {
	button.disabled = true

	let {animeId} = button.dataset
	let apiEndpoint = arn.findAPIEndpoint(button)

	try {
		await arn.post(apiEndpoint + "/add/" + animeId, "")
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
export function removeAnimeFromCollection(arn: AnimeNotifier, button: HTMLElement) {
	if(!confirm("Are you sure you want to remove it from your collection?")) {
		return
	}

	button.innerText = "Removing..."

	let {animeId, nick} = button.dataset
	let apiEndpoint = arn.findAPIEndpoint(button)

	arn.post(apiEndpoint + "/remove/" + animeId, "")
	.then(() => arn.app.load(`/+${nick}/animelist/` + (document.getElementById("Status") as HTMLSelectElement).value))
	.catch(err => arn.statusMessage.showError(err))
}