import AnimeNotifier from "../AnimeNotifier"

// Toggle sidebar
export function toggleSidebar(_: AnimeNotifier) {
	const sidebar = document.getElementById("sidebar") as HTMLElement
	sidebar.classList.toggle("sidebar-visible")
}
