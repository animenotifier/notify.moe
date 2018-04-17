// pack:ignore

// This is the service worker for notify.moe.
// When installed, it will intercept all requests made by the browser
// and return a cache-first response.

//         request                 request
// Browser -------> Service Worker ------> notify.moe Server
//         <-------
//         response (cache)
//                                 <------
//                                 response (network)
//         <-------
//         response (network)
//
// -> Diff cache with network response.

// By always returning cache first,
// we avoid latency problems on high latency connections like mobile
// networks. While the cache is being served, we start a real network
// request to the server to see if the resource changed. We compare the
// E-Tag of the cached and latest version of the resource. If the E-Tag
// of the current document changed, we send a message to the client
// that will cause the client to reload (diff) the page. It is not a real
// page reload as we will only calculate a DOM diff on the contents.
// If the style or script resources changed after being served, we need
// to force a real page reload.

// Promises
// const CACHEREFRESH = new Map<string, Promise<void>>()

// E-Tags that we served for a given URL
// const ETAGS = new Map<string, string>()

// MyServiceWorker is the process that controls all the tabs in a browser.
class MyServiceWorker {
	cache: MyCache
	reloads: Map<string, Promise<Response>>
	excludeCache: Set<string>

	constructor() {
		this.cache = new MyCache("v-6")
		this.reloads = new Map<string, Promise<Response>>()

		// When these patterns are matched for the request URL, we exclude them from being
		// served cache-first and instead serve them via a network request.
		// Note that the service worker URL is automatically excluded from fetch events
		// and therefore doesn't need to be added here.
		this.excludeCache = new Set<string>([
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

		self.addEventListener("install", (evt: InstallEvent) => evt.waitUntil(this.onInstall(evt)))
		self.addEventListener("activate", (evt: any) => evt.waitUntil(this.onActivate(evt)))
		self.addEventListener("fetch", (evt: FetchEvent) => evt.waitUntil(this.onRequest(evt)))
		self.addEventListener("message", (evt: any) => evt.waitUntil(this.onMessage(evt)))
		self.addEventListener("push", (evt: PushEvent) => evt.waitUntil(this.onPush(evt)))
		self.addEventListener("pushsubscriptionchange", (evt: any) => evt.waitUntil(this.onPushSubscriptionChange(evt)))
		self.addEventListener("notificationclick", (evt: NotificationEvent) => evt.waitUntil(this.onNotificationClick(evt)))
	}

	async onInstall(evt: InstallEvent) {
		console.log("service worker install")

		await self.skipWaiting()
		await this.installCache()
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
		let immediateClaim = self.clients.claim()

		return Promise.all([
			deleteOldCache,
			immediateClaim
		])
	}

	// onRequest intercepts all browser requests.
	// Simply returning, without calling evt.respondWith(),
	// will let the browser deal with the request normally.
	async onRequest(evt: FetchEvent) {
		let request = evt.request as Request

		// If it's not a GET request, fetch it normally.
		// Let the browser handle XHR upload requests via POST,
		// so that we can receive upload progress events.
		if(request.method !== "GET") {
			return
		}

		// DevTools opening will trigger these "only-if-cached" requests.
		// https://bugs.chromium.org/p/chromium/issues/detail?id=823392
		if((request.cache as string) === "only-if-cached" && request.mode !== "same-origin") {
			return
		}

		// Exclude certain URLs from being cached.
		for(let pattern of this.excludeCache.keys()) {
			if(request.url.includes(pattern)) {
				return
			}
		}

		// If the request has cache set to "force-cache", return a cache-only response.
		// This is used in reloads to avoid generating a 2nd request after a cache refresh.
		if(request.headers.get("X-Force-Cache") === "true") {
			return evt.respondWith(this.cache.serve(request))
		}

		// --------------------------------------------------------------------------------
		// Cross-origin requests.
		// --------------------------------------------------------------------------------

		// These hosts don't support CORS. Always load via network.
		if(request.url.startsWith("https://img.youtube.com/")) {
			return
		}

		// Use CORS for cross-origin requests.
		if(!request.url.startsWith("https://notify.moe/") && !request.url.startsWith("https://beta.notify.moe/")) {
			request = new Request(request.url, {
				credentials: "omit",
				mode: "cors"
			})
		} else {
			// let relativePath = trimPrefix(request.url, "https://notify.moe")
			// relativePath = trimPrefix(relativePath, "https://beta.notify.moe")
			// console.log(relativePath)
		}

		// --------------------------------------------------------------------------------
		// Network refresh.
		// --------------------------------------------------------------------------------

		// Save response in cache.
		let saveResponseInCache = (response: Response) => {
			let contentType = response.headers.get("Content-Type")

			// Don't cache anything other than text and images.
			if(!contentType.includes("text/") && !contentType.includes("application/javascript") && !contentType.includes("image/")) {
				return response
			}

			// Save response in cache.
			let clone = response.clone()
			this.cache.store(request, clone)

			return response
		}

		let onResponse = (response: Response | null) => {
			return response
		}

		// Refresh resource via a network request.
		let refresh = fetch(request).then(saveResponseInCache)

		// --------------------------------------------------------------------------------
		// Final response.
		// --------------------------------------------------------------------------------

		// Clear cache on authentication and fetch it normally.
		if(request.url.includes("/auth/") || request.url.includes("/logout")) {
			return evt.respondWith(this.cache.clear().then(() => refresh))
		}

		// If the request has cache set to "no-cache",
		// return the network-only response even if it fails.
		if(request.headers.get("X-No-Cache") === "true") {
			return evt.respondWith(refresh)
		}

		// Styles and scripts will be served via network first and fallback to cache.
		if(request.url.endsWith("/styles") || request.url.endsWith("/scripts")) {
			evt.respondWith(this.networkFirst(request, refresh, onResponse))
			return refresh
		}

		// --------------------------------------------------------------------------------
		// Default behavior for most requests.
		// --------------------------------------------------------------------------------

		// // Respond via cache first.
		// evt.respondWith(this.cacheFirst(request, refresh, onResponse))
		// return refresh

		// Serve via network first and fallback to cache.
		evt.respondWith(this.networkFirst(request, refresh, onResponse))
		return refresh
	}

	// onMessage is called when the service worker receives a message from a client (browser tab).
	async onMessage(evt: ServiceWorkerMessageEvent) {
		let message = JSON.parse(evt.data)
		let clientId = (evt.source as any).id
		let client = await MyClient.get(clientId)

		client.onMessage(message)
	}

	// onPush is called on push events and requires the payload to contain JSON information about the notification.
	onPush(evt: PushEvent) {
		var payload = evt.data ? evt.data.json() : {}

		// Notify all clients about the new notification so they can update their notification counter
		this.broadcast({
			type: "new notification"
		})

		// Display the notification
		return self.registration.showNotification(payload.title, {
			body: payload.message,
			icon: payload.icon,
			data: payload.link,
			badge: "https://media.notify.moe/images/brand/64.png"
		})
	}

	onPushSubscriptionChange(evt: any) {
		return self.registration.pushManager.subscribe(evt.oldSubscription.options)
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

			let response = await fetch("/api/me", {credentials: "same-origin"})
			let user = await response.json()

			return fetch("/api/pushsubscriptions/" + user.id + "/add", {
				method: "POST",
				credentials: "same-origin",
				body: JSON.stringify(pushSubscription)
			})
		})
	}

