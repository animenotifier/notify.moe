import { Application } from "./Application"
import { AnimeNotifier } from "./AnimeNotifier"

let app = new Application()
let arn = new AnimeNotifier(app)

document.addEventListener("DOMContentLoaded", arn.onContentLoaded.bind(arn))
document.addEventListener("readystatechange", arn.onReadyStateChange.bind(arn))
document.addEventListener("keydown", arn.onKeyDown.bind(arn), false)

window.addEventListener("popstate", arn.onPopState.bind(arn))
// window.addEventListener("resize", arn.onResize.bind(arn))