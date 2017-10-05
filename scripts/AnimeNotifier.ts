import * as actions from "./Actions"
import { displayAiringDate, displayDate } from "./DateView"
import { findAll, delay, canUseWebP, swapElements } from "./Utils"
import { Application } from "./Application"
import { Diff } from "./Diff"
import { MutationQueue } from "./MutationQueue"
import { StatusMessage } from "./StatusMessage"
import { PushManager } from "./PushManager"
import { TouchController } from "./TouchController"
import { Analytics } from "./Analytics"
import { SideBar } from "./SideBar"

export class AnimeNotifier {
	app: Application
	analytics: Analytics
	user: HTMLElement
	title: string
	webpEnabled: boolean
	contentLoadedActions: Promise<any>
	statusMessage: StatusMessage
	visibilityObserver: IntersectionObserver
	pushManager: PushManager
	touchController: TouchController
	sideBar: SideBar
	mainPageLoaded: boolean
	lastReloadContentPath: string

	imageFound: MutationQueue
	imageNotFound: MutationQueue
	unmount: MutationQueue

	constructor(app: Application) {
		this.app = app
		this.user = null
		this.title = "Anime Notifier"

		this.imageFound = new MutationQueue(elem => elem.classList.add("image-found"))
		this.imageNotFound = new MutationQueue(elem => elem.classList.add("image-not-found"))
		this.unmount = new MutationQueue(elem => elem.classList.remove("mounted"))

		// These classes will never be removed on DOM diffs
		Diff.persistentClasses.add("mounted")
		Diff.persistentClasses.add("image-found")

		// Never remove src property on diffs
		Diff.persistentAttributes.add("src")

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
		// App init
		this.app.init()

		// Event listeners
		document.addEventListener("readystatechange", this.onReadyStateChange.bind(this))
		document.addEventListener("DOMContentLoaded", this.onContentLoaded.bind(this))
		document.addEventListener("keydown", this.onKeyDown.bind(this), false)
		window.addEventListener("popstate", this.onPopState.bind(this))

		// Idle
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
		// Check for WebP support
		this.webpEnabled = canUseWebP()

		// Initiate the elements we need
		this.user = this.app.find("user")
		this.app.content = this.app.find("content")
		this.app.loading = this.app.find("loading")

		// Status message
		this.statusMessage = new StatusMessage(
			this.app.find("status-message"),
			this.app.find("status-message-text")
		)

		// Push manager
		this.pushManager = new PushManager()

		// Analytics
		this.analytics = new Analytics()

		// Sidebar control
		this.sideBar = new SideBar(this.app.find("sidebar"))

		// Let"s start
		this.app.run()
	}

	onContentLoaded() {
		// Stop watching all the objects from the previous page.
		this.visibilityObserver.disconnect()
		
		this.contentLoadedActions = Promise.all([
			Promise.resolve().then(() => this.mountMountables()),
			Promise.resolve().then(() => this.lazyLoadImages()),
			Promise.resolve().then(() => this.displayLocalDates()),
			Promise.resolve().then(() => this.setSelectBoxValue()),
			Promise.resolve().then(() => this.assignActions()),
			Promise.resolve().then(() => this.updatePushUI()),
			Promise.resolve().then(() => this.dragAndDrop()),
			Promise.resolve().then(() => this.countUp())
		])

		// Apply page title
		let headers = document.getElementsByTagName("h1")

		if(this.app.currentPath === "/" || headers.length === 0) {
			if(document.title !== this.title) {
				document.title = this.title
			}
		} else {
			document.title = headers[0].innerText
		}
	}

	onIdle() {
		// Service worker
		this.registerServiceWorker()

		// Analytics
		if(this.user) {
			this.analytics.push()
		}

		// Offline message
		if(navigator.onLine === false) {
			this.statusMessage.showError("You are viewing an offline version of the site now.")
		}
	}

