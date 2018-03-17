import { AnimeNotifier } from "../AnimeNotifier"

// Search page reference
var emptySearchHTML = ""
var searchPage: HTMLElement
var correctResponseRendered = {
	"anime": false,
	"character": false,
	"forum": false,
	"soundtrack": false,
	"user": false,
	"company": false
}

// Save old term to compare
var oldTerm = ""

// Containers for all the search results
var animeSearchResults: HTMLElement
var characterSearchResults: HTMLElement
var forumSearchResults: HTMLElement
var soundtrackSearchResults: HTMLElement
var userSearchResults: HTMLElement
var companySearchResults: HTMLElement

// Search
export async function search(arn: AnimeNotifier, search: HTMLInputElement, e: KeyboardEvent) {
	if(e.ctrlKey || e.altKey) {
		return
	}

	// Check if the search term really changed
	let term = search.value.trim()

	if(term === oldTerm) {
		return
	}

	oldTerm = term

	// Determine if we're already seeing the search page
	let searchPageActivated = (searchPage === arn.app.content.children[0])

	// Reset
	correctResponseRendered.anime = false
	correctResponseRendered.character = false
	correctResponseRendered.forum = false
	correctResponseRendered.soundtrack = false
	correctResponseRendered.user = false

	// Set browser URL
	let url = "/search/" + term
	history.pushState(url, null, url)
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
		}

		if(!animeSearchResults) {
			animeSearchResults = document.getElementById("anime-search-results")
			characterSearchResults = document.getElementById("character-search-results")
			forumSearchResults = document.getElementById("forum-search-results")
			soundtrackSearchResults = document.getElementById("soundtrack-search-results")
			userSearchResults = document.getElementById("user-search-results")
			companySearchResults = document.getElementById("company-search-results")
		}

		if(!term || term.length < 1) {
			await arn.innerHTML(searchPage, emptySearchHTML)
			arn.app.emit("DOMContentLoaded")
			return
		}

		const fetchOptions: RequestInit = {
			credentials: "same-origin"
		}

		// Start searching
		fetch("/_/anime-search/" + term, fetchOptions)
		.then(showResponseInElement(arn, url, "anime", animeSearchResults))
		.catch(console.error)

		fetch("/_/character-search/" + term, fetchOptions)
		.then(showResponseInElement(arn, url, "character", characterSearchResults))
		.catch(console.error)

		fetch("/_/forum-search/" + term, fetchOptions)
		.then(showResponseInElement(arn, url, "forum", forumSearchResults))
		.catch(console.error)

		fetch("/_/soundtrack-search/" + term, fetchOptions)
		.then(showResponseInElement(arn, url, "soundtrack", soundtrackSearchResults))
		.catch(console.error)

		fetch("/_/user-search/" + term, fetchOptions)
		.then(showResponseInElement(arn, url, "user", userSearchResults))
		.catch(console.error)

		fetch("/_/company-search/" + term, fetchOptions)
		.then(showResponseInElement(arn, url, "company", companySearchResults))
		.catch(console.error)
	} catch(err) {
		console.error(err)
	} finally {
		arn.loading(false)
	}
}

function showResponseInElement(arn: AnimeNotifier, url: string, typeName: string, element: HTMLElement) {
	return async response => {
		let html = await response.text()

		if(arn.app.currentPath !== url) {
			// Return if this result would overwrite the already arrived correct result
			if(correctResponseRendered[typeName]) {
				return
			}
		} else {
			correctResponseRendered[typeName] = true
		}

		await arn.innerHTML(element, html)

		// Emit content loaded event
		arn.app.emit("DOMContentLoaded")
	}
}

// Search database
export function searchDB(arn: AnimeNotifier, input: HTMLInputElement, e: KeyboardEvent) {
	if(e.ctrlKey || e.altKey) {
		return
	}

	let dataType = (arn.app.find("data-type") as HTMLInputElement).value || "+"
	let field = (arn.app.find("field") as HTMLInputElement).value || "+"
	let fieldValue = (arn.app.find("field-value") as HTMLInputElement).value || "+"
	let records = arn.app.find("records")

	arn.loading(true)

	fetch(`/api/select/${dataType}/where/${field}/is/${fieldValue}`)
	.then(response => {
		if(response.status !== 200) {
			throw response
		}

		return response
	})
	.then(response => response.json())
	.then(data => {
		records.innerHTML = ""
		let count = 0

		if(data.results.length === 0) {
			records.innerText = "No results."
			return
		}

		for(let record of data.results) {
			count++

			let container = document.createElement("div")
			container.classList.add("record")

			let id = document.createElement("div")
			id.innerText = record.id
			id.classList.add("record-id")
			container.appendChild(id)

			let link = document.createElement("a")
			link.classList.add("record-view")
			link.innerText = "Open " + dataType.toLowerCase()

			if(dataType === "User") {
				link.href = "/+" + record.nick
			} else {
				link.href = "/" + dataType.toLowerCase() + "/" + record.id
			}

			link.target = "_blank"
			container.appendChild(link)

			let apiLink = document.createElement("a")
			apiLink.classList.add("record-view-api")
			apiLink.innerText = "JSON data"
			apiLink.href = "/api/" + dataType.toLowerCase() + "/" + record.id
			apiLink.target = "_blank"
			container.appendChild(apiLink)

			let recordCount = document.createElement("div")
			recordCount.innerText = count + "/" + data.results.length
			recordCount.classList.add("record-count")
			container.appendChild(recordCount)

			records.appendChild(container)
		}
	})
	.catch(response => {
		response.text().then(text => {
			arn.statusMessage.showError(text)
		})
	})
	.then(() => arn.loading(false))
}