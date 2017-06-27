import { Application } from "./Application"
import { AnimeNotifier } from "./AnimeNotifier"
import { Diff } from "./Diff"

// Save new data from an input field
export function save(arn: AnimeNotifier, input: HTMLInputElement | HTMLTextAreaElement) {
	arn.loading(true)

	let obj = {}
	let value = input.value
	
	if(input.type === "number") {
		if(input.getAttribute("step") === "1") {
			obj[input.id] = parseInt(value)
		} else {
			obj[input.id] = parseFloat(value)
		}
	} else {
		obj[input.id] = value
	}

	// console.log(input.type, input.dataset.api, obj, JSON.stringify(obj))

	let apiObject: HTMLElement
	let parent = input as HTMLElement

	while(parent = parent.parentElement) {
		if(parent.dataset.api !== undefined) {
			apiObject = parent
			break
		}
	}

	if(!apiObject) {
		throw "API object not found"
	}

	input.disabled = true

	fetch(apiObject.dataset.api, {
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
	.catch(console.error)
	.then(() => {
		arn.loading(false)
		input.disabled = false

		return arn.reloadContent()
	})
}

// Load
export function load(arn: AnimeNotifier, element: HTMLElement) {
	let url = element.dataset.url || (element as HTMLAnchorElement).getAttribute("href")
	arn.app.load(url)
}

// Diff
export function diff(arn: AnimeNotifier, element: HTMLElement) {
	let url = element.dataset.url || (element as HTMLAnchorElement).getAttribute("href")
	arn.diff(url)
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

	arn.post("/api/post/new", post)
	.then(() => arn.reloadContent())
	.then(() => textarea.value = "")
	.catch(console.error)
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

	arn.post("/api/thread/new", thread)
	.then(() => arn.app.load("/forum/" + thread.tags[0]))
	.catch(console.error)
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

	button.innerText = "Adding..."
	button.disabled = true

	arn.post("/api/soundtrack/new", soundtrack)
	.then(() => arn.app.load("/music"))
	.catch(err => {
		console.error(err)
		arn.reloadContent()
	})
}

// Search
export function search(arn: AnimeNotifier, search: HTMLInputElement, e: KeyboardEvent) {
	if(e.ctrlKey || e.altKey) {
		return
	}

	let term = search.value

	if(!window.location.pathname.startsWith("/search/")) {
		history.pushState("search", null, "/search/" + term)
	} else {
		history.replaceState("search", null, "/search/" + term)
	}

	if(!term || term.length < 2) {
		arn.app.content.innerHTML = "Please enter at least 2 characters to start searching."
		return
	}

	var results = arn.app.find("results")

	if(!results) {
		results = document.createElement("div")
		results.id = "results"
		arn.app.content.innerHTML = ""
		arn.app.content.appendChild(results)
	}

	arn.app.get("/_/search/" + encodeURI(term))
	.then(html => {
		if(!search.value) {
			return
		}

		Diff.innerHTML(results, html)
		arn.app.emit("DOMContentLoaded")
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
	.catch(console.error)
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
		
		return arn.app.load("/+" + userNick + "/animelist")
	})
	.catch(console.error)
	.then(() => arn.loading(false))
}

// Chrome extension installation
export function installExtension(arn: AnimeNotifier, button: HTMLElement) {
	let browser: any = window["chrome"]
	browser.webstore.install()
}