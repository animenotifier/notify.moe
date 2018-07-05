import AnimeNotifier from "../AnimeNotifier"
import { findAll } from "scripts/Utils";

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

// Hides anime that are already in your list. And "toggles" localStorage.hide if button is pressed
export function hideAddedAnime(arn?: AnimeNotifier, input?: HTMLButtonElement) {
	if(input) {
		if(localStorage.getItem("hideAdded") === "true") {
			localStorage.setItem("hideAdded", "false")
		} else {
			localStorage.setItem("hideAdded", "true")
		}
	}

	if(localStorage.getItem("hideAdded") === "true" || input.type === "submit") {
		for(let anime of findAll("anime-grid-cell")) {
			if(anime.dataset.added === "true") {
				anime.classList.toggle("anime-grid-cell-hide")
			}
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