import { Application } from "./Application"
import { Diff } from "./Diff"
import { displayLocalDate } from "./DateView"
import { findAll, delay } from "./Utils"
import * as actions from "./Actions"

export class AnimeNotifier {
	app: Application
	visibilityObserver: IntersectionObserver
	user: HTMLElement

	constructor(app: Application) {
		this.app = app
		this.user = null

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

	init() {
		document.addEventListener("readystatechange", this.onReadyStateChange.bind(this))
		document.addEventListener("DOMContentLoaded", this.onContentLoaded.bind(this))
		document.addEventListener("keydown", this.onKeyDown.bind(this), false)
		window.addEventListener("popstate", this.onPopState.bind(this))

		this.requestIdleCallback(this.onIdle.bind(this))
	}

	requestIdleCallback(func: Function) {
		if("requestIdleCallback" in window) {
			window["requestIdleCallback"](func)
		} else {
			func()
		}
	}

	onReadyStateChange() {
		if(document.readyState !== "interactive") {
			return
		}

		this.run()
	}

	run() {
		// Add "osx" class on macs so we can set a proper font-size
		if(navigator.platform.includes("Mac")) {
			document.documentElement.classList.add("osx")
		}

		// Initiate the elements we need
		this.user = this.app.find("user")
		this.app.content = this.app.find("content")
		this.app.loading = this.app.find("loading")

		// Let's start
		this.app.run()
	}

	onContentLoaded() {
		// Stop watching all the objects from the previous page.
		this.visibilityObserver.disconnect()
		
		Promise.resolve().then(() => this.mountMountables()),
		Promise.resolve().then(() => this.lazyLoadImages()),
		Promise.resolve().then(() => this.displayLocalDates()),
		Promise.resolve().then(() => this.setSelectBoxValue()),
		Promise.resolve().then(() => this.assignActions())
	}

	onIdle() {
		if(!this.user) {
			return
		}

		let analytics = {
			general: {
				timezoneOffset: new Date().getTimezoneOffset()
			},
			screen: {
				width: screen.width,
				height: screen.height,
				availableWidth: screen.availWidth,
				availableHeight: screen.availHeight,
				pixelRatio: window.devicePixelRatio
			},
			system: {
				cpuCount: navigator.hardwareConcurrency,
				platform: navigator.platform
			}
		}

		fetch("/api/analytics/new", {
			method: "POST",
			credentials: "same-origin",
			body: JSON.stringify(analytics)
		})
	}

	setSelectBoxValue() {
		for(let element of document.getElementsByTagName("select")) {
			element.value = element.getAttribute("value")
		}
	}

	displayLocalDates() {
		const now = new Date()

		for(let element of findAll("utc-date")) {
			displayLocalDate(element, now)
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
			document.body.style.cursor = "progress"
			this.app.loading.classList.remove(this.app.fadeOutClass)
		} else {
			document.body.style.cursor = "auto"
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
			if(element.classList.contains("never-unmount")) {
				continue
			}

			element.classList.remove("mounted")
		}
	}

	modifyDelayed(className: string, func: (element: HTMLElement) => void) {
		let mountableTypes = {
			general: 0
		}

		const delay = 20
		const maxDelay = 1000
		
		let time = 0
		let collection = document.getElementsByClassName(className)

		for(let i = 0; i < collection.length; i++) {
			let element = collection.item(i) as HTMLElement
			let type = element.dataset.mountableType || "general"

			if(type in mountableTypes) {
				time = mountableTypes[type] += delay
			} else {
				time = mountableTypes[type] = 0
			}

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
		if(url == this.app.currentPath) {
			return
		}

		let request = fetch("/_" + url, {
			credentials: "same-origin"
		})
		.then(response => response.text())
		
		history.pushState(url, null, url)
		this.app.currentPath = url
		this.app.markActiveLinks()
		this.unmountMountables()
		this.loading(true)

		// Delay by transition-speed
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
		let activeElement = document.activeElement

		// Ignore hotkeys on input elements
		switch(activeElement.tagName) {
			case "INPUT":
			case "TEXTAREA":
				return
		}

		// Disallow Enter key in contenteditables
		if(activeElement.getAttribute("contenteditable") === "true" && e.keyCode == 13) {
			if("blur" in activeElement) {
				activeElement["blur"]()
			}
			
			e.preventDefault()
			e.stopPropagation()
			return
		}

		// F = Search
		if(e.keyCode == 70 && !e.ctrlKey) {
			let search = this.app.find("search") as HTMLInputElement

			search.focus()
			search.select()

			e.preventDefault()
			e.stopPropagation()
			return
		}
	}
}