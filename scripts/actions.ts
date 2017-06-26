import { Application } from "./Application"
import { AnimeNotifier } from "./AnimeNotifier"
import { Diff } from "./Diff"
import { delay, findAll } from "./utils"

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

// Diff
export function diff(arn: AnimeNotifier, element: HTMLElement) {
	let url = element.dataset.url || (element as HTMLAnchorElement).getAttribute("href")
	let request = fetch("/_" + url).then(response => response.text())

	history.pushState(url, null, url)
	arn.app.currentPath = url
	arn.app.markActiveLinks()
	arn.loading(true)
	arn.unmountMountables()

	// for(let element of findAll("mountable")) {
	// 	element.classList.remove("mountable")
	// }

	delay(300).then(() => {
		request
		.then(html => arn.app.setContent(html, true))
		.then(() => arn.app.markActiveLinks())
		// .then(() => {
		// 	for(let element of findAll("mountable")) {
		// 		element.classList.remove("mountable")
		// 	}
		// })
		.then(() => arn.app.emit("DOMContentLoaded"))
		.then(() => arn.loading(false))
		.catch(console.error)
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