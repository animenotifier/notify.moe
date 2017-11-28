import { AnimeNotifier } from "../AnimeNotifier"

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
	button.innerText = "Removing..."

	let {animeId, nick} = button.dataset
	let apiEndpoint = arn.findAPIEndpoint(button)

	arn.post(apiEndpoint + "/remove/" + animeId, "")
	.then(() => arn.app.load("/animelist/" + (arn.app.find("Status") as HTMLSelectElement).value))
	.catch(err => arn.statusMessage.showError(err))
}