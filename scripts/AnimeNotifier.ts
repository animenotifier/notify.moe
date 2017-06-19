import { Application } from "./Application"

export class AnimeNotifier {
	app: Application

	constructor(app: Application) {
		this.app = app
	}

	run() {
		this.app.content = this.app.find("content")
		this.app.loading = this.app.find("loading")
		this.app.run()
	}
}