import { Application } from "./Application"
import { AnimeNotifier } from "./AnimeNotifier"

let app = new Application()
let arn = new AnimeNotifier(app)

document.addEventListener("DOMContentLoaded", arn.onContentLoaded.bind(arn))
document.addEventListener("readystatechange", arn.onReadyStateChange.bind(arn))

window.onpopstate = e => {
	if(e.state)
		app.load(e.state, false)
	else if(app.currentPath !== app.originalPath)
		app.load(app.originalPath, false)
}