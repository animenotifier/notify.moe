import Application from "./Application"
import Diff from "./Diff"
import StatusMessage from "./StatusMessage"
import PushManager from "./PushManager"
import TouchController from "./TouchController"
import NotificationManager from "./NotificationManager"
import AudioPlayer from "./AudioPlayer"
import Analytics from "./Analytics"
import SideBar from "./SideBar"
import InfiniteScroller from "./InfiniteScroller"
import ServiceWorkerManager from "./ServiceWorkerManager"
import ServerEvents from "./ServerEvents"
import { displayAiringDate, displayDate, displayTime } from "./DateView"
import { findAll, canUseWebP, requestIdleCallback, swapElements, delay, findAllInside } from "./Utils"
import * as actions from "./Actions"

export default class AnimeNotifier {
	app: Application
	analytics: Analytics
	user: HTMLElement
	title: string
	webpEnabled: boolean
	contentLoadedActions: Promise<any>
	statusMessage: StatusMessage
	visibilityObserver: IntersectionObserver
	pushManager: PushManager
	serviceWorkerManager: ServiceWorkerManager
	notificationManager: NotificationManager
	touchController: TouchController
	audioPlayer: AudioPlayer
	sideBar: SideBar
	infiniteScroller: InfiniteScroller
	mainPageLoaded: boolean
	isLoading: boolean
	diffCompletedForCurrentPath: boolean
	lastReloadContentPath: string
	currentSoundTrackId: string
	serverEvents: ServerEvents

