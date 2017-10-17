import { AnimeNotifier } from "../AnimeNotifier"

// Search
export function search(arn: AnimeNotifier, search: HTMLInputElement, e: KeyboardEvent) {
	if(e.ctrlKey || e.altKey) {
		return
	}

	let term = search.value

	if(!term || term.length < 2) {
		arn.app.content.innerHTML = "Please enter at least 2 characters to start searching."
		return
	}

	arn.diff("/search/" + term)
}