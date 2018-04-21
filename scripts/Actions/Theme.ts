import AnimeNotifier from "../AnimeNotifier"

let currentTheme = "light"
let timeoutID: number = 0

const light = {}
const dark = {
	"hue": "45",
	"saturation": "100%",

	"text-color-l": "90%",
	"bg-color": "hsl(0, 0%, 18%)",
	"link-color": "hsl(var(--hue), var(--saturation), 66%)",
	"link-hover-color": "hsl(var(--hue), var(--saturation), 76%)",
	"link-hover-text-shadow": "0 0 8px hsla(var(--hue), var(--saturation), 66%, 0.5)",
	"reverse-light-color": "rgba(255, 255, 255, 0.1)",
	"reverse-light-hover-color": "rgba(255, 255, 255, 0.2)",
	"ui-background": "hsl(0, 0%, 14%)",
	"sidebar-background": "hsla(0, 0%, 0%, 0.2)",
	"sidebar-opaque-background": "hsl(0, 0%, 18%)",
	"table-row-hover-background": "hsla(0, 0%, 100%, 0.01)",

	"theme-white": "var(--bg-color)",
	"theme-black": "var(--text-color)",

	"main-color": "var(--link-color)",
	"link-active-color": "var(--link-hover-color)",
	"button-hover-color": "var(--link-hover-color)",
	"button-hover-background": "hsl(0, 0%, 10%)",
	"tab-background": "hsla(0, 0%, 0%, 0.1)",
	"tab-hover-background": "hsla(0, 0%, 0%, 0.2)",
	"tab-active-color": "hsl(0, 0%, 95%)",
	"tab-active-background": "hsla(0, 0%, 2%, 0.5)",
	"loading-anim-color": "var(--link-color)",
	"anime-alternative-title-color": "hsla(0, 0%, 100%, 0.5)",
	"anime-list-item-name-color": "var(--text-color)",
	"tip-bg-color": "hsl(0, 0%, 10%)",

	"post-like-color": "var(--link-color)",
	"post-unlike-color": "var(--link-color)",
	"post-permalink-color": "var(--link-color)",

	"quote-color": "var(--text-color)",

	"calendar-hover-color": "var(--reverse-light-color)"
}

// Toggle theme
export function toggleTheme(arn: AnimeNotifier) {
	// Clear preview interval
	clearTimeout(timeoutID)

	if(currentTheme === "light") {
		darkTheme(arn)
		return
	}

	lightTheme(arn)
}

// Light theme
export function lightTheme(arn: AnimeNotifier) {
	let root = document.documentElement

	for(let property in light) {
		if(!light.hasOwnProperty(property)) {
			continue
		}

		root.style.setProperty(`--${property}`, light[property])
	}

	currentTheme = "light"
}

// Dark theme
export function darkTheme(arn: AnimeNotifier) {
	let root = document.documentElement

	if(!arn.user || arn.user.dataset.pro !== "true") {
		arn.statusMessage.showInfo("Previewing Dark theme for 30 seconds. If you would like to use it permanently, please support us.", 5000)

		// After 30 seconds, switch back to default theme if the user doesn't own a PRO account
		timeoutID = setTimeout(() => {
			if(currentTheme === "dark") {
				toggleTheme(arn)
				arn.statusMessage.showInfo("Dark theme preview time has ended. If you would like to use it permanently, please support us.", 5000)
			}
		}, 30000)
	}

	for(let property in dark) {
		if(!dark.hasOwnProperty(property)) {
			continue
		}

		if(light[property] === undefined) {
			light[property] = root.style.getPropertyValue(`--${property}`)
		}

		root.style.setProperty(`--${property}`, dark[property])
	}

	currentTheme = "dark"
}