	// onNotificationClick is called when the user clicks on a notification.
	onNotificationClick(evt: NotificationEvent) {
		let notification = evt.notification
		notification.close()

		return self.clients.matchAll().then(function(clientList) {
			// If we have a link, use that link to open a new window.
			let url = notification.data

			if(url) {
				return self.clients.openWindow(url)
			}

			// If there is at least one client, focus it.
			if(clientList.length > 0) {
				return (clientList[0] as WindowClient).focus()
			}

			// Otherwise open a new window
			return self.clients.openWindow("https://notify.moe")
		})
	}

	// Broadcast sends a message to all clients (open tabs etc.)
	broadcast(msg: object) {
		const msgText = JSON.stringify(msg)

		self.clients.matchAll().then(function(clientList) {
			for(let client of clientList) {
				client.postMessage(msgText)
			}
		})
	}

	// installCache is called when the service worker is installed for the first time.
	async installCache() {
		let urls = [
			"/",
			"/scripts",
			"/styles",
			"/manifest.json",
			"https://media.notify.moe/images/elements/noise-strong.png",
		]

		let promises = []

		for(let url of urls) {
			let request = new Request(url, {
				credentials: "same-origin",
				mode: "cors"
			})

			let promise = fetch(request).then(response => this.cache.store(request, response))
			promises.push(promise)
		}

		return Promise.all(promises)
	}