	registerServiceWorker() {
		if(!("serviceWorker" in navigator)) {
			return
		}

		console.log("register service worker")

		navigator.serviceWorker.register("/service-worker").then(registration => {
			// registration.update()
		})

		navigator.serviceWorker.addEventListener("message", evt => {
			this.onServiceWorkerMessage(evt)
		})

		// This will send a message to the service worker that the DOM has been loaded
		let sendContentLoadedEvent = () => {
			if(!navigator.serviceWorker.controller) {
				return
			}

			// A reloadContent call should never trigger another reload
			if(this.app.currentPath === this.lastReloadContentPath) {
				console.log("reload finished.")
				this.lastReloadContentPath = ""
				return
			}

			let message = {
				type: "loaded",
				url: ""
			}

			// If mainPageLoaded is set, it means every single request is now an AJAX request for the /_/ prefixed page
			if(this.mainPageLoaded) {
				message.url = window.location.origin + "/_" + window.location.pathname
			} else {
				this.mainPageLoaded = true
				message.url = window.location.href
			}

			console.log("checking for updates:", message.url)

			navigator.serviceWorker.controller.postMessage(JSON.stringify(message))
		}

		// For future loaded events
		document.addEventListener("DOMContentLoaded", sendContentLoadedEvent)

		// If the page is loaded already, send the loaded event right now.
		if(document.readyState !== "loading") {
			sendContentLoadedEvent()
		}
	}

	onServiceWorkerMessage(evt: ServiceWorkerMessageEvent) {
		let message = JSON.parse(evt.data)

		switch(message.type) {
			case "new content":
				if(message.url.includes("/_/")) {
					// Content reload
					this.contentLoadedActions.then(() => {
						this.reloadContent(true)
					})
				} else {
					// Full page reload
					this.contentLoadedActions.then(() => {
						this.reloadPage()
					})
				}
				
				break
		}
	}

	dragAndDrop() {
		for(let element of findAll("inventory-slot")) {
			// Skip elements that have their event listeners attached already
			if(element["listeners-attached"]) {
				continue
			}

			element.addEventListener("dragstart", e => {
				if(!element.draggable) {
					return
				}

				e.dataTransfer.setData("text", element.dataset.index)
			}, false)

			element.addEventListener("dblclick", e => {
				if(!element.draggable) {
					return
				}

				let itemName = element.title

				if(element.dataset.consumable !== "true") {
					return this.statusMessage.showError(itemName + " is not a consumable item.")
				}
				
				let apiEndpoint = this.findAPIEndpoint(element)
				
				this.post(apiEndpoint + "/use/" + element.dataset.index, "")
				.then(() => this.reloadContent())
				.then(() => this.statusMessage.showInfo(`You used ${itemName}.`))
				.catch(err => this.statusMessage.showError(err))
			}, false)

			element.addEventListener("dragenter", e => {
				element.classList.add("drag-enter")
			}, false)

			element.addEventListener("dragleave", e => {
				element.classList.remove("drag-enter")
			}, false)

			element.addEventListener("dragover", e => {
				e.preventDefault()
			}, false)

			element.addEventListener("drop", e => {
				let toElement = e.toElement as HTMLElement
				toElement.classList.remove("drag-enter")

				e.stopPropagation()
				e.preventDefault()

				let inventory = e.toElement.parentElement
				let fromIndex = e.dataTransfer.getData("text")

				if(!fromIndex) {
					return
				}

				let fromElement = inventory.childNodes[fromIndex] as HTMLElement
				
				let toIndex = toElement.dataset.index

				if(fromElement === toElement || fromIndex === toIndex) {
					return
				}

				// Swap in database
				let apiEndpoint = this.findAPIEndpoint(inventory)
				
				this.post(apiEndpoint + "/swap/" + fromIndex + "/" + toIndex, "")
				.catch(err => this.statusMessage.showError(err))

				// Swap in UI
				swapElements(fromElement, toElement)
				
				fromElement.dataset.index = toIndex
				toElement.dataset.index = fromIndex
			}, false)

			// Prevent re-attaching the same listeners
			element["listeners-attached"] = true
		}
	}

	async updatePushUI() {
		if(!this.pushManager.pushSupported || !this.app.currentPath.includes("/settings")) {
			return
		}
		
		let subscription = await this.pushManager.subscription()

		if(subscription) {
			this.app.find("enable-notifications").style.display = "none"
			this.app.find("disable-notifications").style.display = "flex"
		} else {
			this.app.find("enable-notifications").style.display = "flex"
			this.app.find("disable-notifications").style.display = "none"
		}
	}

