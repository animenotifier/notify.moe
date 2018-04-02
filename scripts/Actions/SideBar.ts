import AnimeNotifier from "../AnimeNotifier"

// Toggle sidebar
export function toggleSidebar(arn: AnimeNotifier) {
	document.getElementById("sidebar").classList.toggle("sidebar-visible")
}