import AnimeNotifier from "../AnimeNotifier"
import Diff from "scripts/Diff"

// Follow user
export async function followUser(arn: AnimeNotifier, element: HTMLElement) {
	return updateFollow(arn, element, "You are now following")
}

// Unfollow user
export async function unfollowUser(arn: AnimeNotifier, element: HTMLElement) {
	return updateFollow(arn, element, "You stopped following")
}

// Update follow
async function updateFollow(arn: AnimeNotifier, element: HTMLElement, message: string) {
	const api = element.dataset.api
	const nick = document.getElementById("nick")

	if(!api || !nick || !nick.textContent) {
		console.error("Missing data-api or invalid nick:", element)
		return
	}

	try {
		await arn.post(api)
		await arn.reloadContent()
		arn.statusMessage.showInfo(`${message} ${nick.textContent}.`)
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}

// Show more
export function showMore(_: AnimeNotifier, showMoreElement: HTMLElement) {
	const elements = [...document.getElementsByClassName("show-more")]

	for(const element of elements) {
		Diff.mutations.queue(() => element.classList.remove("show-more"))
	}

	Diff.mutations.queue(() => showMoreElement.classList.add("show-more"))
}
