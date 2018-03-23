import { AnimeNotifier } from "../AnimeNotifier"

let currentTheme = "light"

const light = {}
const dark = {
	"hue": "45",
	"saturation": "100%",

	"text-color": "hsl(0, 0%, 90%)",
	"bg-color": "hsl(0, 0%, 24%)",
	"link-color": "hsl(var(--hue), var(--saturation), 66%)",
	"link-hover-color": "hsl(var(--hue), var(--saturation), 76%)",
	"link-hover-text-shadow": "0 0 8px hsla(var(--hue), var(--saturation), 66%, 0.5)",
	"reverse-light-color": "rgba(255, 255, 255, 0.1)",
	"reverse-light-hover-color": "rgba(255, 255, 255, 0.2)",
	"ui-background": "hsl(0, 0%, 18%)",
	"sidebar-background": "hsla(0, 0%, 0%, 0.2)",
	"sidebar-opaque-background": "hsl(0, 0%, 18%)",
	"table-row-hover-background": "hsla(0, 0%, 100%, 0.01)",

	"theme-white": "var(--bg-color)",
	"theme-black": "var(--text-color)",

	"main-color": "var(--link-color)",
	"link-active-color": "var(--link-hover-color)",
	"button-hover-color": "var(--link-hover-color)",
	"button-hover-background": "hsl(0, 0%, 14%)",
	"tab-background": "hsla(0, 0%, 0%, 0.1)",
	"tab-hover-background": "hsla(0, 0%, 0%, 0.2)",
	"tab-active-color": "hsl(0, 0%, 95%)",
	"tab-active-background": "hsla(0, 0%, 2%, 0.5)",
	"loading-anim-color": "var(--link-color)",
	"anime-alternative-title-color": "hsla(0, 0%, 100%, 0.5)",
	"anime-list-item-name-color": "var(--text-color)",

	"post-like-color": "var(--link-color)",
	"post-unlike-color": "var(--link-color)",
	"post-permalink-color": "var(--link-color)",

	"quote-color": "var(--text-color)",

	"calendar-hover-color": "var(--reverse-light-color)"
}

// Toggle theme
export function toggleTheme(arn: AnimeNotifier) {
	if(currentTheme === "light") {
		darkTheme(arn)
		arn.statusMessage.showInfo("Previewing Dark theme. If you would like to use it permanently, please buy a PRO account.", 4000)
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

	// if(arn.user.dataset.pro !== "true") {
	// 	alert("You need a PRO account!")
	// }

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