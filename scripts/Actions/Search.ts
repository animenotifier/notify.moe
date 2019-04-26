import AnimeNotifier from "../AnimeNotifier"
import { delay, requestIdleCallback } from "../Utils"
import Diff from "scripts/Diff";

// Search page reference
var emptySearchHTML = ""
var searchPage: HTMLElement
var searchPageTitle: HTMLElement
var correctResponseRendered = {
	"anime": false,
	"character": false,
	"posts": false,
	"threads": false,
	"soundtrack": false,
	"user": false,
	"amv": false,
	"company": false
}

// Search types
var searchTypes = Object.keys(correctResponseRendered)

// Save old term to compare
var oldTerm = ""

// Containers for all the search results
var results = new Map<string, HTMLElement>()

// Delay before a request is sent
const searchDelay = 140

// Fetch options
const fetchOptions: RequestInit = {
	credentials: "same-origin"
}

// Error message
const searchErrorMessage = "Looks like the website is offline, please retry later."

// Speech recognition
let recognition: SpeechRecognition

// Search
export async function search(arn: AnimeNotifier, search: HTMLInputElement, evt?: KeyboardEvent) {
	if(evt && (evt.ctrlKey || evt.altKey)) {
		return
	}

	// Determine if we're already seeing the search page
	let searchPageActivated = (searchPage === arn.app.content.children[0])

	// Check if the search term really changed
	let term = search.value.trim()

	if(term === oldTerm && searchPageActivated) {
		return
	}

	oldTerm = term

	// Reset
	for(let key of searchTypes) {
		correctResponseRendered[key] = false
	}

	// Set browser URL
	let url = "/search/" + term
	document.title = "Search: " + term
	arn.app.currentPath = url

	// Unmount mountables to improve visual responsiveness on key press
	arn.unmountMountables()

	// Show loading spinner
	arn.loading(true)

	try {
		// Fetch empty search frame if needed
		if(emptySearchHTML === "") {
			let response = await fetch("/_/empty-search")
			emptySearchHTML = await response.text()
		}

		if(!searchPageActivated) {
			if(!searchPage) {
				searchPage = document.createElement("div")
				searchPage.innerHTML = emptySearchHTML
			}

			arn.app.content.innerHTML = ""
			arn.app.content.appendChild(searchPage)

			history.pushState(url, document.title, url)
		} else {
			history.replaceState(url, document.title, url)

			// Delay
			await delay(searchDelay)
		}

		if(term !== search.value.trim()) {
			arn.mountMountables()
			return
		}

		if(!results["anime"]) {
			for(let key of searchTypes) {
				results[key] = document.getElementById(`${key}-search-results`)
			}

			searchPageTitle = document.getElementsByTagName("h1")[0]
		}

		searchPageTitle.textContent = document.title

		if(!term || term.length < 1) {
			await arn.innerHTML(searchPage, emptySearchHTML)
			arn.app.emit("DOMContentLoaded")
			return
		}

		// Start searching anime
		fetch("/_/anime-search/" + term, fetchOptions)
		.then(showResponseInElement(arn, url, "anime", results["anime"]))
		.catch(_ => arn.statusMessage.showError(searchErrorMessage))

		requestIdleCallback(() => {
			// Check that the term hasn't changed in the meantime
			if(term !== search.value.trim()) {
				return
			}

			// Search the other types (everything except anime)
			for(let key of searchTypes) {
				if(key === "anime") {
					continue
				}

				fetch(`/_/${key}-search/` + term, fetchOptions)
				.then(showResponseInElement(arn, url, key, results[key]))
				.catch(_ => arn.statusMessage.showError(searchErrorMessage))
			}
		})
	} catch(err) {
		console.error(err)
	} finally {
		arn.loading(false)
	}
}

function showResponseInElement(arn: AnimeNotifier, url: string, typeName: string, element: HTMLElement) {
	return async (response: Response) => {
		if(!response.ok) {
			throw response.statusText
		}

		let html = await response.text()

		if(html.includes("no-search-results")) {
			Diff.mutations.queue(() => (element.parentElement as HTMLElement).classList.add("search-section-disabled"))
		} else {
			Diff.mutations.queue(() => (element.parentElement as HTMLElement).classList.remove("search-section-disabled"))
		}

		if(arn.app.currentPath !== url) {
			// Return if this result would overwrite the already arrived correct result
			if(correctResponseRendered[typeName]) {
				return
			}
		} else {
			correctResponseRendered[typeName] = true
		}

		await arn.innerHTML(element, html)
		arn.onNewContent(element)
	}
}

export function searchBySpeech(arn: AnimeNotifier, element: HTMLElement) {
	if(recognition) {
		recognition.stop()
		return
	}

	let searchInput = document.getElementById("search") as HTMLInputElement
	let oldPlaceholder = searchInput.placeholder

	let SpeechRecognition: any = window["SpeechRecognition"] || window["webkitSpeechRecognition"]
	recognition = new SpeechRecognition()
	recognition.continuous = false
	recognition.interimResults = false
	recognition.lang = navigator.language

	recognition.onresult = evt => {
		if(evt.results.length > 0) {
			let result = evt.results.item(0).item(0)
			let term = result.transcript

			if(term !== "") {
				searchInput.value = term
				arn.sideBar.hide()
				search(arn, searchInput)
			}
		}

		recognition.stop()
	}

	recognition.onerror = _ => {
		recognition.stop()
	}

	recognition.onend = _ => {
		searchInput.placeholder = oldPlaceholder
		element.classList.remove("speech-listening")
	}

	// Focus search field
	searchInput.placeholder = "Listening..."
	searchInput.value = ""
	searchInput.focus()
	searchInput.select()

	// Highlight microphone icon
	element.classList.add("speech-listening")

	// Start voice recognition
	recognition.start()
}
