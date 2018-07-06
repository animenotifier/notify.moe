import AnimeNotifier from "../AnimeNotifier"
import { findAll, hideValues } from "scripts/Utils"

// Filter anime on explore page
export function filterAnime(arn: AnimeNotifier, input: HTMLInputElement) {
	let root = document.getElementById("filter-root")

	let elementYear = document.getElementById("filter-year") as HTMLSelectElement
	let elementSeason = document.getElementById("filter-season") as HTMLSelectElement
	let elementStatus = document.getElementById("filter-status") as HTMLSelectElement
	let elementType = document.getElementById("filter-type") as HTMLSelectElement

	for(let element of findAll("anime-grid-image")) {
		let img = element as HTMLImageElement
		img.src = arn.emptyPixel()
		img.classList.remove("element-found")
		img.classList.remove("element-color-preview")
	}

	let year = elementYear.value || "any"
	let season = elementSeason.value || "any"
	let status = elementStatus.value || "any"
	let type = elementType.value || "any"

	arn.diff(`${root.dataset.url}/${year}/${season}/${status}/${type}`)
}

// Toggle hiding added anime.
export function toggleHideAddedAnime(arn: AnimeNotifier, input: HTMLButtonElement) {

	// Toggle state
	let whereAmI
	if(arn.app.currentPath.includes("/explore")) {
		whereAmI = hideValues.hideExplore
		localStorage.setItem(hideValues.hideGenre, "false")
	} else {
		whereAmI = hideValues.hideGenre
		localStorage.setItem(hideValues.hideExplore, "false")
	}

	if(localStorage.getItem(whereAmI) === "true") {
		localStorage.setItem(whereAmI, "false")
	} else {
		localStorage.setItem(whereAmI, "true")
	}

	// Hide anime
	hideAddedAnime(whereAmI)
}

// Hides anime that are already in your list.
export function hideAddedAnime(whereAmI: string) {
	for(let anime of findAll("anime-grid-cell")) {
		if(anime.dataset.added !== "true") {
			continue
		}

		if(localStorage.getItem(whereAmI) === "true") {
			anime.classList.add("anime-grid-cell-hide")
		} else {
			anime.classList.remove("anime-grid-cell-hide")
		}
	}
}

// Hides anime that are not in your list.
export function calendarShowAddedAnimeOnly() {
	for(let anime of findAll("calendar-entry")) {
		if(anime.dataset.added === "false") {
			anime.classList.toggle("hidden")
		}
	}
}