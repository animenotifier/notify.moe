import { AnimeNotifier } from "../AnimeNotifier"

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

// Filter anime on maldiff page
export function malDiffFilterAnime(arn: AnimeNotifier, input: HTMLInputElement) {
	let year = arn.app.find("filter-year") as HTMLSelectElement
	let status = arn.app.find("filter-status") as HTMLSelectElement
	let type = arn.app.find("filter-type") as HTMLSelectElement

	arn.app.load(`/editor/anime/maldiff/${year.value}/${status.value}/${type.value}`)
}

// Import Kitsu anime
export async function importKitsuAnime(arn: AnimeNotifier, button: HTMLButtonElement) {
	let response = await fetch(`/api/import/kitsu/anime/${button.dataset.id}`, {
		method: "POST",
		credentials: "same-origin"
	})

	if(response.ok) {
		arn.reloadContent()
	} else {
		arn.statusMessage.showError(await response.text())
	}
}