import AnimeNotifier from "../AnimeNotifier"

// newAnimeDiffIgnore
export function newAnimeDiffIgnore(arn: AnimeNotifier, button: HTMLButtonElement) {
	if(!confirm("Are you sure you want to permanently ignore this difference?")) {
		return
	}

	let id = button.dataset.id
	let hash = button.dataset.hash

	arn.post(`/api/new/ignoreanimedifference`, {
		id,
		hash
	})
	.then(() => {
		arn.reloadContent()
	})
	.catch(err => arn.statusMessage.showError(err))
}

// Import Kitsu anime
export async function importKitsuAnime(arn: AnimeNotifier, button: HTMLButtonElement) {
	let newTab = window.open()
	let animeId = button.dataset.id
	let response = await fetch(`/api/import/kitsu/anime/${animeId}`, {
		method: "POST",
		credentials: "same-origin"
	})

	if(response.ok) {
		newTab.location.href = `/anime/${animeId}`
		arn.reloadContent()
	} else {
		arn.statusMessage.showError(await response.text())
	}
}