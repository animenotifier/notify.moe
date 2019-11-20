import * as actions from "./Actions"
import { uploadAnalytics } from "./Analytics"
import Application from "./Application"
import AudioPlayer from "./AudioPlayer"
import { displayAiringDate, displayDate, displayTime } from "./DateView"
import Diff from "./Diff"
import ToolTip from "./Elements/tool-tip/tool-tip"
import infiniteScroll from "./infiniteScroll"
import NotificationManager from "./NotificationManager"
import PushManager from "./PushManager"
import receiveServerEvents from "./ServerEvent/receiveServerEvents"
import ServiceWorkerManager from "./ServiceWorkerManager"
import SideBar from "./SideBar"
import StatusMessage from "./StatusMessage"
import User from "./User"
import delay from "./Utils/delay"
import emptyPixel from "./Utils/emptyPixel"
import findAll from "./Utils/findAll"
import findAllInside from "./Utils/findAllInside"
import requestIdleCallback from "./Utils/requestIdleCallback"
import supportsWebP from "./Utils/supportsWebP"
import swapElements from "./Utils/swapElements"
import VideoPlayer from "./VideoPlayer"
import * as WebComponents from "./WebComponents"

export default class AnimeNotifier {
	public isLoading: boolean
	public app: Application
	public statusMessage: StatusMessage
	public notificationManager: NotificationManager | undefined
	public currentMediaId: string
	public audioPlayer: AudioPlayer
	public videoPlayer: VideoPlayer
	public user: User | null
	public sideBar: SideBar
	public pushManager: PushManager

	private title: string
	private webpCheck: Promise<boolean>
	private webpEnabled: boolean
	private visibilityObserver: IntersectionObserver
	private serviceWorkerManager: ServiceWorkerManager
	private diffCompletedForCurrentPath: boolean
	private tip: ToolTip

	constructor() {
		this.app = new Application()
		this.title = "Anime Notifier"
		this.isLoading = true

		// These classes will never be removed on DOM diffs
		Diff.persistentClasses.add("mounted")
		Diff.persistentClasses.add("element-found")
		Diff.persistentClasses.add("active")

		// Never remove src property on diffs
		Diff.persistentAttributes.add("src")
	}

	public init() {
		// App init
		this.app.init()

		// Event listeners
		document.addEventListener("readystatechange", () => this.onReadyStateChange())
		document.addEventListener("DOMContentLoaded", () => this.onContentLoaded())

		// If we finished loading the DOM (either "interactive" or "complete" state),
		// immediately trigger the event listener functions.
		if(document.readyState !== "loading") {
			this.app.emit("DOMContentLoaded")
			this.run()
		}

		// Idle
		requestIdleCallback(() => this.onIdle())
	}

