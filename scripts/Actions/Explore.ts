import { AnimeNotifier } from "../AnimeNotifier"
import { findAll } from "scripts/Utils";

// Filter anime on explore page
export function filterAnime(arn: AnimeNotifier, input: HTMLInputElement) {
	let year = arn.app.find("filter-year") as HTMLSelectElement
	let status = arn.app.find("filter-status") as HTMLSelectElement
	let type = arn.app.find("filter-type") as HTMLSelectElement

	arn.app.load(`/explore/anime/${year.value}/${status.value}/${type.value}`)
}

// Hides anime that are already in your list.
export function hideAddedAnime() {
	for(let anime of findAll("anime-grid-cell")) {
		if(anime.dataset.added === "true") {
			anime.classList.toggle("anime-grid-cell-hide")
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