	countUp() {
		for(let element of findAll("count-up")) {
			let final = parseInt(element.innerText)
			let duration = 2000.0
			let start = Date.now()

			element.innerText = "0"

			let callback = () => {
				let progress = (Date.now() - start) / duration

				if(progress > 1) {
					progress = 1
				}

				element.innerText = String(Math.round(progress * final))

				if(progress < 1) {
					window.requestAnimationFrame(callback)
				}
			}

			window.requestAnimationFrame(callback)
		}
	}

	setSelectBoxValue() {
		for(let element of document.getElementsByTagName("select")) {
			element.value = element.getAttribute("value")
		}
	}

	displayLocalDates() {
		const now = new Date()

		for(let element of findAll("utc-airing-date")) {
			displayAiringDate(element, now)
		}

		for(let element of findAll("utc-date")) {
			displayDate(element, now)
		}
	}

	reloadContent(cached?: boolean) {
		// console.log("reload content", "/_" + this.app.currentPath)

		let headers = new Headers()

		if(!cached) {
			headers.append("X-Reload", "true")
		} else {
			headers.append("X-CacheOnly", "true")
		}

		let path = this.app.currentPath
		this.lastReloadContentPath = path

		return fetch("/_" + path, {
			credentials: "same-origin",
			headers
		})
		.then(response => {
			if(this.app.currentPath !== path) {
				return Promise.reject("old request")
			}

			return Promise.resolve(response)
		})
		.then(response => response.text())
		.then(html => Diff.innerHTML(this.app.content, html))
		.then(() => this.app.emit("DOMContentLoaded"))
	}

	reloadPage() {
		console.log("reload page", this.app.currentPath)

		let path = this.app.currentPath
		this.lastReloadContentPath = path

		return fetch(path, {
			credentials: "same-origin"
		})
		.then(response => {
			if(this.app.currentPath !== path) {
				return Promise.reject("old request")
			}

			return Promise.resolve(response)
		})
		.then(response => response.text())
		.then(html => {
			Diff.root(document.documentElement, html)
		})
		.then(() => this.app.emit("DOMContentLoaded"))
		.then(() => this.loading(false)) // Because our loading element gets reset due to full page diff
	}

	loading(isLoading: boolean) {
		if(isLoading) {
			document.documentElement.style.cursor = "progress"
			this.app.loading.classList.remove(this.app.fadeOutClass)
		} else {
			document.documentElement.style.cursor = "auto"
			this.app.loading.classList.add(this.app.fadeOutClass)
		}
	}
	
