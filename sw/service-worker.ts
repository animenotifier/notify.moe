// pack:ignore

const RELOADS = new Map<string, Promise<Response>>()
const ETAGS = new Map<string, string>()
const CACHEREFRESH = new Map<string, Promise<void>>()
const EXCLUDECACHE = new Set<string>([
	// API requests
	"/api/",

	// PayPal stuff
	"/paypal/",

	// List imports
	"/import/",

	// Infinite scrolling
	"/from/",

	// Chrome extension
	"chrome-extension",

	// Authorization paths /auth/ and /logout are not listed here because they are handled in a special way.
])

class MyCache {
	version: string

	constructor(version: string) {
		this.version = version
	}

	store(request: RequestInfo, response: Response) {
		return caches.open(this.version).then(cache => {
			return cache.put(request, response)
		})
	}
}

class MyServiceWorker {
	cache: MyCache
	currentCSP: string

	constructor() {
		this.cache = new MyCache("v-4")
		this.currentCSP = ""

		self.addEventListener("install", (evt: InstallEvent) => evt.waitUntil(this.onInstall(evt)))
		self.addEventListener("activate", (evt: any) => evt.waitUntil(this.onActivate(evt)))
		self.addEventListener("fetch", (evt: FetchEvent) => evt.waitUntil(this.onRequest(evt)))
		self.addEventListener("message", (evt: any) => evt.waitUntil(this.onMessage(evt)))
		self.addEventListener("push", (evt: PushEvent) => evt.waitUntil(this.onPush(evt)))
		self.addEventListener("pushsubscriptionchange", (evt: any) => evt.waitUntil(this.onPushSubscriptionChange(evt)))
		self.addEventListener("notificationclick", (evt: NotificationEvent) => evt.waitUntil(this.onNotificationClick(evt)))
	}

	onInstall(evt: InstallEvent) {
		console.log("service worker install")

		return (self as any).skipWaiting().then(() => {
			return this.installCache()
		})
	}

	onActivate(evt: any) {
		console.log("service worker activate")

		// Only keep current version of the cache and delete old caches
		let cacheWhitelist = [this.cache.version]

		let deleteOldCache = caches.keys().then(keyList => {
			return Promise.all(keyList.map(key => {
				if(cacheWhitelist.indexOf(key) === -1) {
					return caches.delete(key)
				}
			}))
		})

		// Immediate claim helps us gain control over a new client immediately
		let immediateClaim = (self as any).clients.claim()

		return Promise.all([
			deleteOldCache,
			immediateClaim
		])
	}

	onRequest(evt: FetchEvent) {
		let request = evt.request as Request

		// If it's not a GET request, fetch it normally
		if(request.method !== "GET") {
			return evt.respondWith(fetch(request))
		}

		// Clear cache on authentication and fetch it normally
		if(request.url.includes("/auth/") || request.url.includes("/logout")) {
			return evt.respondWith(caches.delete(this.cache.version).then(() => fetch(request)))
		}

		// Exclude certain URLs from being cached
		for(let pattern of EXCLUDECACHE.keys()) {
			if(request.url.includes(pattern)) {
				return evt.respondWith(fetch(request))
			}
		}

		// If the request included the header "X-CacheOnly", return a cache-only response.
		// This is used in reloads to avoid generating a 2nd request after a cache refresh.
		if(request.headers.get("X-CacheOnly") === "true") {
			return evt.respondWith(this.fromCache(request))
		}

		// Start fetching the request
		let refresh = fetch(request).then(response => {
			let clone = response.clone()

			// Save the new version of the resource in the cache
			let cacheRefresh = this.cache.store(request, clone)

			CACHEREFRESH.set(request.url, cacheRefresh)

			return response
		})

		// Save in map
		RELOADS.set(request.url, refresh)

		// Forced reload
		let servedETag = undefined

		let onResponse = response => {
			servedETag = response.headers.get("ETag")
			ETAGS.set(request.url, servedETag)
			return response
		}

		if(request.headers.get("X-Reload") === "true") {
			return evt.respondWith(refresh.then(onResponse))
		}

		// Try to serve cache first and fall back to network response
		let networkOrCache = this.fromCache(request).then(onResponse).catch(error => {
			// console.log("Cache MISS:", request.url)
			return refresh
		})

		return evt.respondWith(networkOrCache)
	}

	onMessage(evt: any) {
		let message = JSON.parse(evt.data)

		switch(message.type) {
			case "loaded":
				this.onDOMContentLoaded(evt, message.url)
				break
		}
	}

