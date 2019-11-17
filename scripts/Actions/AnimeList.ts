import AnimeNotifier from "../AnimeNotifier"

// Add anime to collection
export async function addAnimeToCollection(arn: AnimeNotifier, button: HTMLButtonElement) {
	button.disabled = true

	const {animeId} = button.dataset

	if(!animeId) {
		console.error("Button without anime ID:", button)
		return
	}

	const apiEndpoint = arn.findAPIEndpoint(button)

	try {
		await arn.post(apiEndpoint + "/add/" + animeId)
		arn.reloadContent()

		// Show status message
		const response = await fetch("/api/anime/" + animeId)
		const anime = await response.json()
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
	const {animeId, nick} = button.dataset

	if(!animeId || !nick) {
		console.error("Button without nick or anime ID:", button)
		return
	}

	const apiEndpoint = arn.findAPIEndpoint(button)
	const status = document.getElementById("Status") as HTMLSelectElement

	try {
		await arn.post(apiEndpoint + "/remove/" + animeId)
		await arn.app.load(`/+${nick}/animelist/` + status.value)
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Delete anime list
export async function deleteAnimeList(arn: AnimeNotifier, button: HTMLElement) {
	if(!confirm("Last confirmation: Are you sure you want to delete your entire anime list?")) {
		return
	}

	button.textContent = "Deleting..."
	const {returnPath} = button.dataset

	if(!returnPath) {
		console.error("Button without data-return-path:", button)
		return
	}

	try {
		await arn.post("/api/delete/animelist")
		await arn.app.load(returnPath)
		arn.statusMessage.showInfo("Your anime list has been deleted.")
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}
