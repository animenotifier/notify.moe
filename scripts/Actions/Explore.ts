import emptyPixel from "scripts/Utils/emptyPixel"
import findAll from "scripts/Utils/findAll"
import AnimeNotifier from "../AnimeNotifier"

// Filter anime on explore page
export function filterAnime(arn: AnimeNotifier, _: HTMLInputElement) {
	const root = document.getElementById("filter-root") as HTMLElement

	const elementYear = document.getElementById("filter-year") as HTMLSelectElement
	const elementSeason = document.getElementById("filter-season") as HTMLSelectElement
	const elementStatus = document.getElementById("filter-status") as HTMLSelectElement
	const elementType = document.getElementById("filter-type") as HTMLSelectElement

	for(const element of findAll("anime-grid-image")) {
		const img = element as HTMLImageElement
		img.src = emptyPixel
		img.classList.remove("element-found")
		img.classList.remove("element-color-preview")
	}

	const year = elementYear.value || "any"
	const season = elementSeason.value || "any"
	const status = elementStatus.value || "any"
	const type = elementType.value || "any"

	arn.diff(`${root.dataset.url}/${year}/${season}/${status}/${type}`)
}

// Toggle hiding added anime.
export function toggleHideAddedAnime() {
	hideAddedAnime()
}

// Hides anime that are already in your list.
export function hideAddedAnime() {
	for(const anime of findAll("anime-grid-cell")) {
		if(anime.dataset.added !== "true") {
			continue
		}

		anime.classList.toggle("anime-grid-cell-hide")
	}
}

// Hides anime that are not in your list.
export async function calendarShowAddedAnimeOnly(arn: AnimeNotifier, element: HTMLInputElement) {
	const calendar = document.getElementById("calendar")

	if(!calendar || calendar.dataset.showAddedAnimeOnly === undefined) {
		return
	}

	// Toggling the switch will trigger the CSS rules
	if(calendar.dataset.showAddedAnimeOnly === "true") {
		calendar.dataset.showAddedAnimeOnly = "false"
	} else {
		calendar.dataset.showAddedAnimeOnly = "true"
	}

	// Save the state in the database
	const showAddedAnimeOnly = calendar.dataset.showAddedAnimeOnly === "true"
	const apiEndpoint = arn.findAPIEndpoint(element)

	try {
		await arn.post(apiEndpoint, {
			"Calendar.ShowAddedAnimeOnly": showAddedAnimeOnly
		})
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}
