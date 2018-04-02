import Application from "./Application"
import AnimeNotifier from "./AnimeNotifier"

let app = new Application()
let arn = new AnimeNotifier(app)

arn.init()

// For debugging purposes
window["arn"] = arn