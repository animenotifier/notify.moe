import AnimeNotifier from "../AnimeNotifier"

// Chrome extension installation
export function installExtension(arn: AnimeNotifier, button: HTMLElement) {
	let browser: any = window["chrome"]
	browser.webstore.install()
}

// Desktop app installation
export function installApp() {
	alert("Open your browser menu > 'More tools' > 'Add to desktop' and enable 'Open as window'.")
}