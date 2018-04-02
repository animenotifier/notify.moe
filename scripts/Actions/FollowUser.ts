import AnimeNotifier from "../AnimeNotifier"

// Follow user
export function followUser(arn: AnimeNotifier, elem: HTMLElement) {
	return arn.post(elem.dataset.api, "")
	.then(() => arn.reloadContent())
	.then(() => arn.statusMessage.showInfo("You are now following " + document.getElementById("nick").innerText + "."))
	.catch(err => arn.statusMessage.showError(err))
}

// Unfollow user
export function unfollowUser(arn: AnimeNotifier, elem: HTMLElement) {
	return arn.post(elem.dataset.api, "")
	.then(() => arn.reloadContent())
	.then(() => arn.statusMessage.showInfo("You stopped following " + document.getElementById("nick").innerText + "."))
	.catch(err => arn.statusMessage.showError(err))
}