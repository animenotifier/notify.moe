import { AnimeNotifier } from "../AnimeNotifier"

// Load
export function load(arn: AnimeNotifier, element: HTMLElement) {
	let url = element.dataset.url || (element as HTMLAnchorElement).getAttribute("href")
	arn.app.load(url)
}

// Diff
export function diff(arn: AnimeNotifier, element: HTMLElement) {
	let url = element.dataset.url || (element as HTMLAnchorElement).getAttribute("href")
	
	arn.diff(url).then(() => arn.scrollTo(element))
}