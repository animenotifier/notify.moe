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