	constructor(app: Application) {
		this.app = app
		this.user = null
		this.title = "Anime Notifier"
		this.isLoading = true

		// These classes will never be removed on DOM diffs
		Diff.persistentClasses.add("mounted")
		Diff.persistentClasses.add("element-found")
		Diff.persistentClasses.add("active")

		// Never remove src property on diffs
		Diff.persistentAttributes.add("src")

		// Is intersection observer supported?
		if("IntersectionObserver" in window) {
			// Enable lazy load
			this.visibilityObserver = new IntersectionObserver(
				entries => {
					for(let entry of entries) {
						if(entry.isIntersecting) {
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

		// If we finished loading the DOM (either "interactive" or "complete" state),
		// immediately trigger the event listener functions.
		if(document.readyState !== "loading") {
			this.app.emit("DOMContentLoaded")
			this.run()
		}

		// Idle
		requestIdleCallback(this.onIdle.bind(this))
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
		this.user = document.getElementById("user")
		this.app.content = document.getElementById("content")
		this.app.loading = document.getElementById("loading")

		// Theme
		if(this.user && this.user.dataset.pro === "true" && this.user.dataset.theme !== "light") {
			actions.applyTheme(this.user.dataset.theme)
		}

		// Status message
		this.statusMessage = new StatusMessage(
			document.getElementById("status-message"),
			document.getElementById("status-message-text")
		)

		this.app.onError = (error: Error) => {
			this.statusMessage.showError(error, 3000)
		}

		// Push manager
		this.pushManager = new PushManager()

		// Notification manager
		this.notificationManager = new NotificationManager()

		// Audio player
		this.audioPlayer = new AudioPlayer(this)

		// Analytics
		this.analytics = new Analytics()

		// Sidebar control
		this.sideBar = new SideBar(document.getElementById("sidebar"))

		// Infinite scrolling
		this.infiniteScroller = new InfiniteScroller(this.app.content.parentElement, 150)

		// Loading
		this.loading(false)
	}

	onContentLoaded() {
		// Stop watching all the objects from the previous page.
		this.visibilityObserver.disconnect()

		this.contentLoadedActions = Promise.all([
			Promise.resolve().then(() => this.mountMountables()),
			Promise.resolve().then(() => this.lazyLoad()),
			Promise.resolve().then(() => this.displayLocalDates()),
			Promise.resolve().then(() => this.setSelectBoxValue()),
			Promise.resolve().then(() => this.textAreaFocus()),
			Promise.resolve().then(() => this.markPlayingSoundTrack()),
			Promise.resolve().then(() => this.assignActions()),
			Promise.resolve().then(() => this.updatePushUI()),
			Promise.resolve().then(() => this.dragAndDrop()),
			Promise.resolve().then(() => this.colorStripes()),
			Promise.resolve().then(() => this.loadCharacterRanking()),
			Promise.resolve().then(() => this.assignTooltipOffsets()),
			Promise.resolve().then(() => this.countUp())
		])

		// Apply page title
		this.applyPageTitle()

		// Auto-focus first input element on welcome page.
		if(location.pathname === "/welcome") {
			let firstInput = this.app.content.getElementsByTagName("input")[0] as HTMLInputElement

			if(firstInput) {
				firstInput.focus()
			}
		}
	}

	applyPageTitle() {
		let headers = document.getElementsByTagName("h1")

		if(this.app.currentPath === "/" || headers.length === 0 || headers[0].textContent === "NOTIFY.MOE") {
			if(document.title !== this.title) {
				document.title = this.title
			}
		} else {
			document.title = headers[0].textContent
		}
	}

	textAreaFocus() {
		const newPostText = document.getElementById("new-post-text") as HTMLTextAreaElement

		if(!newPostText || newPostText["has-input-listener"]) {
			return
		}

		newPostText.addEventListener("input", () => {
			if(newPostText.value.length > 0) {
				const newPostActions = document.getElementsByClassName("new-post-actions")[0]
				newPostActions.classList.add("new-post-actions-enabled")
			} else {
				const newPostActions = document.getElementsByClassName("new-post-actions")[0]
				newPostActions.classList.remove("new-post-actions-enabled")
			}
		})

		newPostText["has-input-listener"] = true
	}

	async onIdle() {
		// Register event listeners
		document.addEventListener("keydown", this.onKeyDown.bind(this), false)
		window.addEventListener("popstate", this.onPopState.bind(this))
		window.addEventListener("error", this.onError.bind(this))

		// Service worker
		this.serviceWorkerManager = new ServiceWorkerManager(this, "/service-worker")
		this.serviceWorkerManager.register()

		// Analytics
		if(this.user) {
			this.analytics.push()
		}

		// Offline message
		if(navigator.onLine === false) {
			this.statusMessage.showInfo("You are viewing an offline version of the site now.")
		}

		// Notification manager
		if(this.user) {
			this.notificationManager.update()
		}

		// Bind unload event
		window.addEventListener("beforeunload", this.onBeforeUnload.bind(this))

		// Show microphone icon if speech input is available
		if(window["SpeechRecognition"] || window["webkitSpeechRecognition"]) {
			document.getElementsByClassName("speech-input")[0].classList.add("speech-input-available")
		}

		// Ensure a minimum size for the desktop app
		const minWidth = 1420
		const minHeight = 800

		if(window.outerWidth <= minWidth || window.outerHeight <= minHeight) {
			let finalWidth = window.outerWidth < minWidth ? minWidth : window.outerWidth
			let finalHeight = window.outerHeight < minHeight ? minHeight : window.outerHeight

			window.resizeTo(finalWidth, finalHeight)
		}

		// Server sent events
		if(this.user && EventSource) {
			this.serverEvents = new ServerEvents(this)
		}

		// // Download popular anime titles for the search
		// let response = await fetch("/api/popular/anime/titles/500")
		// let titles = await response.json()
		// let titleList = document.createElement("datalist")
		// titleList.id = "popular-anime-titles-list"

		// for(let title of titles) {
		// 	let option = document.createElement("option")
		// 	option.value = title
		// 	titleList.appendChild(option)
		// }

		// document.body.appendChild(titleList)

		// let search = document.getElementById("search") as HTMLInputElement
		// search.setAttribute("list", titleList.id)
	}

	async onBeforeUnload(e: BeforeUnloadEvent) {
		let message = undefined

		// Prevent closing tab on new thread page
		if(this.app.currentPath === "/new/thread" && document.activeElement.tagName === "TEXTAREA" && (document.activeElement as HTMLTextAreaElement).value.length > 20) {
			message = "You have unsaved changes on the current page. Are you sure you want to leave?"
		}

		if(message) {
			e.returnValue = message
			return message
		}
	}

	assignTooltipOffsets(elements?: IterableIterator<HTMLElement>) {
		requestIdleCallback(() => {
			const distanceToBorder = 5
			let contentRect: ClientRect

			if(!elements) {
				elements = findAll("tip")
			}

			// Assign mouse enter event handler
			for(let element of elements) {
				Diff.mutations.queue(() => {
					element.classList.add("tip-active")
				})

				element.onmouseenter = () => {
					Diff.mutations.queue(() => {
						if(!contentRect) {
							contentRect = this.app.content.getBoundingClientRect()
						}

						// Dynamic label assignment to prevent label texts overflowing
						// and taking horizontal space at page load.
						element.dataset.label = element.getAttribute("aria-label")

						// This is the most expensive call in this whole function,
						// it consumes about 2-4 ms every time you call it.
						let rect = element.getBoundingClientRect()

						// Calculate offsets
						let tipStyle = window.getComputedStyle(element, ":before")
						let tipWidth = parseInt(tipStyle.width) + parseInt(tipStyle.paddingLeft) * 2
						let tipStartX = rect.left + rect.width / 2 - tipWidth / 2 - contentRect.left
						let tipEndX = tipStartX + tipWidth
						let leftOffset = 0

						if(tipStartX < distanceToBorder) {
							leftOffset = -tipStartX + distanceToBorder
						} else if(tipEndX > contentRect.width - distanceToBorder) {
							leftOffset = -(tipEndX - contentRect.width + distanceToBorder)
						}

						if(leftOffset !== 0) {
							element.classList.remove("tip-active")
							element.classList.add("tip-offset-root")

							let tipChild = document.createElement("div")
							tipChild.classList.add("tip-offset-child")
							tipChild.setAttribute("data-label", element.getAttribute("data-label"))
							tipChild.style.left = Math.round(leftOffset) + "px"
							tipChild.style.width = rect.width + "px"
							tipChild.style.height = rect.height + "px"
							element.appendChild(tipChild)
						}

						// Unassign event listener
						element.onmouseenter = null
					})
				}
			}
		})
	}

	dragAndDrop() {
		if(location.pathname.includes("/animelist/")) {
			for(let element of findAll("anime-list-item")) {
				// Skip elements that have their event listeners attached already
				if(element["drag-listeners-attached"]) {
					continue
				}

				element.addEventListener("dragstart", e => {
					if(!element.draggable) {
						return
					}

					let image = element.getElementsByClassName("anime-list-item-image")[0]
					e.dataTransfer.setDragImage(image, 0, 0)

					let name = element.getElementsByClassName("anime-list-item-name")[0]

					e.dataTransfer.setData("text/plain", JSON.stringify({
						api: element.dataset.api,
						animeTitle: name.textContent
					}))
					e.dataTransfer.effectAllowed = "move"
				}, false)

				// Prevent re-attaching the same listeners
				element["drag-listeners-attached"] = true
			}

			for(let element of findAll("tab")) {
				// Skip elements that have their event listeners attached already
				if(element["drop-listeners-attached"]) {
					continue
				}

				element.addEventListener("drop", async e => {
					let toElement = e.toElement as HTMLElement

					// Find tab element
					while(toElement && !toElement.classList.contains("tab")) {
						toElement = toElement.parentElement
					}

					// Ignore a drop on the current status tab
					if(!toElement || toElement.classList.contains("active")) {
						return
					}

					let data = e.dataTransfer.getData("text/plain")
					let json = null

					try {
						json = JSON.parse(data)
					} catch {
						return
					}

					if(!json || !json.api) {
						return
					}

					e.stopPropagation()
					e.preventDefault()

					let tabText = toElement.textContent
					let newStatus = tabText.toLowerCase()

					if(newStatus === "on hold") {
						newStatus = "hold"
					}

					try {
						await this.post(json.api, {
							Status: newStatus
						})
						await this.reloadContent()

						this.statusMessage.showInfo(`Moved "${json.animeTitle}" to "${tabText}".`)
					} catch(err) {
						this.statusMessage.showError(err)
					}

				}, false)

				element.addEventListener("dragenter", e => {
					e.preventDefault()
				}, false)

				element.addEventListener("dragleave", e => {
					e.preventDefault()
				}, false)

				element.addEventListener("dragover", e => {
					e.preventDefault()
				}, false)

				// Prevent re-attaching the same listeners
				element["drop-listeners-attached"] = true
			}
		}

		if(location.pathname.startsWith("/inventory")) {
			for(let element of findAll("inventory-slot")) {
				// Skip elements that have their event listeners attached already
				if(element["drag-listeners-attached"]) {
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

					let itemName = element.getAttribute("aria-label")

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
				element["drag-listeners-attached"] = true
			}
		}
	}

	async updatePushUI() {
		if(!this.app.currentPath.includes("/settings/notifications")) {
			return
		}

		let enableButton = document.getElementById("enable-notifications") as HTMLButtonElement
		let disableButton = document.getElementById("disable-notifications") as HTMLButtonElement
		let testButton = document.getElementById("test-notification") as HTMLButtonElement

		if(!this.pushManager.pushSupported) {
			enableButton.classList.add("hidden")
			disableButton.classList.add("hidden")
			testButton.innerHTML = "Your browser doesn't support push notifications!"
			return
		}

		let subscription = await this.pushManager.subscription()

		if(subscription) {
			enableButton.classList.add("hidden")
			disableButton.classList.remove("hidden")
		} else {
			enableButton.classList.remove("hidden")
			disableButton.classList.add("hidden")
		}
	}

	loadCharacterRanking() {
		if(!this.app.currentPath.includes("/character/")) {
			return
		}

		for(let element of findAll("character-ranking")) {
			fetch(`/api/character/${element.dataset.characterId}/ranking`).then(async response => {
				let ranking = await response.json()

				if(!ranking.rank) {
					return
				}

				Diff.mutations.queue(() => {
					let percentile = Math.ceil(ranking.percentile * 100)

					element.textContent = "#" + ranking.rank.toString()
					element.title = "Top " + percentile + "%"
				})
			})
		}
	}

	colorStripes() {
		if(!this.app.currentPath.includes("/explore/color/")) {
			return
		}

		for(let element of findAll("color-stripe")) {
			Diff.mutations.queue(() => {
				element.style.backgroundColor = element.dataset.color
			})
		}
	}

	countUp() {
		if(!this.app.currentPath.includes("/paypal/success")) {
			return
		}

		for(let element of findAll("count-up")) {
			let final = parseInt(element.textContent)
			let duration = 2000.0
			let start = Date.now()

			element.textContent = "0"

			let callback = () => {
				let progress = (Date.now() - start) / duration

				if(progress > 1) {
					progress = 1
				}

				element.textContent = String(Math.round(progress * final))

				if(progress < 1) {
					window.requestAnimationFrame(callback)
				}
			}

			window.requestAnimationFrame(callback)
		}
	}

	markPlayingSoundTrack() {
		for(let element of findAll("soundtrack-play-area")) {
			if(element.dataset.soundtrackId === this.currentSoundTrackId) {
				element.classList.add("playing")
			}
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

		for(let element of findAll("utc-date-absolute")) {
			displayTime(element, now)
		}
	}

	reloadContent(cached?: boolean) {
		let headers = new Headers()

		if(cached) {
			headers.set("X-Force-Cache", "true")
		} else {
			headers.set("X-No-Cache", "true")
		}

		let path = this.lastReloadContentPath = this.app.currentPath

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
		.then(html => Diff.root(document.documentElement, html))
		.then(() => this.app.emit("DOMContentLoaded"))
		.then(() => this.loading(false)) // Because our loading element gets reset due to full page diff
	}

	loading(newState: boolean) {
		this.isLoading = newState

		if(this.isLoading) {
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

			// Filter out invalid definitions
			if(!actionTrigger || !actionName) {
				continue
			}

			let oldAction = element["action assigned"]

			if(oldAction) {
				if(oldAction.trigger === actionTrigger && oldAction.action === actionName) {
					continue
				}

				element.removeEventListener(oldAction.trigger, oldAction.handler)
			}

			// This prevents default actions on links
			if(actionTrigger === "click" && element.tagName === "A") {
				element.onclick = null
			}

			// Warn us about undefined actions
			if(!(actionName in actions)) {
				this.statusMessage.showError(`Action '${actionName}' has not been defined`)
				continue
			}

			// Register the actual action handler
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

	lazyLoad(elements?: IterableIterator<Element>) {
		if(!elements) {
			elements = findAll("lazy")
		}

		for(let element of elements) {
			switch(element.tagName) {
				case "IMG":
					this.lazyLoadImage(element as HTMLImageElement)
					break

				case "VIDEO":
					this.lazyLoadVideo(element as HTMLVideoElement)
					break

				case "IFRAME":
					this.lazyLoadIFrame(element as HTMLIFrameElement)
					break
			}
		}
	}

	emptyPixel() {
		return "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII="
	}

	lazyLoadImage(element: HTMLImageElement) {
		let pixelRatio = window.devicePixelRatio

		// Once the image becomes visible, load it
		element["became visible"] = () => {
			let dataSrc = element.dataset.src
			let dotPos = dataSrc.lastIndexOf(".")
			let base = dataSrc.substring(0, dotPos)
			let extension = ""

			// Replace URL with WebP if supported
			if(this.webpEnabled && element.dataset.webp === "true") {
				let queryPos = dataSrc.lastIndexOf("?")

				if(queryPos !== -1) {
					extension = ".webp" + dataSrc.substring(queryPos)
				} else {
					extension = ".webp"
				}
			} else {
				extension = dataSrc.substring(dotPos)
			}

			// Anime and character images on Retina displays
			if(pixelRatio > 1) {
				if(base.includes("/anime/") || (base.includes("/characters/") && !base.includes("/large/"))) {
					base += "@2"
				}
			}

			let finalSrc = base + extension

			if(element.src !== finalSrc && element.src !== "https:" + finalSrc && element.src !== "https://notify.moe" + finalSrc) {
				// Show average color
				if(element.dataset.color) {
					element.src = this.emptyPixel()
					element.style.backgroundColor = element.dataset.color
					Diff.mutations.queue(() => element.classList.add("element-color-preview"))
				}

				Diff.mutations.queue(() => element.classList.remove("element-found"))
				element.src = finalSrc
			}

			if(element.naturalWidth === 0) {
				element.onload = () => {
					if(element.src.startsWith("data:")) {
						return
					}

					Diff.mutations.queue(() => element.classList.add("element-found"))
				}

				element.onerror = () => {
					if(element.classList.contains("element-found")) {
						return
					}

					Diff.mutations.queue(() => element.classList.add("element-not-found"))
				}
			} else {
				Diff.mutations.queue(() => element.classList.add("element-found"))
			}
		}

		this.visibilityObserver.observe(element)
	}

	lazyLoadIFrame(element: HTMLIFrameElement) {
		// Once the iframe becomes visible, load it
		element["became visible"] = () => {
			// If the source is already set correctly, don't set it again to avoid iframe flickering.
			if(element.src !== element.dataset.src && element.src !== (window.location.protocol + element.dataset.src)) {
				element.src = element.dataset.src
			}

			Diff.mutations.queue(() => element.classList.add("element-found"))
		}

		this.visibilityObserver.observe(element)
	}

	lazyLoadVideo(video: HTMLVideoElement) {
		// Once the video becomes visible, load it
		video["became visible"] = () => {
			video.pause()

			// Prevent context menu
			video.addEventListener("contextmenu", e => e.preventDefault())

			let modified = false

			for(let child of video.children) {
				let element = child as HTMLSourceElement

				if(element.src !== element.dataset.src) {
					element.src = element.dataset.src
					modified = true
				}
			}

			if(modified) {
				Diff.mutations.queue(() => {
					video.load()
					video.classList.add("element-found")
				})
			}
		}

		this.visibilityObserver.observe(video)
	}

	mountMountables(elements?: IterableIterator<HTMLElement>) {
		if(!elements) {
			elements = findAll("mountable")
		}

		this.modifyDelayed(elements, element => element.classList.add("mounted"))
	}

	unmountMountables() {
		for(let element of findAll("mountable")) {
			if(element.classList.contains("never-unmount")) {
				continue
			}

			Diff.mutations.queue(() => element.classList.remove("mounted"))
		}
	}

	modifyDelayed(elements: IterableIterator<HTMLElement>, func: (element: HTMLElement) => void) {
		const maxDelay = 2500
		const delay = 20

		let time = 0
		let start = Date.now()
		let maxTime = start + maxDelay

		let mountableTypes = new Map<string, number>()
		let mountableTypeMutations = new Map<string, Array<any>>()

		for(let element of elements) {
			// Skip already mounted elements.
			// This helps a lot when dealing with infinite scrolling
			// where the first elements are already mounted.
			if(element.classList.contains("mounted")) {
				continue
			}

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

	async diff(url: string) {
		if(url === this.app.currentPath) {
			return null
		}

		let path = "/_" + url

		try {
			// Start the request
			let request = fetch(path, {
				credentials: "same-origin"
			})
			.then(response => response.text())

			history.pushState(url, null, url)
			this.app.currentPath = url
			this.diffCompletedForCurrentPath = false
			this.app.markActiveLinks()
			this.unmountMountables()
			this.loading(true)

			// Delay by mountable-transition-speed
			await delay(150)

			let html = await request

			// If the response for the correct path has not arrived yet, show this response
			if(!this.diffCompletedForCurrentPath) {
				// If this response was the most recently requested one, mark the requests as completed
				if(this.app.currentPath === url) {
					this.diffCompletedForCurrentPath = true
				}

				// Update contents
				await Diff.innerHTML(this.app.content, html)
				this.app.emit("DOMContentLoaded")
			}
		} catch(err) {
			console.error(err)
		} finally {
			this.loading(false)
		}
	}

	innerHTML(element: HTMLElement, html: string) {
		return Diff.innerHTML(element, html)
	}

	post(url: string, body?: any) {
		if(this.isLoading) {
			return Promise.resolve(null)
		}

		if(body !== undefined && typeof body !== "string") {
			body = JSON.stringify(body)
		}

		this.loading(true)

		return fetch(url, {
			method: "POST",
			body,
			credentials: "same-origin"
		})
		.then(response => {
			this.loading(false)

			if(response.status === 200) {
				return Promise.resolve(response)
			}

			return response.text().then(err => {
				throw err
			})
		})
		.catch(err => {
			this.loading(false)
			throw err
		})
	}

	onNewContent(element: HTMLElement) {
		// Do the same as for the content loaded event,
		// except here we are limiting it to the element.
		this.app.ajaxify(element.getElementsByTagName("a"))
		this.lazyLoad(findAllInside("lazy", element))
		this.mountMountables(findAllInside("mountable", element))
		this.assignTooltipOffsets(findAllInside("tip", element))
		this.textAreaFocus()
	}

	scrollTo(target: HTMLElement) {
		const duration = 250.0
		const fullSin = Math.PI / 2
		const contentPadding = 24

		let newScroll = 0
		let finalScroll = Math.max(target.offsetTop - contentPadding, 0)

		// Calculating scrollTop will force a layout - careful!
		let oldScroll = this.app.content.parentElement.scrollTop
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
			this.statusMessage.showError("API object not found")
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
				// If the active element is the search and we press Enter, re-activate search.
				if(activeElement.id === "search" && e.keyCode === 13) {
					actions.search(this, activeElement as HTMLInputElement, e)
				}

				return

			case "TEXTAREA":
				return
		}

		// When called, this will prevent the default action for that key.
		let preventDefault = () => {
			e.preventDefault()
			e.stopPropagation()
		}

		// Ignore hotkeys on contentEditable elements
		if(activeElement.getAttribute("contenteditable") === "true") {
			// Disallow Enter key in contenteditables and make it blur the element instead
			if(e.keyCode === 13) {
				if("blur" in activeElement) {
					(activeElement["blur"] as Function)()
				}

				return preventDefault()
			}

			return
		}

		// "Ctrl" + "," = Settings
		if(e.ctrlKey && e.keyCode === 188) {
			this.app.load("/settings")
			return preventDefault()
		}

		// The following keycodes should not be activated while Ctrl or Alt is held down
		if(e.ctrlKey || e.altKey) {
			return
		}

		// "F" = Search
		if(e.keyCode === 70) {
			let search = document.getElementById("search") as HTMLInputElement

			search.focus()
			search.select()
			return preventDefault()
		}

		// "S" = Toggle sidebar
		if(e.keyCode === 83) {
			this.sideBar.toggle()
			return preventDefault()
		}


		// "+" = Audio speed up
		if(e.key == "+") {
			this.audioPlayer.addSpeed(0.05)
			return preventDefault()
		}

		// "-" = Audio speed down
		if(e.key == "-") {
			this.audioPlayer.addSpeed(-0.05)
			return preventDefault()
		}

		// "J" = Previous track
		if(e.keyCode === 74) {
			this.audioPlayer.previous()
			return preventDefault()
		}

		// "K" = Play/pause
		if(e.keyCode === 75) {
			this.audioPlayer.playPause()
			return preventDefault()
		}

		// "L" = Next track
		if(e.keyCode === 76) {
			this.audioPlayer.next()
			return preventDefault()
		}

		// Number keys activate sidebar menus
		for(let i = 48; i <= 57; i++) {
			if(e.keyCode === i) {
				let index = i === 48 ? 9 : i - 49
				let links = [...findAll("sidebar-link")]

				if(index < links.length) {
					let element = links[index] as HTMLElement

					element.click()
					return preventDefault()
				}
			}
		}
	}

	// This is called every time an uncaught JavaScript error is thrown
	onError(evt: ErrorEvent) {
		let report = {
			message: evt.message,
			stack: evt.error.stack,
			fileName: evt.filename,
			lineNumber: evt.lineno,
			columnNumber: evt.colno,
		}

		this.post("/api/new/clienterrorreport", report)
		.then(() => console.log("Successfully reported the error to the website staff."))
		.catch(() => console.warn("Failed reporting the error to the website staff."))
	}
}