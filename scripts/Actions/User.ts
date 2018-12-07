import AnimeNotifier from "../AnimeNotifier"
import Diff from "scripts/Diff"

// Follow user
export async function followUser(arn: AnimeNotifier, element: HTMLElement) {
	try {
		await arn.post(element.dataset.api)
		await arn.reloadContent()
		arn.statusMessage.showInfo("You are now following " + document.getElementById("nick").textContent + ".")
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Unfollow user
export async function unfollowUser(arn: AnimeNotifier, element: HTMLElement) {
	try {
		await arn.post(element.dataset.api)
		await arn.reloadContent()
		arn.statusMessage.showInfo("You stopped following " + document.getElementById("nick").textContent + ".")
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Show more
export function showMore(arn: AnimeNotifier, showMoreElement: HTMLElement) {
	const elements = [...document.getElementsByClassName("show-more")]

	for(let element of elements) {
		Diff.mutations.queue(() => element.classList.remove("show-more"))
	}

	Diff.mutations.queue(() => showMoreElement.classList.add("show-more"))
}