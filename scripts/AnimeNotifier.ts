import { Application } from "./Application"
import { Diff } from "./Diff"
import { findAll, delay } from "./utils"
import * as actions from "./actions"

var monthNames = [
	"January", "February", "March",
	"April", "May", "June", "July",
	"August", "September", "October",
	"November", "December"
]

var dayNames = [
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday"
]

export class AnimeNotifier {
	app: Application
	visibilityObserver: IntersectionObserver

	constructor(app: Application) {
		this.app = app

		if("IntersectionObserver" in window) {
			// Enable lazy load
			this.visibilityObserver = new IntersectionObserver(
				entries => {
					for(let entry of entries) {
						if(entry.intersectionRatio > 0) {
							entry.target["became visible"]()
							this.visibilityObserver.unobserve(entry.target)
						}
					}
				},
				{}
			)
		} else {
			// Disable lazy load feature
			this.visibilityObserver = {
				disconnect: () => {},
				observe: (elem: HTMLElement) => {
					elem["became visible"]()
				},
				unobserve: (elem: HTMLElement) => {}
			} as IntersectionObserver
		}
	}

	onReadyStateChange() {
		if(document.readyState !== "interactive") {
			return
		}

		this.run()
	}

	run() {
		this.app.content = this.app.find("content")
		this.app.loading = this.app.find("loading")
		this.app.run()
	}

	onContentLoaded() {
		this.visibilityObserver.disconnect()
		
		// Update each of these asynchronously
		Promise.resolve().then(() => this.mountMountables())
		Promise.resolve().then(() => this.assignActions())
		Promise.resolve().then(() => this.lazyLoadImages())
		Promise.resolve().then(() => this.displayLocalDates())
	}

	displayLocalDates() {
		const oneDay = 24 * 60 * 60 * 1000
		const now = new Date()

		for(let element of findAll("utc-date")) {
			let startDate = new Date(element.dataset.startDate)
			let endDate = new Date(element.dataset.endDate)

			let h = startDate.getHours()
			let m = startDate.getMinutes()
			let startTime = (h <= 9 ? "0" + h : h) + ":" + (m <= 9 ? "0" + m : m)

			h = endDate.getHours()
			m = endDate.getMinutes()
			let endTime = (h <= 9 ? "0" + h : h) + ":" + (m <= 9 ? "0" + m : m)
			
			let dayDifference = Math.round(Math.abs((startDate.getTime() - now.getTime()) / oneDay))
			let dayInfo = dayNames[startDate.getDay()] + ", " + monthNames[startDate.getMonth()] + " " + startDate.getDate()

			if(dayDifference > 1) {
				element.innerText = dayDifference + " day" + (dayDifference == 1 ? "" : "s")
			} else if(dayDifference == 1) {
				element.innerText = "Tomorrow"
			} else {
				element.innerText = "Today"
			}

			element.title = "Episode " + element.dataset.episodeNumber + " will be airing " + startTime + " - " + endTime + " your time"
		}
	}

	reloadContent() {
		return fetch("/_" + this.app.currentPath, {
			credentials: "same-origin"
		})
		.then(response => response.text())
		.then(html => Diff.innerHTML(this.app.content, html))
		.then(() => this.app.emit("DOMContentLoaded"))
	}

	loading(isLoading: boolean) {
		if(isLoading) {
			this.app.loading.classList.remove(this.app.fadeOutClass)
		} else {
			this.app.loading.classList.add(this.app.fadeOutClass)
		}
	}
	
	assignActions() {
		for(let element of findAll("action")) {
			if(element["action assigned"]) {
				continue
			}

			let actionName = element.dataset.action

			element.addEventListener(element.dataset.trigger, e => {
				actions[actionName](this, element, e)

				e.stopPropagation()
				e.preventDefault()
			})

			// Use "action assigned" flag instead of removing the class.
			// This will make sure that DOM diffs which restore the class name
			// will not assign the action multiple times to the same element.
			element["action assigned"] = true
		}
	}

	lazyLoadImages() {
		for(let element of findAll("lazy")) {
			this.lazyLoadImage(element as HTMLImageElement)
		}
	}

	lazyLoadImage(img: HTMLImageElement) {
		// Once the image becomes visible, load it
		img["became visible"] = () => {
			img.src = img.dataset.src

			if(img.naturalWidth === 0) {
				img.onload = function() {
					this.classList.add("image-found")
				}

				img.onerror = function() {
					this.classList.add("image-not-found")
				}
			} else {
				img.classList.add("image-found")
			}
		}

		this.visibilityObserver.observe(img)
	}

	mountMountables() {
		this.modifyDelayed("mountable", element => element.classList.add("mounted"))
	}

	unmountMountables() {
		for(let element of findAll("mountable")) {
			element.classList.remove("mounted")
		}
	}

	modifyDelayed(className: string, func: (element: HTMLElement) => void) {
		const delay = 20
		const maxDelay = 500
		
		let time = 0

		for(let element of findAll(className)) {
			time += delay

			if(time > maxDelay) {
				func(element)
			} else {
				setTimeout(() => {
					window.requestAnimationFrame(() => func(element))
				}, time)
			}
		}
	}

	diff(url: string) {
		let request = fetch("/_" + url, {
			credentials: "same-origin"
		})
		.then(response => response.text())
		
		history.pushState(url, null, url)
		this.app.currentPath = url
		this.app.markActiveLinks()
		this.unmountMountables()
		this.loading(true)

		return delay(300).then(() => {
			request
			.then(html => this.app.setContent(html, true))
			.then(() => this.app.markActiveLinks())
			.then(() => this.app.emit("DOMContentLoaded"))
			.then(() => this.loading(false))
			.catch(console.error)
		})
	}

	post(url, obj) {
		return fetch(url, {
			method: "POST",
			body: JSON.stringify(obj),
			credentials: "same-origin"
		})
		.then(response => response.text())
		.then(body => {
			if(body !== "ok") {
				throw body
			}
		})
	}

	onPopState(e: PopStateEvent) {
		if(e.state) {
			this.app.load(e.state, {
				addToHistory: false
			})
		} else if(this.app.currentPath !== this.app.originalPath) {
			this.app.load(this.app.originalPath, {
				addToHistory: false
			})
		}
	}

	onKeyDown(e: KeyboardEvent) {
		// Ignore hotkeys on input elements
		switch(document.activeElement.tagName) {
			case "INPUT":
			case "TEXTAREA":
				return
		}

		// F = Search
		if(e.keyCode == 70) {
			let search = this.app.find("search") as HTMLInputElement

			search.focus()
			search.select()

			e.preventDefault()
			e.stopPropagation()
		}
	}

	// onResize(e: UIEvent) {
	// 	let hasScrollbar = this.app.content.clientHeight === this.app.content.scrollHeight

	// 	if(hasScrollbar) {
	// 		this.app.content.classList.add("has-scrollbar")
	// 	} else {
	// 		this.app.content.classList.remove("has-scrollbar")
	// 	}
	// }
}