	// Serve network first.
	// Fall back to cache.
	async networkFirst(request: Request, network: Promise<Response>, onResponse: (r: Response) => Response): Promise<Response> {
		let response: Response | null

		try {
			response = await network
			// console.log("Network HIT:", request.url)
		} catch(error) {
			// console.log("Network MISS:", request.url, error)

			try {
				response = await this.cache.serve(request)
			} catch(error) {
				return Promise.reject(error)
			}
		}

		return onResponse(response)
	}

	// Serve cache first.
	// Fall back to network.
	async cacheFirst(request: Request, network: Promise<Response>, onResponse: (r: Response) => Response): Promise<Response> {
		let response: Response | null

		try {
			response = await this.cache.serve(request)
			// console.log("Cache HIT:", request.url)
		} catch(error) {
			// console.log("Cache MISS:", request.url, error)

			try {
				response = await network
			} catch(error) {
				return Promise.reject(error)
			}
		}

		return onResponse(response)
	}
}

// MyCache is the cache used by the service worker.
class MyCache {
	version: string

	constructor(version: string) {
		this.version = version
	}

	clear() {
		return caches.delete(this.version)
	}

	async store(request: RequestInfo, response: Response) {
		try {
			// This can fail if the disk space quota has been exceeded.
			let cache = await caches.open(this.version)
			await cache.put(request, response)
		} catch(err) {
			console.log("Disk quota exceeded, can't store in cache:", request, response, err)
		}
	}

	async serve(request: RequestInfo): Promise<Response> {
		let cache = await caches.open(this.version)
		let matching = await cache.match(request)

		if(matching) {
			return matching
		}

		return Promise.reject("no-match")
	}
}

// MyClient represents a single tab in the browser.
class MyClient {
	// MyClient.idToClient is a Map of clients
	static idToClient = new Map<string, MyClient>()

	// MyClient.get retrieves a client by ID
	static async get(id: string): Promise<MyClient> {
		let client = MyClient.idToClient.get(id)

		if(!client) {
			client = new MyClient(await self.clients.get(id))
			MyClient.idToClient.set(id, client)
		}

		return client
	}

	// The actual client
	client: ServiceWorkerClient

	constructor(client: ServiceWorkerClient) {
		this.client = client
	}

	onMessage(message: any) {
		switch(message.type) {
			case "loaded":
				this.onDOMContentLoaded(message.url)
				break

			case "broadcast":
				message.type = message.realType
				delete message.realType
				serviceWorker.broadcast(message)
				break
		}
	}

	// postMessage sends a message to the client.
	postMessage(message: object) {
		this.client.postMessage(JSON.stringify(message))
	}

	// onDOMContentLoaded is called when the client sent this service worker
	// a message that the page has been loaded.
	onDOMContentLoaded(url: string) {
		// ...
	}
}

// trimPrefix removes the prefix from the text.
function trimPrefix(text, prefix) {
	if(text.startsWith(prefix)) {
		return text.slice(prefix.length)
	}

	return text
}

// Initialize the service worker
const serviceWorker = new MyServiceWorker()
