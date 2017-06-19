import { Application } from "./Application"
import { AnimeNotifier } from "./AnimeNotifier"

let app = new Application()
let arn = new AnimeNotifier(app)

document.onreadystatechange = function() {
	if(document.readyState === "interactive") {
		arn.run()
	}
}

window.onpopstate = e => {
	if(e.state)
		app.load(e.state, false)
	else if(app.currentURL !== app.originalURL)
		app.load(app.originalURL, false)
}