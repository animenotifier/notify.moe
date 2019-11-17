import AnimeNotifier from "./AnimeNotifier"
import Application from "./Application"

const app = new Application()
const arn = new AnimeNotifier(app)

arn.init()
