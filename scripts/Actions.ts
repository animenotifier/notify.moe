import { Application } from "./Application"
import { AnimeNotifier } from "./AnimeNotifier"
import { Diff } from "./Diff"
import { findAll } from "./Utils"

// Follow user
export function followUser(arn: AnimeNotifier, elem: HTMLElement) {
	return arn.post(elem.dataset.api, elem.dataset.viewUserId)
	.then(() => arn.reloadContent())
	.then(() => arn.statusMessage.showInfo("You are now following " + arn.app.find("nick").innerText + "."))
	.catch(err => arn.statusMessage.showError(err))
}

// Unfollow user
export function unfollowUser(arn: AnimeNotifier, elem: HTMLElement) {
	return arn.post(elem.dataset.api, elem.dataset.viewUserId)
	.then(() => arn.reloadContent())
	.then(() => arn.statusMessage.showInfo("You stopped following " + arn.app.find("nick").innerText + "."))
	.catch(err => arn.statusMessage.showError(err))
}

// Toggle sidebar
export function toggleSidebar(arn: AnimeNotifier) {
	arn.app.find("sidebar").classList.toggle("sidebar-visible")
}

// Save new data from an input field
export function save(arn: AnimeNotifier, input: HTMLElement) {
	arn.loading(true)

	let isContentEditable = input.isContentEditable
	let obj = {}
	let value = isContentEditable ? input.innerText : (input as HTMLInputElement).value
	
	if((input as HTMLInputElement).type === "number" || input.dataset.type === "number") {
		if(input.getAttribute("step") === "1" || input.dataset.step === "1") {
			obj[input.dataset.field] = parseInt(value)
		} else {
			obj[input.dataset.field] = parseFloat(value)
		}
	} else {
		obj[input.dataset.field] = value
	}

	if(isContentEditable) {
		input.contentEditable = "false"
	} else {
		(input as HTMLInputElement).disabled = true
	}

	let apiEndpoint = arn.findAPIEndpoint(input)

	fetch(apiEndpoint, {
		method: "POST",
		body: JSON.stringify(obj),
		credentials: "same-origin"
	})
	.then(response => response.text())
	.then(body => {
		if(body !== "ok") {
			throw body
		}
	})
	.catch(err => arn.statusMessage.showError(err))
	.then(() => {
		arn.loading(false)

		if(isContentEditable) {
			input.contentEditable = "true"
		} else {
			(input as HTMLInputElement).disabled = false
		}

		return arn.reloadContent()
	})
}

// Close status message
export function closeStatusMessage(arn: AnimeNotifier) {
	arn.statusMessage.close()
}

// Increase episode
export function increaseEpisode(arn: AnimeNotifier, element: HTMLElement) {
	let prev = element.previousSibling as HTMLElement
	let episodes = parseInt(prev.innerText)
	prev.innerText = String(episodes + 1)
	save(arn, prev)
}

// Load
export function load(arn: AnimeNotifier, element: HTMLElement) {
	let url = element.dataset.url || (element as HTMLAnchorElement).getAttribute("href")
	arn.app.load(url)
}

// Soon
export function soon() {
	alert("Coming Soon™")
}

// Diff
export function diff(arn: AnimeNotifier, element: HTMLElement) {
	let url = element.dataset.url || (element as HTMLAnchorElement).getAttribute("href")
	
	arn.diff(url).then(() => arn.scrollTo(element))
}

// Edit post
export function editPost(arn: AnimeNotifier, element: HTMLElement) {
	let postId = element.dataset.id

	let render = arn.app.find("render-" + postId)
	let toolbar = arn.app.find("toolbar-" + postId)
	let title = arn.app.find("title-" + postId)
	let source = arn.app.find("source-" + postId)
	let edit = arn.app.find("edit-toolbar-" + postId)

	render.classList.toggle("hidden")
	toolbar.classList.toggle("hidden")
	source.classList.toggle("hidden")
	edit.classList.toggle("hidden")

	if(title) {
		title.classList.toggle("hidden")
	}
}

// Save post
export function savePost(arn: AnimeNotifier, element: HTMLElement) {
	let postId = element.dataset.id
	let source = arn.app.find("source-" + postId) as HTMLTextAreaElement
	let title = arn.app.find("title-" + postId) as HTMLInputElement
	let text = source.value

	let updates: any = {
		Text: text,
	}

	// Add title for threads only
	if(title) {
		updates.Title = title.value
	}

	let apiEndpoint = arn.findAPIEndpoint(element)

	arn.post(apiEndpoint, updates)
	.then(() => arn.reloadContent())
	.catch(err => arn.statusMessage.showError(err))
}

// like
export function like(arn: AnimeNotifier, element: HTMLElement) {
	let apiEndpoint = arn.findAPIEndpoint(element)

	arn.post(apiEndpoint + "/like", null)
	.then(() => arn.reloadContent())
	.catch(err => arn.statusMessage.showError(err))
}

// unlike
export function unlike(arn: AnimeNotifier, element: HTMLElement) {
	let apiEndpoint = arn.findAPIEndpoint(element)

	arn.post(apiEndpoint + "/unlike", null)
	.then(() => arn.reloadContent())
	.catch(err => arn.statusMessage.showError(err))
}

