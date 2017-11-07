import { AnimeNotifier } from "../AnimeNotifier"

// Filter anime on explore page
export function filterAnime(arn: AnimeNotifier, input: HTMLInputElement) {
	let year = arn.app.find("filter-year") as HTMLSelectElement
	let status = arn.app.find("filter-status") as HTMLSelectElement
	let type = arn.app.find("filter-type") as HTMLSelectElement

	arn.app.load(`/explore/anime/${year.value}/${status.value}/${type.value}`)
}