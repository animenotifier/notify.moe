import { aero as app } from "./Aero/Aero"

class AnimeNotifier {
	constructor() {
		app.content = app.find("content")
		app.loading = app.find("loading")
		app.run()
	}
}

document.onreadystatechange = function() {
	if(document.readyState === "interactive") {
		let arn = new AnimeNotifier()
	}
}

window.onpopstate = e => {
	if(e.state)
		app.load(e.state, false)
	else if(app.currentURL !== app.originalURL)
		app.load(app.originalURL, false)
}