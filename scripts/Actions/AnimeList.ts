import { AnimeNotifier } from "../AnimeNotifier"

// Add anime to collection
export function addAnimeToCollection(arn: AnimeNotifier, button: HTMLElement) {
	button.innerText = "Adding..."
	
	let {animeId} = button.dataset
	let apiEndpoint = arn.findAPIEndpoint(button)

	arn.post(apiEndpoint + "/add/" + animeId, "")
	.then(() => arn.reloadContent())
	.catch(err => arn.statusMessage.showError(err))
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