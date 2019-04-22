import AnimeNotifier from "../AnimeNotifier"

// Toggle sidebar
export function toggleSidebar(arn: AnimeNotifier) {
	let sidebar = document.getElementById("sidebar") as HTMLElement
	sidebar.classList.toggle("sidebar-visible")
}