	public reloadContent(cached?: boolean) {
		const headers = new Headers()

		if(cached) {
			headers.set("X-Force-Cache", "true")
		} else {
			headers.set("X-No-Cache", "true")
		}

		const path = this.app.currentPath

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

	public reloadPage() {
		console.log("reload page", this.app.currentPath)

		const path = this.app.currentPath

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

	public async diff(url: string) {
		if(url === this.app.currentPath) {
			return
		}

		const path = "/_" + url

		try {
			// Start the request
			const request = fetch(path, {
				credentials: "same-origin"
			})
			.then(response => response.text())

			history.pushState(url, "", url)
			this.app.currentPath = url
			this.diffCompletedForCurrentPath = false
			this.app.markActiveLinks()
			this.unmountMountables()
			this.loading(true)

			// Delay by mountable-transition-speed
			await delay(150)

			const html = await request

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

	public post(url: string, body?: any) {
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

	public loading(newState: boolean) {
		this.isLoading = newState

		if(this.isLoading) {
			document.documentElement.style.cursor = "progress"
			this.app.loading.classList.remove(this.app.fadeOutClass)
		} else {
			document.documentElement.style.cursor = "auto"
			this.app.loading.classList.add(this.app.fadeOutClass)
		}
	}

	public onNewContent(element: HTMLElement) {
		// Do the same as for the content loaded event,
		// except here we are limiting it to the element.
		this.app.ajaxify(element.getElementsByTagName("a"))
		this.lazyLoad(findAllInside("lazy", element))
		this.mountMountables(findAllInside("mountable", element))
		this.prepareTooltips(findAllInside("tip", element))
		this.textAreaFocus()
	}

	public scrollTo(target: HTMLElement) {
		const duration = 250.0
		const fullSin = Math.PI / 2
		const contentPadding = 23

		let newScroll = 0
		const finalScroll = Math.max(target.getBoundingClientRect().top - contentPadding, 0)

		// Calculating scrollTop will force a layout - careful!
		const contentContainer = this.app.content.parentElement as HTMLElement
		const oldScroll = contentContainer.scrollTop
		const scrollDistance = finalScroll - oldScroll

		if(scrollDistance > 0 && scrollDistance < 1) {
			return
		}

		const timeStart = Date.now()
		const timeEnd = timeStart + duration

		const scroll = () => {
			const time = Date.now()
			let progress = (time - timeStart) / duration

			if(progress > 1.0) {
				progress = 1.0
			}

			newScroll = oldScroll + scrollDistance * Math.sin(progress * fullSin)
			contentContainer.scrollTop = newScroll

			if(time < timeEnd && newScroll !== finalScroll) {
				window.requestAnimationFrame(scroll)
			}
		}

		window.requestAnimationFrame(scroll)
	}

	public findAPIEndpoint(element: HTMLElement | null): string {
		while(element) {
			if(element.dataset.api !== undefined) {
				return element.dataset.api
			}

			element = element.parentElement
		}

		this.statusMessage.showError("API object not found")
		throw "API object not found"
	}

	public markPlayingMedia() {
		for(const element of findAll("media-play-area")) {
			if(element.dataset.mediaId === this.currentMediaId) {
				element.classList.add("playing")
			}
		}
	}

	public mountMountables(elements?: IterableIterator<HTMLElement>) {
		if(!elements) {
			elements = findAll("mountable")
		}

		this.modifyDelayed(elements, element => element.classList.add("mounted"))
	}

	public unmountMountables() {
		for(const element of findAll("mountable")) {
			if(element.classList.contains("never-unmount")) {
				continue
			}

			Diff.mutations.queue(() => element.classList.remove("mounted"))
		}
	}

	public async updatePushUI() {
		if(!this.app.currentPath.includes("/settings/notifications")) {
			return
		}

		const enableButton = document.getElementById("notifications-enable") as HTMLButtonElement
		const disableButton = document.getElementById("notifications-disable") as HTMLButtonElement
		const testButton = document.getElementById("notifications-test") as HTMLButtonElement
		const footer = document.getElementById("notifications-footer") as HTMLElement

		if(!this.pushManager.pushSupported) {
			enableButton.classList.add("hidden")
			disableButton.classList.add("hidden")
			testButton.classList.add("hidden")
			footer.innerHTML = "Your browser doesn't support push notifications!"
			return
		}

		const subscription = await this.pushManager.subscription()

		if(subscription) {
			enableButton.classList.add("hidden")
			disableButton.classList.remove("hidden")
		} else {
			enableButton.classList.remove("hidden")
			disableButton.classList.add("hidden")
		}
	}

	public assignActions() {
		for(const element of findAll("action")) {
			const actionTrigger = element.dataset.trigger
			const actionName = element.dataset.action

			// Filter out invalid definitions
			if(!actionTrigger || !actionName) {
				continue
			}

			const oldAction = element["action assigned"]

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
			const actionHandler = e => {
				if(!actionName) {
					return
				}

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

	private onReadyStateChange() {
		if(document.readyState !== "interactive") {
			return
		}

		this.run()
	}

	private run() {
		// Initiate the elements we need
		this.app.content = document.getElementById("content") as HTMLElement
		this.app.loading = document.getElementById("loading") as HTMLElement

		// Web components
		WebComponents.register()

		// Tooltip
		this.tip = new ToolTip()
		document.body.appendChild(this.tip)
		document.addEventListener("linkclicked", () => this.tip.classList.add("fade-out"))

		// Enable lazy load
		this.visibilityObserver = new IntersectionObserver(
			entries => {
				for(const entry of entries) {
					if(entry.isIntersecting) {
						entry.target["became visible"]()
						this.visibilityObserver.unobserve(entry.target)
					}
				}
			},
			{}
		)

		// Status message
		this.statusMessage = new StatusMessage(
			document.getElementById("status-message") as HTMLElement,
			document.getElementById("status-message-text") as HTMLElement
		)

		this.app.onError = (error: Error) => {
			this.statusMessage.showError(error, 3000)
		}

		// User
		const userElement = document.getElementById("user")

		if(userElement && userElement.dataset.id) {
			this.user = new User(userElement.dataset.id)

			if(userElement.dataset.pro === "true") {
				const theme = userElement.dataset.theme

				// Don't apply light theme on load because
				// it's already the standard theme.
				if(theme && theme !== "light") {
					actions.applyTheme(theme)
				}
			}
		}

		// Push manager
		this.pushManager = new PushManager()

		// Notification manager
		if(this.user) {
			this.notificationManager = new NotificationManager(
				document.getElementById("notification-icon") as HTMLElement,
				document.getElementById("notification-count") as HTMLElement
			)
		}

		// Audio player
		this.audioPlayer = new AudioPlayer(this)

		// Video player
		this.videoPlayer = new VideoPlayer(this)

		// Sidebar control
		this.sideBar = new SideBar(document.getElementById("sidebar"))

		// Infinite scrolling
		if(this.app.content.parentElement) {
			infiniteScroll(this.app.content.parentElement, 150)
		}

		// WebP
		this.webpCheck = supportsWebP().then(val => this.webpEnabled = val)

		// Loading
		this.loading(false)
	}

	private onContentLoaded() {
		// Stop watching all the objects from the previous page.
		this.visibilityObserver.disconnect()

		Promise.all([
			Promise.resolve().then(() => this.mountMountables()),
			Promise.resolve().then(() => this.lazyLoad()),
			Promise.resolve().then(() => this.displayLocalDates()),
			Promise.resolve().then(() => this.setSelectBoxValue()),
			Promise.resolve().then(() => this.textAreaFocus()),
			Promise.resolve().then(() => this.markPlayingMedia()),
			Promise.resolve().then(() => this.assignActions()),
			Promise.resolve().then(() => this.updatePushUI()),
			Promise.resolve().then(() => this.dragAndDrop()),
			Promise.resolve().then(() => this.colorBoxes()),
			Promise.resolve().then(() => this.loadCharacterRanking()),
			Promise.resolve().then(() => this.prepareTooltips()),
			Promise.resolve().then(() => this.countUp())
		])

		// Apply page title
		this.applyPageTitle()

		// Auto-focus first input element on welcome page.
		if(location.pathname === "/welcome") {
			const firstInput = this.app.content.getElementsByTagName("input")[0] as HTMLInputElement

			if(firstInput) {
				firstInput.focus()
			}
		}
	}

	private applyPageTitle() {
		const headers = document.getElementsByTagName("h1")

		if(this.app.currentPath === "/" || headers.length === 0 || headers[0].textContent === "NOTIFY.MOE") {
			if(document.title !== this.title) {
				document.title = this.title
			}
		} else if(headers[0].textContent) {
			document.title = headers[0].textContent
		}
	}

	private textAreaFocus() {
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

	private async onIdle() {
		// Register event listeners
		document.addEventListener("keydown", this.onKeyDown.bind(this), false)
		window.addEventListener("popstate", this.onPopState.bind(this))
		window.addEventListener("error", this.onError.bind(this))

		// Service worker
		this.serviceWorkerManager = new ServiceWorkerManager(this, "/service-worker")
		this.serviceWorkerManager.register()

		// Analytics
		if(this.user) {
			uploadAnalytics()
		}

		// Offline message
		if(navigator.onLine === false) {
			this.statusMessage.showInfo("You are viewing an offline version of the site now.")
		}

		// Notification manager
		if(this.notificationManager) {
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
			const finalWidth = window.outerWidth < minWidth ? minWidth : window.outerWidth
			const finalHeight = window.outerHeight < minHeight ? minHeight : window.outerHeight

			window.resizeTo(finalWidth, finalHeight)
		}

		// Server sent events
		if(this.user && EventSource) {
			receiveServerEvents(this)
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

	private onBeforeUnload(e: BeforeUnloadEvent) {
		if(this.app.currentPath !== "/new/thread") {
			return
		}

		if(!document.activeElement) {
			return
		}

		if(document.activeElement.tagName !== "TEXTAREA") {
			return
		}

		if((document.activeElement as HTMLTextAreaElement).value.length < 20) {
			return
		}

		// Prevent closing tab on new thread page
		e.returnValue = "You have unsaved changes on the current page. Are you sure you want to leave?"
	}

	private prepareTooltips(elements?: IterableIterator<HTMLElement>) {
		if(!elements) {
			elements = findAll("tip")
		}

		this.tip.setAttribute("active", "false")

		// Assign mouse enter event handler
		for(const element of elements) {
			element.onmouseenter = () => {
				this.tip.classList.remove("fade-out")
				this.tip.show(element)
			}

			element.onmouseleave = () => {
				this.tip.hide()
			}
		}
	}

	private dragAndDrop() {
		if(location.pathname.includes("/animelist/")) {
			for(const listItem of findAll("anime-list-item")) {
				// Skip elements that have their event listeners attached already
				if(listItem["drag-listeners-attached"]) {
					continue
				}

				const name = listItem.getElementsByClassName("anime-list-item-name")[0]
				const imageContainer = listItem.getElementsByClassName("anime-list-item-image-container")[0]

				const onDrag = evt => {
					if(!evt.dataTransfer) {
						return
					}

					const image = imageContainer.getElementsByClassName("anime-list-item-image")[0]

					if(image) {
						evt.dataTransfer.setDragImage(image, 0, 0)
					}

					evt.dataTransfer.setData("text/plain", JSON.stringify({
						api: listItem.dataset.api,
						animeTitle: name.textContent
					}))

					evt.dataTransfer.effectAllowed = "move"
				}

				name.addEventListener("dragstart", onDrag, false)
				imageContainer.addEventListener("dragstart", onDrag, false)

				// Prevent re-attaching the same listeners
				listItem["drag-listeners-attached"] = true
			}

			for(const element of findAll("tab")) {
				// Skip elements that have their event listeners attached already
				if(element["drop-listeners-attached"]) {
					continue
				}

				element.addEventListener("drop", async e => {
					let toElement: HTMLElement | null = e.target as HTMLElement

					// Find tab element
					while(toElement && !toElement.classList.contains("tab")) {
						toElement = toElement.parentElement
					}

					// Ignore a drop on the current status tab
					if(!toElement || toElement.classList.contains("active") || !e.dataTransfer) {
						return
					}

					const data = e.dataTransfer.getData("text/plain")
					let json: any

					try {
						json = JSON.parse(data)
					} catch(err) {
						return
					}

					if(!json || !json.api) {
						return
					}

					e.stopPropagation()
					e.preventDefault()

					const tabText = toElement.textContent

					if(!tabText) {
						return
					}

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
			for(const element of findAll("inventory-slot")) {
				// Skip elements that have their event listeners attached already
				if(element["drag-listeners-attached"]) {
					continue
				}

				element.addEventListener("dragstart", e => {
					if(!element.draggable || !element.dataset.index || !e.dataTransfer) {
						return
					}

					e.dataTransfer.setData("text", element.dataset.index)
				}, false)

				element.addEventListener("dblclick", async _ => {
					if(!element.draggable || !element.dataset.index) {
						return
					}

					const itemName = element.getAttribute("aria-label")

					if(element.dataset.consumable !== "true") {
						return this.statusMessage.showError(itemName + " is not a consumable item.")
					}

					const apiEndpoint = this.findAPIEndpoint(element)

					try {
						await this.post(apiEndpoint + "/use/" + element.dataset.index)
						await this.reloadContent()
						this.statusMessage.showInfo(`You used ${itemName}.`)
					} catch(err) {
						this.statusMessage.showError(err)
					}
				}, false)

				element.addEventListener("dragenter", _ => {
					element.classList.add("drag-enter")
				}, false)

				element.addEventListener("dragleave", _ => {
					element.classList.remove("drag-enter")
				}, false)

				element.addEventListener("dragover", e => {
					e.preventDefault()
				}, false)

				element.addEventListener("drop", async e => {
					element.classList.remove("drag-enter")

					e.stopPropagation()
					e.preventDefault()

					const inventory = element.parentElement

					if(!inventory || !e.dataTransfer) {
						return
					}

					const fromIndex = e.dataTransfer.getData("text")

					if(!fromIndex) {
						return
					}

					const fromElement = inventory.childNodes[fromIndex] as HTMLElement
					const toIndex = element.dataset.index

					if(!toIndex || fromElement === element || fromIndex === toIndex) {
						console.error("Invalid drag & drop from", fromIndex, "to", toIndex)
						return
					}

					// Swap in database
					const apiEndpoint = this.findAPIEndpoint(inventory)

					try {
						await this.post(apiEndpoint + "/swap/" + fromIndex + "/" + toIndex)
					} catch(err) {
						this.statusMessage.showError(err)
					}

					// Swap in UI
					swapElements(fromElement, element)

					fromElement.dataset.index = toIndex
					element.dataset.index = fromIndex
				}, false)

				// Prevent re-attaching the same listeners
				element["drag-listeners-attached"] = true
			}
		}
	}

	private loadCharacterRanking() {
		if(!this.app.currentPath.includes("/character/")) {
			return
		}

		for(const element of findAll("character-ranking")) {
			fetch(`/api/character/${element.dataset.characterId}/ranking`).then(async response => {
				const ranking = await response.json()

				if(!ranking.rank) {
					return
				}

				Diff.mutations.queue(() => {
					const percentile = Math.ceil(ranking.percentile * 100)

					element.textContent = "#" + ranking.rank.toString()
					element.title = "Top " + percentile + "%"
				})
			})
		}
	}

	private colorBoxes() {
		if(!this.app.currentPath.includes("/explore/color/") && !this.app.currentPath.includes("/settings")) {
			return
		}

		for(const element of findAll("color-box")) {
			Diff.mutations.queue(() => {
				if(!element.dataset.color) {
					console.error("color-box missing data-color attribute:", element)
					return
				}

				element.style.backgroundColor = element.dataset.color
			})
		}
	}

	private countUp() {
		if(!this.app.currentPath.includes("/paypal/success")) {
			return
		}

		for(const element of findAll("count-up")) {
			if(!element.textContent) {
				console.error("count-up missing text content:", element)
				continue
			}

			const final = parseInt(element.textContent, 10)
			const duration = 2000.0
			const start = Date.now()

			element.textContent = "0"

			const callback = () => {
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

	private setSelectBoxValue() {
		for(const element of document.getElementsByTagName("select")) {
			const attributeValue = element.getAttribute("value")

			if(!attributeValue) {
				console.error("Select box without a value:", element)
				continue
			}

			element.value = attributeValue
		}
	}

	private displayLocalDates() {
		const now = new Date()

		for(const element of findAll("utc-airing-date")) {
			displayAiringDate(element, now)
		}

		for(const element of findAll("utc-date")) {
			displayDate(element, now)
		}

		for(const element of findAll("utc-date-absolute")) {
			displayTime(element)
		}
	}

	private async lazyLoad(elements?: IterableIterator<Element>) {
		if(!elements) {
			elements = findAll("lazy")
		}

		await this.webpCheck

		for(const element of elements) {
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

	private lazyLoadImage(element: HTMLImageElement) {
		const pixelRatio = window.devicePixelRatio

		// Once the image becomes visible, load it
		element["became visible"] = () => {
			const dataSrc = element.dataset.src

			if(!dataSrc) {
				console.error("Image missing data-src attribute:", element)
				return
			}

			const dotPos = dataSrc.lastIndexOf(".")
			let base = dataSrc.substring(0, dotPos)
			let extension = ""

			// Replace URL with WebP if supported
			if(this.webpEnabled && element.dataset.webp === "true" && !dataSrc.endsWith(".svg")) {
				const queryPos = dataSrc.lastIndexOf("?")

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
				if(base.includes("/anime/") || base.includes("/groups/") || (base.includes("/characters/") && !base.includes("/large/"))) {
					base += "@2"
				}
			}

			const finalSrc = base + extension

			if(element.src !== finalSrc && element.src !== "https:" + finalSrc && element.src !== "https://notify.moe" + finalSrc) {
				// Show average color
				if(element.dataset.color) {
					element.src = emptyPixel
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
					// Try loading from the origin server if our CDN failed
					if(element.src.includes("media.notify.moe/")) {
						console.warn(`CDN failed loading ${element.src}`)
						element.src = element.src.replace("media.notify.moe/", "notify.moe/")
						return
					}

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

	private lazyLoadIFrame(element: HTMLIFrameElement) {
		// Once the iframe becomes visible, load it
		element["became visible"] = () => {
			if(!element.dataset.src) {
				console.error("IFrame missing data-src attribute:", element)
				return
			}

			// If the source is already set correctly, don't set it again to avoid iframe flickering.
			if(element.src !== element.dataset.src && element.src !== (window.location.protocol + element.dataset.src)) {
				element.src = element.dataset.src
			}

			Diff.mutations.queue(() => element.classList.add("element-found"))
		}

		this.visibilityObserver.observe(element)
	}

	private lazyLoadVideo(video: HTMLVideoElement) {
		const hideControlsDelay = 1500

		// Once the video becomes visible, load it
		video["became visible"] = () => {
			if(!video["listeners attached"]) {
				const videoParent = video.parentElement

				if(!videoParent) {
					console.error("video has no parent element")
					return
				}

				// Prevent context menu
				video.addEventListener("contextmenu", e => e.preventDefault())

				// Show and hide controls on mouse movement
				const controls = videoParent.getElementsByClassName("video-controls")[0]
				const playButton = videoParent.getElementsByClassName("video-control-play")[0] as HTMLElement
				const pauseButton = videoParent.getElementsByClassName("video-control-pause")[0] as HTMLElement

				const hideControls = () => {
					controls.classList.add("fade-out")
					video.style.cursor = "none"
				}

				const showControls = () => {
					controls.classList.remove("fade-out")
					video.style.cursor = "default"
				}

				video.addEventListener("mousemove", () => {
					showControls()
					clearTimeout(video["hideControlsTimeout"])
					video["hideControlsTimeout"] = setTimeout(hideControls, hideControlsDelay)
				})

				const progressElement = videoParent.getElementsByClassName("video-progress")[0] as HTMLElement
				const progressClickable = videoParent.getElementsByClassName("video-progress-clickable")[0]
				const timeElement = videoParent.getElementsByClassName("video-time")[0]

				video.addEventListener("canplay", () => {
					video["playable"] = true
				})

				video.addEventListener("timeupdate", () => {
					if(!video["playable"]) {
						return
					}

					const time = video.currentTime
					const minutes = Math.trunc(time / 60)
					const seconds = Math.trunc(time) % 60
					const paddedSeconds = ("00" + seconds).slice(-2)

					Diff.mutations.queue(() => {
						timeElement.textContent = `${minutes}:${paddedSeconds}`
						progressElement.style.transform = `scaleX(${time / video.duration})`
					})
				})

				video.addEventListener("waiting", () => {
					this.loading(true)
				})

				video.addEventListener("playing", () => {
					this.loading(false)
				})

				video.addEventListener("play", () => {
					playButton.style.display = "none"
					pauseButton.style.display = "block"
				})

				video.addEventListener("pause", () => {
					playButton.style.display = "block"
					pauseButton.style.display = "none"
				})

				progressClickable.addEventListener("click", (e: MouseEvent) => {
					const rect = progressClickable.getBoundingClientRect()
					const x = e.clientX
					const progress = (x - rect.left) / rect.width
					video.currentTime = progress * video.duration
					video.dispatchEvent(new Event("timeupdate"))
					e.stopPropagation()
				})

				video["listeners attached"] = true
			}

			let modified = false

			for(const child of video.children) {
				if(child.tagName !== "SOURCE") {
					continue
				}

				const element = child as HTMLSourceElement

				if(!element.dataset.src || !element.dataset.type) {
					console.error("Source element missing data-src or data-type attribute:", element)
					continue
				}

				if(element.src !== element.dataset.src) {
					element.src = element.dataset.src
					modified = true
				}

				if(element.type !== element.dataset.type) {
					element.type = element.dataset.type
					modified = true
				}
			}

			if(modified) {
				video["playable"] = false

				Diff.mutations.queue(() => {
					video.load()
					video.classList.add("element-found")
				})
			}
		}

		this.visibilityObserver.observe(video)
	}

	private modifyDelayed(elements: IterableIterator<HTMLElement>, func: (element: HTMLElement) => void) {
		const maxDelay = 2500
		const delayTime = 20

		let time = 0
		const start = Date.now()
		const maxTime = start + maxDelay

		const mountableTypes = new Map<string, number>()
		const mountableTypeMutations = new Map<string, any[]>()

		for(const element of elements) {
			// Skip already mounted elements.
			// This helps a lot when dealing with infinite scrolling
			// where the first elements are already mounted.
			if(element.classList.contains("mounted")) {
				continue
			}

			const type = element.dataset.mountableType || "general"
			const typeTime = mountableTypes.get(type)

			if(typeTime !== undefined) {
				time = typeTime + delayTime
				mountableTypes.set(type, time)
			} else {
				time = start
				mountableTypes.set(type, time)
				mountableTypeMutations.set(type, [])
			}

			if(time > maxTime) {
				time = maxTime
			}

			const mutations = mountableTypeMutations.get(type) as any[]

			mutations.push({
				element,
				time
			})
		}

		for(const mutations of mountableTypeMutations.values()) {
			let mutationIndex = 0

			const updateBatch = () => {
				const now = Date.now()

				for(; mutationIndex < mutations.length; mutationIndex++) {
					const mutation = mutations[mutationIndex]

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

	private onPopState(e: PopStateEvent) {
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

	private onKeyDown(e: KeyboardEvent) {
		const activeElement = document.activeElement

		if(!activeElement) {
			return
		}

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
		const preventDefault = () => {
			e.preventDefault()
			e.stopPropagation()
		}

		// Ignore hotkeys on contentEditable elements
		if(activeElement.getAttribute("contenteditable") === "true") {
			// Disallow Enter key in contenteditables and make it blur the element instead
			if(e.keyCode === 13) {
				if("blur" in activeElement) {
					(activeElement["blur"] as () => void)()
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
			const search = document.getElementById("search") as HTMLInputElement

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
		if(e.key === "+") {
			this.audioPlayer.addSpeed(0.05)
			return preventDefault()
		}

		// "-" = Audio speed down
		if(e.key === "-") {
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

		// Space = Toggle play
		if(e.keyCode === 32) {
			// this.audioPlayer.playPause()
			this.videoPlayer.playPause()
			return preventDefault()
		}

		// Number keys activate sidebar menus
		for(let i = 48; i <= 57; i++) {
			if(e.keyCode === i) {
				const index = i === 48 ? 9 : i - 49
				const links = [...findAll("sidebar-link")]

				if(index < links.length) {
					const element = links[index] as HTMLElement

					element.click()
					return preventDefault()
				}
			}
		}
	}

	// This is called every time an uncaught JavaScript error is thrown
	private async onError(evt: ErrorEvent) {
		const report = {
			message: evt.message,
			stack: evt.error.stack,
			fileName: evt.filename,
			lineNumber: evt.lineno,
			columnNumber: evt.colno,
		}

		try {
			await this.post("/api/new/clienterrorreport", report)
			console.log("Successfully reported the error to the website staff.")
		} catch(err) {
			console.warn("Failed reporting the error to the website staff:", err)
		}
	}
}