	// onDOMContentLoaded is called when the client sent this service worker
	// a message that the page has been loaded.
	onDOMContentLoaded(evt: any, url: string) {
		let refresh = RELOADS.get(url)
		let servedETag = ETAGS.get(url)

		// If the user requests a sub-page we should prefetch the full page, too.
		if(url.includes("/_/") && !url.includes("/_/search/")) {
			this.prefetchFullPage(url)
		}

		if(!refresh || !servedETag) {
			return Promise.resolve()
		}

		return refresh.then((response: Response) => {
			// When the actual network request was used by the client, response.bodyUsed is set.
			// In that case the client is already up to date and we don"t need to tell the client to do a refresh.
			if(response.bodyUsed) {
				return
			}

			// Get the ETag of the cached response we sent to the client earlier.
			let eTag = response.headers.get("ETag")

			// Update ETag
			ETAGS.set(url, eTag)

			// Get CSP
			let oldCSP = this.currentCSP
			let csp = response.headers.get("Content-Security-Policy")

			// If the CSP and therefore the sha-1 hash of the CSS changed, we need to do a reload.
			if(csp != oldCSP) {
				this.currentCSP = csp

				if(oldCSP !== "") {
					return this.forceClientReloadPage(url, evt.source)
				}
			}

			// If the ETag changed, we need to do a reload.
			if(eTag !== servedETag) {
				return this.forceClientReloadContent(url, evt.source)
			}

			// Do nothing
			return Promise.resolve()
		})
	}

	prefetchFullPage(url: string) {
		let fullPage = new Request(url.replace("/_/", "/"))

		// Disable HTTP/2 push responses
		let headers = new Headers({
			"X-Source": "service-worker"
		})

		let fullPageRefresh = fetch(fullPage, {
			credentials: "same-origin",
			headers
		}).then(response => {
			// Save the new version of the resource in the cache
			let cacheRefresh = caches.open(this.cache.version).then(cache => {
				return cache.put(fullPage, response)
			})

			CACHEREFRESH.set(fullPage.url, cacheRefresh)
			return response
		})

		// Save in map
		RELOADS.set(fullPage.url, fullPageRefresh)
	}

	onPush(evt: PushEvent) {
		var payload = evt.data ? evt.data.json() : {}

		return (self as any).registration.showNotification(payload.title, {
			body: payload.message,
			icon: payload.icon,
			image: payload.image,
			data: payload.link,
			badge: "https://notify.moe/brand/64.png"
		})
	}

	onPushSubscriptionChange(evt: any) {
		return (self as any).registration.pushManager.subscribe(evt.oldSubscription.options)
		.then(async subscription => {
			console.log("send subscription to server...")

			let rawKey = subscription.getKey("p256dh")
			let key = rawKey ? btoa(String.fromCharCode.apply(null, new Uint8Array(rawKey))) : ""

			let rawSecret = subscription.getKey("auth")
			let secret = rawSecret ? btoa(String.fromCharCode.apply(null, new Uint8Array(rawSecret))) : ""

			let endpoint = subscription.endpoint

			let pushSubscription = {
				endpoint,
				p256dh: key,
				auth: secret,
				platform: navigator.platform,
				userAgent: navigator.userAgent,
				screen: {
					width: window.screen.width,
					height: window.screen.height
				}
			}

			let user = await fetch("/api/me").then(response => response.json())

			return fetch("/api/pushsubscriptions/" + user.id + "/add", {
				method: "POST",
				credentials: "same-origin",
				body: JSON.stringify(pushSubscription)
			})
		})
	}

	onNotificationClick(evt: NotificationEvent) {
		let notification = evt.notification
		notification.close()

		return (self as any).clients.matchAll().then(function(clientList) {
			// If we have a link, use that link to open a new window.
			let url = notification.data

			if(url) {
				return (self as any).clients.openWindow(url)
			}

			// If there is at least one client, focus it.
			if(clientList.length > 0) {
				return clientList[0].focus()
			}

			// Otherwise open a new window
			return (self as any).clients.openWindow("https://notify.moe")
		})
	}

	forceClientReloadContent(url: string, eventSource: any) {
		let message = {
			type: "new content",
			url
		}

		this.postMessageAfterPromise(message, CACHEREFRESH.get(url), eventSource)
	}

	forceClientReloadPage(url: string, eventSource: any) {
		let message = {
			type: "reload page",
			url
		}

		this.postMessageAfterPromise(message, RELOADS.get(url.replace("/_/", "/")), eventSource)
	}

	postMessageAfterPromise(message: any, promise: Promise<any>, eventSource: any) {
		if(!promise) {
			console.log("forcing reload, cache refresh null")
			return eventSource.postMessage(JSON.stringify(message))
		}

		return promise.then(() => {
			console.log("forcing reload after cache refresh")
			eventSource.postMessage(JSON.stringify(message))
		})
	}

	installCache() {
		// TODO: Implement a solution that caches resources with credentials: "same-origin"
		return Promise.resolve()

		// return caches.open(this.cache.version).then(cache => {
		// 	return cache.addAll([
		// 		"./",
		// 		"./scripts",
		// 		"https://fonts.gstatic.com/s/ubuntu/v11/4iCs6KVjbNBYlgoKfw7z.ttf"
		// 	])
		// })
	}

	fromCache(request) {
		return caches.open(this.cache.version).then(cache => {
			return cache.match(request).then(matching => {
				if(matching) {
					// console.log("Cache HIT:", request.url)
					return Promise.resolve(matching)
				}

				return Promise.reject("no-match")
			})
		})
	}
}

const serviceWorker = new MyServiceWorker()