	assignActions() {
		for(let element of findAll("action")) {
			let actionTrigger = element.dataset.trigger
			let actionName = element.dataset.action

			let oldAction = element["action assigned"]

			if(oldAction) {
				if(oldAction.trigger === actionTrigger && oldAction.action === actionName) {
					continue
				}

				element.removeEventListener(oldAction.trigger, oldAction.handler)
			}

			if(!(actionName in actions)) {
				this.statusMessage.showError(`Action '${actionName}' has not been defined`)
			}

			let actionHandler = e => {
				actions[actionName](this, element, e)

				e.stopPropagation()
				e.preventDefault()
			}

			element.addEventListener(actionTrigger, actionHandler)

			// Use "action assigned" flag instead of removing the class.
			// This will make sure that DOM diffs which restore the class name
			// will not assign the action multiple times to the same element.
			element["action assigned"] = {
				trigger: actionTrigger,
				action: actionName,
				handler: actionHandler
			}
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
			// Replace URL with WebP if supported
			if(this.webpEnabled && img.dataset.webp) {
				let dot = img.dataset.src.lastIndexOf(".")
				img.src = img.dataset.src.substring(0, dot) + ".webp"
			} else {
				img.src = img.dataset.src
			}

			if(img.naturalWidth === 0) {
				img.onload = () => {
					this.imageFound.queue(img)
				}

				img.onerror = () => {
					this.imageNotFound.queue(img)
				}
			} else {
				this.imageFound.queue(img)
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

			this.unmount.queue(element)
		}
	}

	modifyDelayed(className: string, func: (element: HTMLElement) => void) {
		const maxDelay = 1000
		const delay = 20
		
		let time = 0
		let start = Date.now()
		let maxTime = start + maxDelay

		let mountableTypes = new Map<string, number>()
		let mountableTypeMutations = new Map<string, Array<any>>()

		let collection = document.getElementsByClassName(className)

		if(collection.length === 0) {
			return
		}

		// let delay = Math.min(maxDelay / collection.length, 20)

		for(let i = 0; i < collection.length; i++) {
			let element = collection.item(i) as HTMLElement
			let type = element.dataset.mountableType || "general"

			if(mountableTypes.has(type)) {
				time = mountableTypes.get(type) + delay
				mountableTypes.set(type, time)
			} else {
				time = start
				mountableTypes.set(type, time)
				mountableTypeMutations.set(type, [])
			}

			if(time > maxTime) {
				time = maxTime
			}

			mountableTypeMutations.get(type).push({
				element,
				time
			})
		}

		for(let mountableType of mountableTypeMutations.keys()) {
			let mutations = mountableTypeMutations.get(mountableType)
			let mutationIndex = 0

			let updateBatch = () => {
				let now = Date.now()

				for(; mutationIndex < mutations.length; mutationIndex++) {
					let mutation = mutations[mutationIndex]

					if(mutation.time > now) {
						break
					}

					func(mutation.element)
				}

				if(mutationIndex < mutations.length) {
					window.requestAnimationFrame(updateBatch)
				}
			}

			window.requestAnimationFrame(updateBatch)
		}
	}

	diff(url: string) {
		if(url === this.app.currentPath) {
			return Promise.reject(null)
		}

		let path = "/_" + url

		let request = fetch(path, {
			credentials: "same-origin"
		})
		.then(response => {
			return response.text()
		})
		
		history.pushState(url, null, url)
		this.app.currentPath = url
		this.app.markActiveLinks()
		this.unmountMountables()
		this.loading(true)

		// Delay by transition-speed
		return delay(300).then(() => request)
		.then(html => this.app.setContent(html, true))
		.then(() => this.app.emit("DOMContentLoaded"))
		.then(() => this.loading(false))
		.catch(console.error)
	}

	post(url, body) {
		if(typeof body !== "string") {
			body = JSON.stringify(body)
		}

		this.loading(true)

		return fetch(url, {
			method: "POST",
			body,
			credentials: "same-origin"
		})
		.then(response => response.text())
		.then(body => {
			this.loading(false)

			if(body !== "ok") {
				throw body
			}
		})
		.catch(err => {
			this.loading(false)
			throw err
		})
	}

	scrollTo(target: HTMLElement) {
		const duration = 250.0
		const steps = 60
		const interval = duration / steps
		const fullSin = Math.PI / 2
		const contentPadding = 24
		
		let scrollHandle: number
		let oldScroll = this.app.content.parentElement.scrollTop
		let newScroll = 0
		let finalScroll = Math.max(target.offsetTop - contentPadding, 0)
		let scrollDistance = finalScroll - oldScroll

		if(scrollDistance > 0 && scrollDistance < 4) {
			return
		}

		let timeStart = Date.now()
		let timeEnd = timeStart + duration

		let scroll = () => {
			let time = Date.now()
			let progress = (time - timeStart) / duration

			if(progress > 1.0) {
				progress = 1.0
			}

			newScroll = oldScroll + scrollDistance * Math.sin(progress * fullSin)
			this.app.content.parentElement.scrollTop = newScroll

			if(time < timeEnd && newScroll != finalScroll) {
				window.requestAnimationFrame(scroll)
			}
		}

		window.requestAnimationFrame(scroll)
	}

	findAPIEndpoint(element: HTMLElement) {
		if(element.dataset.api !== undefined) {
			return element.dataset.api
		}

		let apiObject: HTMLElement
		let parent = element

		while(parent = parent.parentElement) {
			if(parent.dataset.api !== undefined) {
				apiObject = parent
				break
			}
		}

		if(!apiObject) {
			throw "API object not found"
		}

		return apiObject.dataset.api
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