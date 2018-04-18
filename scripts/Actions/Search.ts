import AnimeNotifier from "../AnimeNotifier"
import { delay, requestIdleCallback, findAllInside } from "../Utils"

// Search page reference
var emptySearchHTML = ""
var searchPage: HTMLElement
var searchPageTitle: HTMLElement
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

// Delay before a request is sent
const searchDelay = 140

// Fetch options
const fetchOptions: RequestInit = {
	credentials: "same-origin"
}

// Search
export async function search(arn: AnimeNotifier, search: HTMLInputElement, e: KeyboardEvent) {
	if(e.ctrlKey || e.altKey) {
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
	correctResponseRendered.anime = false
	correctResponseRendered.character = false
	correctResponseRendered.forum = false
	correctResponseRendered.soundtrack = false
	correctResponseRendered.user = false
	correctResponseRendered.company = false

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

		if(!animeSearchResults) {
			animeSearchResults = document.getElementById("anime-search-results")
			characterSearchResults = document.getElementById("character-search-results")
			forumSearchResults = document.getElementById("forum-search-results")
			soundtrackSearchResults = document.getElementById("soundtrack-search-results")
			userSearchResults = document.getElementById("user-search-results")
			companySearchResults = document.getElementById("company-search-results")
			searchPageTitle = document.getElementsByTagName("h1")[0]
		}

		searchPageTitle.innerText = document.title

		if(!term || term.length < 1) {
			await arn.innerHTML(searchPage, emptySearchHTML)
			arn.app.emit("DOMContentLoaded")
			return
		}

		// Start searching
		fetch("/_/anime-search/" + term, fetchOptions)
		.then(showResponseInElement(arn, url, "anime", animeSearchResults))
		.catch(console.error)

		requestIdleCallback(() => {
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
		})
	} catch(err) {
		console.error(err)
	} finally {
		arn.loading(false)
	}
}

function showResponseInElement(arn: AnimeNotifier, url: string, typeName: string, element: HTMLElement) {
	return async (response: Response) => {
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

		showSearchResults(arn, element)
	}
}

export function showSearchResults(arn: AnimeNotifier, element: HTMLElement) {
	// Do the same as for the content loaded event,
	// except here we are limiting it to the element.
	arn.app.ajaxify(element.getElementsByTagName("a"))
	arn.lazyLoad(findAllInside("lazy", element))
	arn.mountMountables(findAllInside("mountable", element))
	arn.assignTooltipOffsets(findAllInside("tip", element))
}
