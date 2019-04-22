import AnimeNotifier from "../AnimeNotifier"

// Toggle sidebar
export function toggleSidebar(_: AnimeNotifier) {
	let sidebar = document.getElementById("sidebar") as HTMLElement
	sidebar.classList.toggle("sidebar-visible")
}