// Forum reply
export function forumReply(arn: AnimeNotifier) {
	let textarea = arn.app.find("new-reply") as HTMLTextAreaElement
	let thread = arn.app.find("thread")

	let post = {
		text: textarea.value,
		threadId: thread.dataset.id,
		tags: []
	}

	arn.post("/api/new/post", post)
	.then(() => arn.reloadContent())
	.then(() => textarea.value = "")
	.catch(err => arn.statusMessage.showError(err))
}

// Create thread
export function createThread(arn: AnimeNotifier) {
	let title = arn.app.find("title") as HTMLInputElement
	let text = arn.app.find("text") as HTMLTextAreaElement
	let category = arn.app.find("tag") as HTMLInputElement

	let thread = {
		title: title.value,
		text: text.value,
		tags: [category.value]
	}

	arn.post("/api/new/thread", thread)
	.then(() => arn.app.load("/forum/" + thread.tags[0]))
	.catch(err => arn.statusMessage.showError(err))
}

// Create soundtrack
export function createSoundTrack(arn: AnimeNotifier, button: HTMLButtonElement) {
	let soundcloud = arn.app.find("soundcloud-link") as HTMLInputElement
	let youtube = arn.app.find("youtube-link") as HTMLInputElement
	let anime = arn.app.find("anime-link") as HTMLInputElement
	let osu = arn.app.find("osu-link") as HTMLInputElement

	let soundtrack = {
		soundcloud: soundcloud.value,
		youtube: youtube.value,
		tags: [anime.value, osu.value],
	}

	arn.post("/api/new/soundtrack", soundtrack)
	.then(() => arn.app.load("/soundtracks"))
	.catch(err => arn.statusMessage.showError(err))
}

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

// Enable notifications
export async function enableNotifications(arn: AnimeNotifier, button: HTMLElement) {
	await arn.pushManager.subscribe(arn.user.dataset.id)
	arn.updatePushUI()
}

// Disable notifications
export async function disableNotifications(arn: AnimeNotifier, button: HTMLElement) {
	await arn.pushManager.unsubscribe(arn.user.dataset.id)
	arn.updatePushUI()
}

// Test notification
export function testNotification(arn: AnimeNotifier) {
	fetch("/api/test/notification", {
		credentials: "same-origin"
	})
}

// Add anime to collection
export function addAnimeToCollection(arn: AnimeNotifier, button: HTMLElement) {
	button.innerText = "Adding..."
	arn.loading(true)

	let {animeId, userId, userNick} = button.dataset

	fetch("/api/animelist/" + userId + "/add", {
		method: "POST",
		body: animeId,
		credentials: "same-origin"
	})
	.then(response => response.text())
	.then(body => {
		if(body !== "ok") {
			throw body
		}
		
		return arn.reloadContent()
	})
	.catch(err => arn.statusMessage.showError(err))
	.then(() => arn.loading(false))
}

// Remove anime from collection
export function removeAnimeFromCollection(arn: AnimeNotifier, button: HTMLElement) {
	button.innerText = "Removing..."
	arn.loading(true)

	let {animeId, userId, userNick} = button.dataset

	fetch("/api/animelist/" + userId + "/remove", {
		method: "POST",
		body: animeId,
		credentials: "same-origin"
	})
	.then(response => response.text())
	.then(body => {
		if(body !== "ok") {
			throw body
		}
		
		return arn.app.load("/+" + userNick + "/animelist/" + (arn.app.find("Status") as HTMLSelectElement).value)
	})
	.catch(err => arn.statusMessage.showError(err))
	.then(() => arn.loading(false))
}

// Charge up
export function chargeUp(arn: AnimeNotifier, button: HTMLElement) {
	let amount = button.dataset.amount

	arn.loading(true)
	arn.statusMessage.showInfo("Creating PayPal transaction... This might take a few seconds.")

	fetch("/api/paypal/payment/create", {
		method: "POST",
		body: amount,
		credentials: "same-origin"
	})
	.then(response => response.json())
	.then(payment => {
		if(!payment || !payment.links) {
			throw "Error creating PayPal payment"
		}

		console.log(payment)
		let link = payment.links.find(link => link.rel === "approval_url")

		if(!link) {
			throw "Error finding PayPal payment link"
		}

		arn.statusMessage.showInfo("Redirecting to PayPal...", 5000)

		let url = link.href
		window.location.href = url
	})
	.catch(err => arn.statusMessage.showError(err))
	.then(() => arn.loading(false))
}

// Buy item
export function buyItem(arn: AnimeNotifier, button: HTMLElement) {
	let itemId = button.dataset.itemId
	let itemName = button.dataset.itemName
	let price = button.dataset.price

	if(!confirm(`Would you like to buy ${itemName} for ${price} gems?`)) {
		return
	}

	arn.loading(true)

	fetch(`/api/shop/buy/${itemId}/1`, {
		method: "POST",
		credentials: "same-origin"
	})
	.then(response => response.text())
	.then(body => {
		if(body !== "ok") {
			throw body
		}
		
		return arn.reloadContent()
	})
	.then(() => arn.statusMessage.showInfo(`You bought ${itemName} for ${price} gems. Check out your inventory to confirm the purchase.`, 4000))
	.catch(err => arn.statusMessage.showError(err))
	.then(() => arn.loading(false))
}

// Chrome extension installation
export function installExtension(arn: AnimeNotifier, button: HTMLElement) {
	let browser: any = window["chrome"]
	browser.webstore.install()
}

// Desktop app installation
export function installApp() {
	alert("Open your browser menu > 'More tools' > 'Add to desktop' and enable 'Open as window'.")
}