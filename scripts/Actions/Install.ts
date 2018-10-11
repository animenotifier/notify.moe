import AnimeNotifier from "../AnimeNotifier"

// Chrome extension installation
export function installExtension(arn: AnimeNotifier, button: HTMLElement) {
	let browser: any = window["chrome"]

	if(browser && browser.webstore) {
		browser.webstore.install()
	} else {
		window.open("https://chrome.google.com/webstore/detail/anime-notifier/hajchfikckiofgilinkpifobdbiajfch", "_blank")
	}
}

// Desktop app installation
export function installApp() {
	alert("Open your browser menu > 'Install Anime Notifier'.")
}