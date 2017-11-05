import { AnimeNotifier } from "../AnimeNotifier"

// Toggle sidebar
export function toggleSidebar(arn: AnimeNotifier) {
	arn.app.find("sidebar").classList.toggle("sidebar-visible")
}